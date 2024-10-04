package orchestrator

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	config "github.com/thirdweb-dev/indexer/configs"
	"github.com/thirdweb-dev/indexer/internal/common"
	"github.com/thirdweb-dev/indexer/internal/metrics"
	"github.com/thirdweb-dev/indexer/internal/storage"
	"github.com/thirdweb-dev/indexer/internal/worker"
)

const DEFAULT_BLOCKS_PER_POLL = 10
const DEFAULT_TRIGGER_INTERVAL = 1000

type Poller struct {
	rpc               common.RPC
	blocksPerPoll     int64
	triggerIntervalMs int64
	storage           storage.IStorage
	lastPolledBlock   *big.Int
	pollUntilBlock    *big.Int
}

type BlockNumberWithError struct {
	BlockNumber *big.Int
	Error       error
}

func NewPoller(rpc common.RPC, storage storage.IStorage) *Poller {
	blocksPerPoll := config.Cfg.Poller.BlocksPerPoll
	if blocksPerPoll == 0 {
		blocksPerPoll = DEFAULT_BLOCKS_PER_POLL
	}
	triggerInterval := config.Cfg.Poller.Interval
	if triggerInterval == 0 {
		triggerInterval = DEFAULT_TRIGGER_INTERVAL
	}
	untilBlock := big.NewInt(int64(config.Cfg.Poller.UntilBlock))
	pollFromBlock := big.NewInt(int64(config.Cfg.Poller.FromBlock))
	lastPolledBlock, err := storage.StagingStorage.GetLastStagedBlockNumber(rpc.ChainID, untilBlock)
	if err != nil || lastPolledBlock == nil || lastPolledBlock.Sign() <= 0 {
		lastPolledBlock = new(big.Int).Sub(pollFromBlock, big.NewInt(1)) // needs to include the first block
		log.Warn().Err(err).Msgf("No last polled block found, setting to %s", lastPolledBlock.String())
	} else {
		// In the case where the start block in staging introduces a gap with main storage,
		// This hack allows us to re-poll from the start block without having to delete the staging data
		if config.Cfg.Poller.ForceFromBlock {
			lastPolledBlock = new(big.Int).Sub(pollFromBlock, big.NewInt(1)) // needs to include the first block
		}
		log.Info().Msgf("Last polled block found: %s", lastPolledBlock.String())
	}
	return &Poller{
		rpc:               rpc,
		triggerIntervalMs: int64(triggerInterval),
		blocksPerPoll:     int64(blocksPerPoll),
		storage:           storage,
		lastPolledBlock:   lastPolledBlock,
		pollUntilBlock:    untilBlock,
	}
}

func (p *Poller) Start() {
	interval := time.Duration(p.triggerIntervalMs) * time.Millisecond
	ticker := time.NewTicker(interval)

	// TODO: make this configurable?
	const numWorkers = 5
	tasks := make(chan struct{}, numWorkers)
	var blockRangeMutex sync.Mutex

	for i := 0; i < numWorkers; i++ {
		go func() {
			for range tasks {
				blockRangeMutex.Lock()
				blockNumbers, err := p.getBlockRange()
				blockRangeMutex.Unlock()

				if err != nil {
					log.Error().Err(err).Msg("Error getting block range")
					continue
				}
				if len(blockNumbers) < 1 {
					log.Debug().Msg("No blocks to poll, skipping")
					continue
				}
				endBlock := blockNumbers[len(blockNumbers)-1]
				if endBlock != nil {
					p.lastPolledBlock = endBlock
				}
				log.Debug().Msgf("Polling %d blocks starting from %s to %s", len(blockNumbers), blockNumbers[0], endBlock)

				endBlockNumberFloat, _ := endBlock.Float64()
				metrics.PollerLastTriggeredBlock.Set(endBlockNumberFloat)

				worker := worker.NewWorker(p.rpc)
				results := worker.Run(blockNumbers)
				p.handleWorkerResults(results)
				if p.reachedPollLimit(endBlock) {
					log.Debug().Msg("Reached poll limit, exiting poller")
					ticker.Stop()
					return
				}
			}
		}()
	}

	for range ticker.C {
		tasks <- struct{}{}
	}

	// Keep the program running (otherwise it will exit)
	select {}
}

