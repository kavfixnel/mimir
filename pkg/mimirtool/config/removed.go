// SPDX-License-Identifier: AGPL-3.0-only

package config

// YAML Paths for config options removed since Cortex 1.11.0.
//nolint // not used right now, but will be soon.
var removedConfigPaths = []string{
	"flusher.concurrent_flushes",                            // -flusher.concurrent-flushes
	"flusher.flush_op_timeout",                              // -flusher.flush-op-timeout
	"flusher.wal_dir",                                       // -flusher.wal-dir
	"ingester.chunk_age_jitter",                             // -ingester.chunk-age-jitter
	"ingester.concurrent_flushes",                           // -ingester.concurrent-flushes
	"ingester.flush_op_timeout",                             // -ingester.flush-op-timeout
	"ingester.flush_period",                                 // -ingester.flush-period
	"ingester.max_chunk_age",                                // -ingester.max-chunk-age
	"ingester.max_chunk_idle_time",                          // -ingester.max-chunk-idle
	"ingester.max_stale_chunk_idle_time",                    // -ingester.max-stale-chunk-idle
	"ingester.max_transfer_retries",                         // -ingester.max-transfer-retries
	"ingester.retain_period",                                // -ingester.retain-period
	"ingester.spread_flushes",                               // -ingester.spread-flushes
	"ingester.walconfig.checkpoint_duration",                // -ingester.checkpoint-duration
	"ingester.walconfig.checkpoint_enabled",                 // -ingester.checkpoint-enabled
	"ingester.walconfig.flush_on_shutdown_with_wal_enabled", // -ingester.flush-on-shutdown-with-wal-enabled
	"ingester.walconfig.recover_from_wal",                   // -ingester.recover-from-wal
	"ingester.walconfig.wal_dir",                            // -ingester.wal-dir
	"ingester.walconfig.wal_enabled",                        // -ingester.wal-enabled
	"limits.max_series_per_query",                           // -ingester.max-series-per-query
	"limits.min_chunk_length",                               // -ingester.min-chunk-length
	"querier.second_store_engine",                           // -querier.second-store-engine
	"querier.use_second_store_before_time",                  // -querier.use-second-store-before-time
	"storage.engine",                                        // -store.engine

	// All table-manager flags.
	"table_manager.chunk_tables_provisioning.enable_inactive_throughput_on_demand_mode", // -table-manager.chunk-table.inactive-enable-ondemand-throughput-mode
	"table_manager.chunk_tables_provisioning.enable_ondemand_throughput_mode",           // -table-manager.chunk-table.enable-ondemand-throughput-mode
	"table_manager.chunk_tables_provisioning.inactive_read_scale.enabled",               // -table-manager.chunk-table.inactive-read-throughput.scale.enabled
	"table_manager.chunk_tables_provisioning.inactive_read_scale.in_cooldown",           // -table-manager.chunk-table.inactive-read-throughput.scale.in-cooldown
	"table_manager.chunk_tables_provisioning.inactive_read_scale.max_capacity",          // -table-manager.chunk-table.inactive-read-throughput.scale.max-capacity
	"table_manager.chunk_tables_provisioning.inactive_read_scale.min_capacity",          // -table-manager.chunk-table.inactive-read-throughput.scale.min-capacity
	"table_manager.chunk_tables_provisioning.inactive_read_scale.out_cooldown",          // -table-manager.chunk-table.inactive-read-throughput.scale.out-cooldown
	"table_manager.chunk_tables_provisioning.inactive_read_scale.role_arn",              // -table-manager.chunk-table.inactive-read-throughput.scale.role-arn
	"table_manager.chunk_tables_provisioning.inactive_read_scale.target",                // -table-manager.chunk-table.inactive-read-throughput.scale.target-value
	"table_manager.chunk_tables_provisioning.inactive_read_scale_lastn",                 // -table-manager.chunk-table.inactive-read-throughput.scale-last-n
	"table_manager.chunk_tables_provisioning.inactive_read_throughput",                  // -table-manager.chunk-table.inactive-read-throughput
	"table_manager.chunk_tables_provisioning.inactive_write_scale.enabled",              // -table-manager.chunk-table.inactive-write-throughput.scale.enabled
	"table_manager.chunk_tables_provisioning.inactive_write_scale.in_cooldown",          // -table-manager.chunk-table.inactive-write-throughput.scale.in-cooldown
	"table_manager.chunk_tables_provisioning.inactive_write_scale.max_capacity",         // -table-manager.chunk-table.inactive-write-throughput.scale.max-capacity
	"table_manager.chunk_tables_provisioning.inactive_write_scale.min_capacity",         // -table-manager.chunk-table.inactive-write-throughput.scale.min-capacity
	"table_manager.chunk_tables_provisioning.inactive_write_scale.out_cooldown",         // -table-manager.chunk-table.inactive-write-throughput.scale.out-cooldown
	"table_manager.chunk_tables_provisioning.inactive_write_scale.role_arn",             // -table-manager.chunk-table.inactive-write-throughput.scale.role-arn
	"table_manager.chunk_tables_provisioning.inactive_write_scale.target",               // -table-manager.chunk-table.inactive-write-throughput.scale.target-value
	"table_manager.chunk_tables_provisioning.inactive_write_scale_lastn",                // -table-manager.chunk-table.inactive-write-throughput.scale-last-n
	"table_manager.chunk_tables_provisioning.inactive_write_throughput",                 // -table-manager.chunk-table.inactive-write-throughput
	"table_manager.chunk_tables_provisioning.provisioned_read_throughput",               // -table-manager.chunk-table.read-throughput
	"table_manager.chunk_tables_provisioning.provisioned_write_throughput",              // -table-manager.chunk-table.write-throughput
	"table_manager.chunk_tables_provisioning.read_scale.enabled",                        // -table-manager.chunk-table.read-throughput.scale.enabled
	"table_manager.chunk_tables_provisioning.read_scale.in_cooldown",                    // -table-manager.chunk-table.read-throughput.scale.in-cooldown
	"table_manager.chunk_tables_provisioning.read_scale.max_capacity",                   // -table-manager.chunk-table.read-throughput.scale.max-capacity
	"table_manager.chunk_tables_provisioning.read_scale.min_capacity",                   // -table-manager.chunk-table.read-throughput.scale.min-capacity
	"table_manager.chunk_tables_provisioning.read_scale.out_cooldown",                   // -table-manager.chunk-table.read-throughput.scale.out-cooldown
	"table_manager.chunk_tables_provisioning.read_scale.role_arn",                       // -table-manager.chunk-table.read-throughput.scale.role-arn
	"table_manager.chunk_tables_provisioning.read_scale.target",                         // -table-manager.chunk-table.read-throughput.scale.target-value
	"table_manager.chunk_tables_provisioning.write_scale.enabled",                       // -table-manager.chunk-table.write-throughput.scale.enabled
	"table_manager.chunk_tables_provisioning.write_scale.in_cooldown",                   // -table-manager.chunk-table.write-throughput.scale.in-cooldown
	"table_manager.chunk_tables_provisioning.write_scale.max_capacity",                  // -table-manager.chunk-table.write-throughput.scale.max-capacity
	"table_manager.chunk_tables_provisioning.write_scale.min_capacity",                  // -table-manager.chunk-table.write-throughput.scale.min-capacity
	"table_manager.chunk_tables_provisioning.write_scale.out_cooldown",                  // -table-manager.chunk-table.write-throughput.scale.out-cooldown
	"table_manager.chunk_tables_provisioning.write_scale.role_arn",                      // -table-manager.chunk-table.write-throughput.scale.role-arn
	"table_manager.chunk_tables_provisioning.write_scale.target",                        // -table-manager.chunk-table.write-throughput.scale.target-value
	"table_manager.creation_grace_period",                                               // -table-manager.periodic-table.grace-period
	"table_manager.index_tables_provisioning.enable_inactive_throughput_on_demand_mode", // -table-manager.index-table.inactive-enable-ondemand-throughput-mode
	"table_manager.index_tables_provisioning.enable_ondemand_throughput_mode",           // -table-manager.index-table.enable-ondemand-throughput-mode
	"table_manager.index_tables_provisioning.inactive_read_scale.enabled",               // -table-manager.index-table.inactive-read-throughput.scale.enabled
	"table_manager.index_tables_provisioning.inactive_read_scale.in_cooldown",           // -table-manager.index-table.inactive-read-throughput.scale.in-cooldown
	"table_manager.index_tables_provisioning.inactive_read_scale.max_capacity",          // -table-manager.index-table.inactive-read-throughput.scale.max-capacity
	"table_manager.index_tables_provisioning.inactive_read_scale.min_capacity",          // -table-manager.index-table.inactive-read-throughput.scale.min-capacity
	"table_manager.index_tables_provisioning.inactive_read_scale.out_cooldown",          // -table-manager.index-table.inactive-read-throughput.scale.out-cooldown
	"table_manager.index_tables_provisioning.inactive_read_scale.role_arn",              // -table-manager.index-table.inactive-read-throughput.scale.role-arn
	"table_manager.index_tables_provisioning.inactive_read_scale.target",                // -table-manager.index-table.inactive-read-throughput.scale.target-value
	"table_manager.index_tables_provisioning.inactive_read_scale_lastn",                 // -table-manager.index-table.inactive-read-throughput.scale-last-n
	"table_manager.index_tables_provisioning.inactive_read_throughput",                  // -table-manager.index-table.inactive-read-throughput
	"table_manager.index_tables_provisioning.inactive_write_scale.enabled",              // -table-manager.index-table.inactive-write-throughput.scale.enabled
	"table_manager.index_tables_provisioning.inactive_write_scale.in_cooldown",          // -table-manager.index-table.inactive-write-throughput.scale.in-cooldown
	"table_manager.index_tables_provisioning.inactive_write_scale.max_capacity",         // -table-manager.index-table.inactive-write-throughput.scale.max-capacity
	"table_manager.index_tables_provisioning.inactive_write_scale.min_capacity",         // -table-manager.index-table.inactive-write-throughput.scale.min-capacity
	"table_manager.index_tables_provisioning.inactive_write_scale.out_cooldown",         // -table-manager.index-table.inactive-write-throughput.scale.out-cooldown
	"table_manager.index_tables_provisioning.inactive_write_scale.role_arn",             // -table-manager.index-table.inactive-write-throughput.scale.role-arn
	"table_manager.index_tables_provisioning.inactive_write_scale.target",               // -table-manager.index-table.inactive-write-throughput.scale.target-value
	"table_manager.index_tables_provisioning.inactive_write_scale_lastn",                // -table-manager.index-table.inactive-write-throughput.scale-last-n
	"table_manager.index_tables_provisioning.inactive_write_throughput",                 // -table-manager.index-table.inactive-write-throughput
	"table_manager.index_tables_provisioning.provisioned_read_throughput",               // -table-manager.index-table.read-throughput
	"table_manager.index_tables_provisioning.provisioned_write_throughput",              // -table-manager.index-table.write-throughput
	"table_manager.index_tables_provisioning.read_scale.enabled",                        // -table-manager.index-table.read-throughput.scale.enabled
	"table_manager.index_tables_provisioning.read_scale.in_cooldown",                    // -table-manager.index-table.read-throughput.scale.in-cooldown
	"table_manager.index_tables_provisioning.read_scale.max_capacity",                   // -table-manager.index-table.read-throughput.scale.max-capacity
	"table_manager.index_tables_provisioning.read_scale.min_capacity",                   // -table-manager.index-table.read-throughput.scale.min-capacity
	"table_manager.index_tables_provisioning.read_scale.out_cooldown",                   // -table-manager.index-table.read-throughput.scale.out-cooldown
	"table_manager.index_tables_provisioning.read_scale.role_arn",                       // -table-manager.index-table.read-throughput.scale.role-arn
	"table_manager.index_tables_provisioning.read_scale.target",                         // -table-manager.index-table.read-throughput.scale.target-value
	"table_manager.index_tables_provisioning.write_scale.enabled",                       // -table-manager.index-table.write-throughput.scale.enabled
	"table_manager.index_tables_provisioning.write_scale.in_cooldown",                   // -table-manager.index-table.write-throughput.scale.in-cooldown
	"table_manager.index_tables_provisioning.write_scale.max_capacity",                  // -table-manager.index-table.write-throughput.scale.max-capacity
	"table_manager.index_tables_provisioning.write_scale.min_capacity",                  // -table-manager.index-table.write-throughput.scale.min-capacity
	"table_manager.index_tables_provisioning.write_scale.out_cooldown",                  // -table-manager.index-table.write-throughput.scale.out-cooldown
	"table_manager.index_tables_provisioning.write_scale.role_arn",                      // -table-manager.index-table.write-throughput.scale.role-arn
	"table_manager.index_tables_provisioning.write_scale.target",                        // -table-manager.index-table.write-throughput.scale.target-value
	"table_manager.poll_interval",                                                       // -table-manager.poll-interval
	"table_manager.retention_deletes_enabled",                                           // -table-manager.retention-deletes-enabled
	"table_manager.retention_period",                                                    // -table-manager.retention-period
	"table_manager.throughput_updates_disabled",                                         // -table-manager.throughput-updates-disabled

	// All `-deletes.*` flags
	"storage.delete_store.requests_table_name",                                // -deletes.requests-table-name
	"storage.delete_store.store",                                              // -deletes.store
	"storage.delete_store.table_provisioning.enable_ondemand_throughput_mode", // -deletes.table.enable-ondemand-throughput-mode
	"storage.delete_store.table_provisioning.provisioned_read_throughput",     // -deletes.table.read-throughput
	"storage.delete_store.table_provisioning.provisioned_write_throughput",    // -deletes.table.write-throughput
	"storage.delete_store.table_provisioning.read_scale.enabled",              // -deletes.table.read-throughput.scale.enabled
	"storage.delete_store.table_provisioning.read_scale.in_cooldown",          // -deletes.table.read-throughput.scale.in-cooldown
	"storage.delete_store.table_provisioning.read_scale.max_capacity",         // -deletes.table.read-throughput.scale.max-capacity
	"storage.delete_store.table_provisioning.read_scale.min_capacity",         // -deletes.table.read-throughput.scale.min-capacity
	"storage.delete_store.table_provisioning.read_scale.out_cooldown",         // -deletes.table.read-throughput.scale.out-cooldown
	"storage.delete_store.table_provisioning.read_scale.role_arn",             // -deletes.table.read-throughput.scale.role-arn
	"storage.delete_store.table_provisioning.read_scale.target",               // -deletes.table.read-throughput.scale.target-value
	"storage.delete_store.table_provisioning.tags",                            // -deletes.table.tags
	"storage.delete_store.table_provisioning.write_scale.enabled",             // -deletes.table.write-throughput.scale.enabled
	"storage.delete_store.table_provisioning.write_scale.in_cooldown",         // -deletes.table.write-throughput.scale.in-cooldown
	"storage.delete_store.table_provisioning.write_scale.max_capacity",        // -deletes.table.write-throughput.scale.max-capacity
	"storage.delete_store.table_provisioning.write_scale.min_capacity",        // -deletes.table.write-throughput.scale.min-capacity
	"storage.delete_store.table_provisioning.write_scale.out_cooldown",        // -deletes.table.write-throughput.scale.out-cooldown
	"storage.delete_store.table_provisioning.write_scale.role_arn",            // -deletes.table.write-throughput.scale.role-arn
	"storage.delete_store.table_provisioning.write_scale.target",              // -deletes.table.write-throughput.scale.target-value

	// All `-purger.*` flags
	"purger.delete_request_cancel_period", // -purger.delete-request-cancel-period
	"purger.enable",                       // -purger.enable
	"purger.num_workers",                  // -purger.num-workers
	"purger.object_store_type",            // -purger.object-store-type

	// All `-metrics.*` flags
	"storage.aws.dynamodb.metrics.ignore_throttle_below", // -metrics.ignore-throttle-below
	"storage.aws.dynamodb.metrics.queue_length_query",    // -metrics.queue-length-query
	"storage.aws.dynamodb.metrics.read_error_query",      // -metrics.read-error-query
	"storage.aws.dynamodb.metrics.read_usage_query",      // -metrics.read-usage-query
	"storage.aws.dynamodb.metrics.scale_up_factor",       // -metrics.scale-up-factor
	"storage.aws.dynamodb.metrics.target_queue_length",   // -metrics.target-queue-length
	"storage.aws.dynamodb.metrics.url",                   // -metrics.url
	"storage.aws.dynamodb.metrics.write_throttle_query",  // -metrics.write-throttle-query
	"storage.aws.dynamodb.metrics.write_usage_query",     // -metrics.usage-query

	// All `-dynamodb.*` flags
	"storage.aws.dynamodb.api_limit",                  // -dynamodb.api-limit
	"storage.aws.dynamodb.backoff_config.max_period",  // -dynamodb.max-backoff
	"storage.aws.dynamodb.backoff_config.max_retries", // -dynamodb.max-retries
	"storage.aws.dynamodb.backoff_config.min_period",  // -dynamodb.min-backoff
	"storage.aws.dynamodb.chunk_gang_size",            // -dynamodb.chunk-gang-size
	"storage.aws.dynamodb.chunk_get_max_parallelism",  // -dynamodb.chunk.get-max-parallelism
	"storage.aws.dynamodb.dynamodb_url",               // -dynamodb.url
	"storage.aws.dynamodb.throttle_limit",             // -dynamodb.throttle-limit

	// All `-s3.*` flags
	"storage.aws.access_key_id",                       // -s3.access-key-id
	"storage.aws.bucketnames",                         // -s3.buckets
	"storage.aws.endpoint",                            // -s3.endpoint
	"storage.aws.http_config.idle_conn_timeout",       // -s3.http.idle-conn-timeout
	"storage.aws.http_config.insecure_skip_verify",    // -s3.http.insecure-skip-verify
	"storage.aws.http_config.response_header_timeout", // -s3.http.response-header-timeout
	"storage.aws.insecure",                            // -s3.insecure
	"storage.aws.region",                              // -s3.region
	"storage.aws.s3",                                  // -s3.url
	"storage.aws.s3forcepathstyle",                    // -s3.force-path-style
	"storage.aws.secret_access_key",                   // -s3.secret-access-key
	"storage.aws.signature_version",                   // -s3.signature-version
	"storage.aws.sse.kms_encryption_context",          // -s3.sse.kms-encryption-context
	"storage.aws.sse.kms_key_id",                      // -s3.sse.kms-key-id
	"storage.aws.sse.type",                            // -s3.sse.type
	"storage.aws.sse_encryption",                      // -s3.sse-encryption

	// All `-azure.*` flags
	"storage.azure.account_key",          // -azure.account-key
	"storage.azure.account_name",         // -azure.account-name
	"storage.azure.container_name",       // -azure.container-name
	"storage.azure.download_buffer_size", // -azure.download-buffer-size
	"storage.azure.environment",          // -azure.environment
	"storage.azure.max_retries",          // -azure.max-retries
	"storage.azure.max_retry_delay",      // -azure.max-retry-delay
	"storage.azure.min_retry_delay",      // -azure.min-retry-delay
	"storage.azure.request_timeout",      // -azure.request-timeout
	"storage.azure.upload_buffer_count",  // -azure.download-buffer-count
	"storage.azure.upload_buffer_size",   // -azure.upload-buffer-size

	// All `-bigtable.*` flags
	"storage.bigtable.grpc_client_config.backoff_config.max_period",  // -bigtable.backoff-max-period
	"storage.bigtable.grpc_client_config.backoff_config.max_retries", // -bigtable.backoff-retries
	"storage.bigtable.grpc_client_config.backoff_config.min_period",  // -bigtable.backoff-min-period
	"storage.bigtable.grpc_client_config.backoff_on_ratelimits",      // -bigtable.backoff-on-ratelimits
	"storage.bigtable.grpc_client_config.grpc_compression",           // -bigtable.grpc-compression
	"storage.bigtable.grpc_client_config.max_recv_msg_size",          // -bigtable.grpc-max-recv-msg-size
	"storage.bigtable.grpc_client_config.max_send_msg_size",          // -bigtable.grpc-max-send-msg-size
	"storage.bigtable.grpc_client_config.rate_limit",                 // -bigtable.grpc-client-rate-limit
	"storage.bigtable.grpc_client_config.rate_limit_burst",           // -bigtable.grpc-client-rate-limit-burst
	"storage.bigtable.grpc_client_config.tls_ca_path",                // -bigtable.tls-ca-path
	"storage.bigtable.grpc_client_config.tls_cert_path",              // -bigtable.tls-cert-path
	"storage.bigtable.grpc_client_config.tls_enabled",                // -bigtable.tls-enabled
	"storage.bigtable.grpc_client_config.tls_insecure_skip_verify",   // -bigtable.tls-insecure-skip-verify
	"storage.bigtable.grpc_client_config.tls_key_path",               // -bigtable.tls-key-path
	"storage.bigtable.grpc_client_config.tls_server_name",            // -bigtable.tls-server-name
	"storage.bigtable.instance",                                      // -bigtable.instance
	"storage.bigtable.project",                                       // -bigtable.project
	"storage.bigtable.table_cache_enabled",                           // -bigtable.table-cache.enabled
	"storage.bigtable.table_cache_expiration",                        // -bigtable.table-cache.expiration

	// All `-gcs.*` flags
	"storage.gcs.bucket_name",       // -gcs.bucketname
	"storage.gcs.chunk_buffer_size", // -gcs.chunk-buffer-size
	"storage.gcs.enable_opencensus", // -gcs.enable-opencensus
	"storage.gcs.request_timeout",   // -gcs.request-timeout

	// All `-cassandra.*` flags
	"storage.cassandra.CA_path",                     // -cassandra.ca-path
	"storage.cassandra.SSL",                         // -cassandra.ssl
	"storage.cassandra.addresses",                   // -cassandra.addresses
	"storage.cassandra.auth",                        // -cassandra.auth
	"storage.cassandra.connect_timeout",             // -cassandra.connect-timeout
	"storage.cassandra.consistency",                 // -cassandra.consistency
	"storage.cassandra.convict_hosts_on_failure",    // -cassandra.convict-hosts-on-failure
	"storage.cassandra.custom_authenticators",       // -cassandra.custom-authenticator
	"storage.cassandra.disable_initial_host_lookup", // -cassandra.disable-initial-host-lookup
	"storage.cassandra.host_selection_policy",       // -cassandra.host-selection-policy
	"storage.cassandra.host_verification",           // -cassandra.host-verification
	"storage.cassandra.keyspace",                    // -cassandra.keyspace
	"storage.cassandra.max_retries",                 // -cassandra.max-retries
	"storage.cassandra.num_connections",             // -cassandra.num-connections
	"storage.cassandra.password",                    // -cassandra.password
	"storage.cassandra.password_file",               // -cassandra.password-file
	"storage.cassandra.port",                        // -cassandra.port
	"storage.cassandra.query_concurrency",           // -cassandra.query-concurrency
	"storage.cassandra.reconnect_interval",          // -cassandra.reconnent-interval
	"storage.cassandra.replication_factor",          // -cassandra.replication-factor
	"storage.cassandra.retry_max_backoff",           // -cassandra.retry-max-backoff
	"storage.cassandra.retry_min_backoff",           // -cassandra.retry-min-backoff
	"storage.cassandra.table_options",               // -cassandra.table-options
	"storage.cassandra.timeout",                     // -cassandra.timeout
	"storage.cassandra.tls_cert_path",               // -cassandra.tls-cert-path
	"storage.cassandra.tls_key_path",                // -cassandra.tls-key-path
	"storage.cassandra.username",                    // -cassandra.username

	// All `-boltdb.*` flags
	"storage.boltdb.directory", // -boltdb.dir

	// All `-local.*` flags
	"storage.filesystem.directory", // -local.chunk-directory

	// All `-swift.*` flags
	"storage.swift.auth_url",            // -swift.auth-url
	"storage.swift.auth_version",        // -swift.auth-version
	"storage.swift.connect_timeout",     // -swift.connect-timeout
	"storage.swift.container_name",      // -swift.container-name
	"storage.swift.domain_id",           // -swift.domain-id
	"storage.swift.domain_name",         // -swift.domain-name
	"storage.swift.max_retries",         // -swift.max-retries
	"storage.swift.password",            // -swift.password
	"storage.swift.project_domain_id",   // -swift.project-domain-id
	"storage.swift.project_domain_name", // -swift.project-domain-name
	"storage.swift.project_id",          // -swift.project-id
	"storage.swift.project_name",        // -swift.project-name
	"storage.swift.region_name",         // -swift.region-name
	"storage.swift.request_timeout",     // -swift.request-timeout
	"storage.swift.user_domain_id",      // -swift.user-domain-id
	"storage.swift.user_domain_name",    // -swift.user-domain-name
	"storage.swift.user_id",             // -swift.user-id
	"storage.swift.username",            // -swift.username

	// All `-store.*` flags except `-store.engine`, `-store.max-query-length`, `-store.max-labels-query-length`
	"chunk_store.cache_lookups_older_than",                                                        // -store.cache-lookups-older-than
	"chunk_store.chunk_cache_config.background.writeback_buffer",                                  // -store.chunks-cache.background.write-back-buffer
	"chunk_store.chunk_cache_config.background.writeback_goroutines",                              // -store.chunks-cache.background.write-back-concurrency
	"chunk_store.chunk_cache_config.default_validity",                                             // -store.chunks-cache.default-validity
	"chunk_store.chunk_cache_config.enable_fifocache",                                             // -store.chunks-cache.cache.enable-fifocache
	"chunk_store.chunk_cache_config.fifocache.max_size_bytes",                                     // -store.chunks-cache.fifocache.max-size-bytes
	"chunk_store.chunk_cache_config.fifocache.max_size_items",                                     // -store.chunks-cache.fifocache.max-size-items
	"chunk_store.chunk_cache_config.fifocache.size",                                               // -store.chunks-cache.fifocache.size
	"chunk_store.chunk_cache_config.fifocache.validity",                                           // -store.chunks-cache.fifocache.duration
	"chunk_store.chunk_cache_config.memcached.batch_size",                                         // -store.chunks-cache.memcached.batchsize
	"chunk_store.chunk_cache_config.memcached.expiration",                                         // -store.chunks-cache.memcached.expiration
	"chunk_store.chunk_cache_config.memcached.parallelism",                                        // -store.chunks-cache.memcached.parallelism
	"chunk_store.chunk_cache_config.memcached_client.addresses",                                   // -store.chunks-cache.memcached.addresses
	"chunk_store.chunk_cache_config.memcached_client.circuit_breaker_consecutive_failures",        // -store.chunks-cache.memcached.circuit-breaker-consecutive-failures
	"chunk_store.chunk_cache_config.memcached_client.circuit_breaker_interval",                    // -store.chunks-cache.memcached.circuit-breaker-interval
	"chunk_store.chunk_cache_config.memcached_client.circuit_breaker_timeout",                     // -store.chunks-cache.memcached.circuit-breaker-timeout
	"chunk_store.chunk_cache_config.memcached_client.consistent_hash",                             // -store.chunks-cache.memcached.consistent-hash
	"chunk_store.chunk_cache_config.memcached_client.host",                                        // -store.chunks-cache.memcached.hostname
	"chunk_store.chunk_cache_config.memcached_client.max_idle_conns",                              // -store.chunks-cache.memcached.max-idle-conns
	"chunk_store.chunk_cache_config.memcached_client.max_item_size",                               // -store.chunks-cache.memcached.max-item-size
	"chunk_store.chunk_cache_config.memcached_client.service",                                     // -store.chunks-cache.memcached.service
	"chunk_store.chunk_cache_config.memcached_client.timeout",                                     // -store.chunks-cache.memcached.timeout
	"chunk_store.chunk_cache_config.memcached_client.update_interval",                             // -store.chunks-cache.memcached.update-interval
	"chunk_store.chunk_cache_config.redis.db",                                                     // -store.chunks-cache.redis.db
	"chunk_store.chunk_cache_config.redis.endpoint",                                               // -store.chunks-cache.redis.endpoint
	"chunk_store.chunk_cache_config.redis.expiration",                                             // -store.chunks-cache.redis.expiration
	"chunk_store.chunk_cache_config.redis.idle_timeout",                                           // -store.chunks-cache.redis.idle-timeout
	"chunk_store.chunk_cache_config.redis.master_name",                                            // -store.chunks-cache.redis.master-name
	"chunk_store.chunk_cache_config.redis.max_connection_age",                                     // -store.chunks-cache.redis.max-connection-age
	"chunk_store.chunk_cache_config.redis.password",                                               // -store.chunks-cache.redis.password
	"chunk_store.chunk_cache_config.redis.pool_size",                                              // -store.chunks-cache.redis.pool-size
	"chunk_store.chunk_cache_config.redis.timeout",                                                // -store.chunks-cache.redis.timeout
	"chunk_store.chunk_cache_config.redis.tls_enabled",                                            // -store.chunks-cache.redis.tls-enabled
	"chunk_store.chunk_cache_config.redis.tls_insecure_skip_verify",                               // -store.chunks-cache.redis.tls-insecure-skip-verify
	"chunk_store.write_dedupe_cache_config.background.writeback_buffer",                           // -store.index-cache-write.background.write-back-buffer
	"chunk_store.write_dedupe_cache_config.background.writeback_goroutines",                       // -store.index-cache-write.background.write-back-concurrency
	"chunk_store.write_dedupe_cache_config.default_validity",                                      // -store.index-cache-write.default-validity
	"chunk_store.write_dedupe_cache_config.enable_fifocache",                                      // -store.index-cache-write.cache.enable-fifocache
	"chunk_store.write_dedupe_cache_config.fifocache.max_size_bytes",                              // -store.index-cache-write.fifocache.max-size-bytes
	"chunk_store.write_dedupe_cache_config.fifocache.max_size_items",                              // -store.index-cache-write.fifocache.max-size-items
	"chunk_store.write_dedupe_cache_config.fifocache.size",                                        // -store.index-cache-write.fifocache.size
	"chunk_store.write_dedupe_cache_config.fifocache.validity",                                    // -store.index-cache-write.fifocache.duration
	"chunk_store.write_dedupe_cache_config.memcached.batch_size",                                  // -store.index-cache-write.memcached.batchsize
	"chunk_store.write_dedupe_cache_config.memcached.expiration",                                  // -store.index-cache-write.memcached.expiration
	"chunk_store.write_dedupe_cache_config.memcached.parallelism",                                 // -store.index-cache-write.memcached.parallelism
	"chunk_store.write_dedupe_cache_config.memcached_client.addresses",                            // -store.index-cache-write.memcached.addresses
	"chunk_store.write_dedupe_cache_config.memcached_client.circuit_breaker_consecutive_failures", // -store.index-cache-write.memcached.circuit-breaker-consecutive-failures
	"chunk_store.write_dedupe_cache_config.memcached_client.circuit_breaker_interval",             // -store.index-cache-write.memcached.circuit-breaker-interval
	"chunk_store.write_dedupe_cache_config.memcached_client.circuit_breaker_timeout",              // -store.index-cache-write.memcached.circuit-breaker-timeout
	"chunk_store.write_dedupe_cache_config.memcached_client.consistent_hash",                      // -store.index-cache-write.memcached.consistent-hash
	"chunk_store.write_dedupe_cache_config.memcached_client.host",                                 // -store.index-cache-write.memcached.hostname
	"chunk_store.write_dedupe_cache_config.memcached_client.max_idle_conns",                       // -store.index-cache-write.memcached.max-idle-conns
	"chunk_store.write_dedupe_cache_config.memcached_client.max_item_size",                        // -store.index-cache-write.memcached.max-item-size
	"chunk_store.write_dedupe_cache_config.memcached_client.service",                              // -store.index-cache-write.memcached.service
	"chunk_store.write_dedupe_cache_config.memcached_client.timeout",                              // -store.index-cache-write.memcached.timeout
	"chunk_store.write_dedupe_cache_config.memcached_client.update_interval",                      // -store.index-cache-write.memcached.update-interval
	"chunk_store.write_dedupe_cache_config.redis.db",                                              // -store.index-cache-write.redis.db
	"chunk_store.write_dedupe_cache_config.redis.endpoint",                                        // -store.index-cache-write.redis.endpoint
	"chunk_store.write_dedupe_cache_config.redis.expiration",                                      // -store.index-cache-write.redis.expiration
	"chunk_store.write_dedupe_cache_config.redis.idle_timeout",                                    // -store.index-cache-write.redis.idle-timeout
	"chunk_store.write_dedupe_cache_config.redis.master_name",                                     // -store.index-cache-write.redis.master-name
	"chunk_store.write_dedupe_cache_config.redis.max_connection_age",                              // -store.index-cache-write.redis.max-connection-age
	"chunk_store.write_dedupe_cache_config.redis.password",                                        // -store.index-cache-write.redis.password
	"chunk_store.write_dedupe_cache_config.redis.pool_size",                                       // -store.index-cache-write.redis.pool-size
	"chunk_store.write_dedupe_cache_config.redis.timeout",                                         // -store.index-cache-write.redis.timeout
	"chunk_store.write_dedupe_cache_config.redis.tls_enabled",                                     // -store.index-cache-write.redis.tls-enabled
	"chunk_store.write_dedupe_cache_config.redis.tls_insecure_skip_verify",                        // -store.index-cache-write.redis.tls-insecure-skip-verify
	"limits.cardinality_limit",                                                                    // -store.cardinality-limit
	"limits.max_chunks_per_query",                                                                 // -store.query-chunk-limit
	"storage.index_cache_validity",                                                                // -store.index-cache-validity
	"storage.index_queries_cache_config.background.writeback_buffer",                              // -store.index-cache-read.background.write-back-buffer
	"storage.index_queries_cache_config.background.writeback_goroutines",                          // -store.index-cache-read.background.write-back-concurrency
	"storage.index_queries_cache_config.default_validity",                                         // -store.index-cache-read.default-validity
	"storage.index_queries_cache_config.enable_fifocache",                                         // -store.index-cache-read.cache.enable-fifocache
	"storage.index_queries_cache_config.fifocache.max_size_bytes",                                 // -store.index-cache-read.fifocache.max-size-bytes
	"storage.index_queries_cache_config.fifocache.max_size_items",                                 // -store.index-cache-read.fifocache.max-size-items
	"storage.index_queries_cache_config.fifocache.size",                                           // -store.index-cache-read.fifocache.size
	"storage.index_queries_cache_config.fifocache.validity",                                       // -store.index-cache-read.fifocache.duration
	"storage.index_queries_cache_config.memcached.batch_size",                                     // -store.index-cache-read.memcached.batchsize
	"storage.index_queries_cache_config.memcached.expiration",                                     // -store.index-cache-read.memcached.expiration
	"storage.index_queries_cache_config.memcached.parallelism",                                    // -store.index-cache-read.memcached.parallelism
	"storage.index_queries_cache_config.memcached_client.addresses",                               // -store.index-cache-read.memcached.addresses
	"storage.index_queries_cache_config.memcached_client.circuit_breaker_consecutive_failures",    // -store.index-cache-read.memcached.circuit-breaker-consecutive-failures
	"storage.index_queries_cache_config.memcached_client.circuit_breaker_interval",                // -store.index-cache-read.memcached.circuit-breaker-interval
	"storage.index_queries_cache_config.memcached_client.circuit_breaker_timeout",                 // -store.index-cache-read.memcached.circuit-breaker-timeout
	"storage.index_queries_cache_config.memcached_client.consistent_hash",                         // -store.index-cache-read.memcached.consistent-hash
	"storage.index_queries_cache_config.memcached_client.host",                                    // -store.index-cache-read.memcached.hostname
	"storage.index_queries_cache_config.memcached_client.max_idle_conns",                          // -store.index-cache-read.memcached.max-idle-conns
	"storage.index_queries_cache_config.memcached_client.max_item_size",                           // -store.index-cache-read.memcached.max-item-size
	"storage.index_queries_cache_config.memcached_client.service",                                 // -store.index-cache-read.memcached.service
	"storage.index_queries_cache_config.memcached_client.timeout",                                 // -store.index-cache-read.memcached.timeout
	"storage.index_queries_cache_config.memcached_client.update_interval",                         // -store.index-cache-read.memcached.update-interval
	"storage.index_queries_cache_config.redis.db",                                                 // -store.index-cache-read.redis.db
	"storage.index_queries_cache_config.redis.endpoint",                                           // -store.index-cache-read.redis.endpoint
	"storage.index_queries_cache_config.redis.expiration",                                         // -store.index-cache-read.redis.expiration
	"storage.index_queries_cache_config.redis.idle_timeout",                                       // -store.index-cache-read.redis.idle-timeout
	"storage.index_queries_cache_config.redis.master_name",                                        // -store.index-cache-read.redis.master-name
	"storage.index_queries_cache_config.redis.max_connection_age",                                 // -store.index-cache-read.redis.max-connection-age
	"storage.index_queries_cache_config.redis.password",                                           // -store.index-cache-read.redis.password
	"storage.index_queries_cache_config.redis.pool_size",                                          // -store.index-cache-read.redis.pool-size
	"storage.index_queries_cache_config.redis.timeout",                                            // -store.index-cache-read.redis.timeout
	"storage.index_queries_cache_config.redis.tls_enabled",                                        // -store.index-cache-read.redis.tls-enabled
	"storage.index_queries_cache_config.redis.tls_insecure_skip_verify",                           // -store.index-cache-read.redis.tls-insecure-skip-verify

	// All `-grpc-store.*` flags
	"storage.grpc_store.server_address", // -grpc-store.server-address
}

// CLI options removed since Cortex 1.11.0. These flags only existed as CLI Flags, and were not included in YAML Config.
//nolint // not used right now, but will be soon.
var removedCLIOptions = []string{
	"schema-config-file",
	"ingester.chunk-encoding",
	"querier.query-parallelism",
}
