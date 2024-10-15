CREATE TABLE transactions (
    `chain_id` UInt256,
    `hash` FixedString(66),
    `nonce` UInt64,
    `block_hash` FixedString(66),
    `block_number` UInt256,
    `block_timestamp` UInt64 CODEC(Delta, ZSTD),
    `transaction_index` UInt64,
    `from_address` FixedString(42),
    `to_address` FixedString(42),
    `value` UInt256,
    `gas` UInt64,
    `gas_price` UInt256,
    `data` String,
    `function_selector` FixedString(10),
    `max_fee_per_gas` UInt128,
    `max_priority_fee_per_gas` UInt128,
    `transaction_type` UInt8,
    `r` UInt256,
    `s` UInt256,
    `v` UInt256,
    `access_list` Nullable(String),
    `contract_address` Nullable(FixedString(42)),
    `gas_used` Nullable(UInt64),
    `cumulative_gas_used` Nullable(UInt64),
    `effective_gas_price` Nullable(UInt256),
    `blob_gas_used` Nullable(UInt64),
    `blob_gas_price` Nullable(UInt256),
    `logs_bloom` Nullable(String),
    `status` Nullable(UInt64),
    `is_deleted` UInt8 DEFAULT 0,
    `insert_timestamp` DateTime DEFAULT now(),
    INDEX idx_block_timestamp block_timestamp TYPE minmax GRANULARITY 1,
    INDEX idx_block_hash block_hash TYPE bloom_filter GRANULARITY 1,
    INDEX idx_hash hash TYPE bloom_filter GRANULARITY 1,
    INDEX idx_from_address from_address TYPE bloom_filter GRANULARITY 1,
    INDEX idx_to_address to_address TYPE bloom_filter GRANULARITY 1,
    INDEX idx_function_selector function_selector TYPE bloom_filter GRANULARITY 1,
) ENGINE = ReplacingMergeTree(insert_timestamp, is_deleted)
ORDER BY (chain_id, block_number, hash) SETTINGS allow_experimental_replacing_merge_with_cleanup = 1;