func (p *Poller) reachedPollLimit(blockNumber *big.Int) bool {
	return p.pollUntilBlock.Sign() > 0 && blockNumber.Cmp(p.pollUntilBlock) >= 0
}

func (p *Poller) getBlockRange() ([]*big.Int, error) {
	latestBlock, err := p.getLatestBlockNumber()
	if err != nil {
		return nil, err
	}
	log.Debug().Msgf("Last polled block: %s", p.lastPolledBlock.String())

	startBlock := new(big.Int).Add(p.lastPolledBlock, big.NewInt(1))
	if startBlock.Cmp(latestBlock) > 0 {
		log.Debug().Msgf("Start block %s is greater than latest block %s, skipping", startBlock, latestBlock)
		return nil, nil
	}
	endBlock := p.getEndBlockForRange(startBlock, latestBlock)
	if startBlock.Cmp(endBlock) > 0 {
		log.Debug().Msgf("Invalid range: start block %s is greater than end block %s, skipping", startBlock, endBlock)
		return nil, nil
	}

	blockCount := new(big.Int).Sub(endBlock, startBlock).Int64() + 1
	blockNumbers := make([]*big.Int, blockCount)
	for i := int64(0); i < blockCount; i++ {
		blockNumbers[i] = new(big.Int).Add(startBlock, big.NewInt(i))
	}

	return blockNumbers, nil
}

func (p *Poller) getLatestBlockNumber() (*big.Int, error) {
	latestBlockUint64, err := p.rpc.EthClient.BlockNumber(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get latest block number: %v", err)
	}
	return new(big.Int).SetUint64(latestBlockUint64), nil
}

func (p *Poller) getEndBlockForRange(startBlock *big.Int, latestBlock *big.Int) *big.Int {
	endBlock := new(big.Int).Add(startBlock, big.NewInt(p.blocksPerPoll-1))
	if endBlock.Cmp(latestBlock) > 0 {
		endBlock = latestBlock
	}
	if p.reachedPollLimit(endBlock) {
		log.Debug().Msgf("End block %s is greater than poll until block %s, setting to poll until block", endBlock, p.pollUntilBlock)
		endBlock = p.pollUntilBlock
	}
	return endBlock
}

func (p *Poller) handleWorkerResults(results []worker.WorkerResult) {
	var successfulResults []worker.WorkerResult
	var failedResults []worker.WorkerResult

	for _, result := range results {
		if result.Error != nil {
			log.Warn().Err(result.Error).Msgf("Error fetching block data for block %s", result.BlockNumber.String())
			failedResults = append(failedResults, result)
		} else {
			successfulResults = append(successfulResults, result)
		}
	}

	blockData := make([]common.BlockData, 0, len(successfulResults))
	for _, result := range successfulResults {
		blockData = append(blockData, common.BlockData{
			Block:        result.Block,
			Logs:         result.Logs,
			Transactions: result.Transactions,
			Traces:       result.Traces,
		})
	}
	if err := p.storage.StagingStorage.InsertBlockData(blockData); err != nil {
		e := fmt.Errorf("error inserting block data: %v", err)
		log.Error().Err(e)
		for _, result := range successfulResults {
			failedResults = append(failedResults, worker.WorkerResult{
				BlockNumber: result.BlockNumber,
				Error:       e,
			})
		}
		metrics.PolledBatchSize.Set(float64(len(blockData)))
	}

	if len(failedResults) > 0 {
		p.handleBlockFailures(failedResults)
	}
}

func (p *Poller) handleBlockFailures(results []worker.WorkerResult) {
	var blockFailures []common.BlockFailure
	for _, result := range results {
		if result.Error != nil {
			blockFailures = append(blockFailures, common.BlockFailure{
				BlockNumber:   result.BlockNumber,
				FailureReason: result.Error.Error(),
				FailureTime:   time.Now(),
				ChainId:       p.rpc.ChainID,
				FailureCount:  1,
			})
		}
	}
	err := p.storage.OrchestratorStorage.StoreBlockFailures(blockFailures)
	if err != nil {
		// TODO: exiting if this fails, but should handle this better
		log.Error().Err(err).Msg("Error saving block failures")
	}
}
