package config

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type LogConfig struct {
	Level    string `mapstructure:"level"`
	Prettify bool   `mapstructure:"prettify"`
}

type PollerConfig struct {
	Enabled         bool `mapstructure:"enabled"`
	Interval        int  `mapstructure:"interval"`
	BlocksPerPoll   int  `mapstructure:"blocksPerPoll"`
	FromBlock       int  `mapstructure:"fromBlock"`
	ForceFromBlock  bool `mapstructure:"forceFromBlock"`
	UntilBlock      int  `mapstructure:"untilBlock"`
	ParallelPollers int  `mapstructure:"parallelPollers"`
}

type CommitterConfig struct {
	Enabled         bool `mapstructure:"enabled"`
	Interval        int  `mapstructure:"interval"`
	BlocksPerCommit int  `mapstructure:"blocksPerCommit"`
	FromBlock       int  `mapstructure:"fromBlock"`
}

type ReorgHandlerConfig struct {
	Enabled        bool `mapstructure:"enabled"`
	Interval       int  `mapstructure:"interval"`
	BlocksPerScan  int  `mapstructure:"blocksPerScan"`
	FromBlock      int  `mapstructure:"fromBlock"`
	ForceFromBlock bool `mapstructure:"forceFromBlock"`
}

type FailureRecovererConfig struct {
	Enabled      bool `mapstructure:"enabled"`
	Interval     int  `mapstructure:"interval"`
	BlocksPerRun int  `mapstructure:"blocksPerRun"`
}

type StorageConfig struct {
	Staging      StorageConnectionConfig `mapstructure:"staging"`
	Main         StorageConnectionConfig `mapstructure:"main"`
	Orchestrator StorageConnectionConfig `mapstructure:"orchestrator"`
}
type StorageType string

const (
	StorageTypeMain         StorageType = "main"
	StorageTypeStaging      StorageType = "staging"
	StorageTypeOrchestrator StorageType = "orchestrator"
)

type StorageConnectionConfig struct {
	Clickhouse *ClickhouseConfig `mapstructure:"clickhouse"`
	Memory     *MemoryConfig     `mapstructure:"memory"`
	Redis      *RedisConfig      `mapstructure:"redis"`
}

type ClickhouseConfig struct {
	Host             string `mapstructure:"host"`
	Port             int    `mapstructure:"port"`
	Username         string `mapstructure:"username"`
	Password         string `mapstructure:"password"`
	Database         string `mapstructure:"database"`
	DisableTLS       bool   `mapstructure:"disableTLS"`
	AsyncInsert      bool   `mapstructure:"asyncInsert"`
	MaxRowsPerInsert int    `mapstructure:"maxRowsPerInsert"`
}

type MemoryConfig struct {
	MaxItems int `mapstructure:"maxItems"`
}

type RedisConfig struct {
	PoolSize int    `mapstructure:"poolSize"`
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type RPCBatchRequestConfig struct {
	BlocksPerRequest int `mapstructure:"blocksPerRequest"`
	BatchDelay       int `mapstructure:"batchDelay"`
}

type ToggleableRPCBatchRequestConfig struct {
	Enabled bool `mapstructure:"enabled"`
	RPCBatchRequestConfig
}

type RPCConfig struct {
	URL           string                          `mapstructure:"url"`
	Blocks        RPCBatchRequestConfig           `mapstructure:"blocks"`
	Logs          RPCBatchRequestConfig           `mapstructure:"logs"`
	BlockReceipts ToggleableRPCBatchRequestConfig `mapstructure:"blockReceipts"`
	Traces        ToggleableRPCBatchRequestConfig `mapstructure:"traces"`
}

type BasicAuthConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type APIConfig struct {
	Host                string          `mapstructure:"host"`
	BasicAuth           BasicAuthConfig `mapstructure:"basicAuth"`
	ThirdwebContractApi string          `mapstructure:"thirdwebContractApi"`
	AbiDecodingEnabled  bool            `mapstructure:"abiDecodingEnabled"`
}

type Config struct {
	RPC              RPCConfig              `mapstructure:"rpc"`
	Log              LogConfig              `mapstructure:"log"`
	Poller           PollerConfig           `mapstructure:"poller"`
	Committer        CommitterConfig        `mapstructure:"committer"`
	FailureRecoverer FailureRecovererConfig `mapstructure:"failureRecoverer"`
	ReorgHandler     ReorgHandlerConfig     `mapstructure:"reorgHandler"`
	Storage          StorageConfig          `mapstructure:"storage"`
	API              APIConfig              `mapstructure:"api"`
}

var Cfg Config

func LoadConfig(cfgFile string) error {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
		if err := viper.ReadInConfig(); err != nil {
			return fmt.Errorf("error reading config file, %s", err)
		}
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath("./configs")

		if err := viper.ReadInConfig(); err != nil {
			log.Warn().Msgf("error reading config file, %s", err)
		}

		viper.SetConfigName("secrets")
		err := viper.MergeInConfig()
		if err != nil {
			log.Warn().Msgf("error loading secrets file: %v", err)
		}
	}

	// sets e.g. RPC_URL to rpc.url
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	viper.AutomaticEnv()

	err := viper.Unmarshal(&Cfg)
	if err != nil {
		return fmt.Errorf("error unmarshalling config: %v", err)
	}

	return nil
}
