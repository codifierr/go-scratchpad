package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// Config holds all application configuration
type Config struct {
	Server        ServerConfig
	Elasticsearch ElasticsearchConfig
	Buffer        BufferConfig
	Retry         RetryConfig
}

// ServerConfig holds HTTP server configuration
type ServerConfig struct {
	Port string
}

// ElasticsearchConfig holds Elasticsearch connection configuration
type ElasticsearchConfig struct {
	URL string
}

// BufferConfig holds bulk buffer configuration
type BufferConfig struct {
	FlushInterval time.Duration
	MaxBatchSize  int64
	MaxBufferSize int64
}

// RetryConfig holds retry configuration
type RetryConfig struct {
	Attempts   int
	BackoffMin time.Duration
}

// Load loads configuration from environment variables and config files
func Load() (*Config, error) {
	v := viper.New()

	// Set default values
	setDefaults(v)

	// Set config name and paths
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./configs")
	v.AddConfigPath(".")
	v.AddConfigPath("/etc/es-proxy")

	// Read config file (optional)
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}
		// Config file not found; using defaults and env vars
	}

	// Enable automatic env variable binding
	v.AutomaticEnv()

	// Bind specific environment variables
	bindEnvVars(v)

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}

// setDefaults sets default configuration values
func setDefaults(v *viper.Viper) {
	// Server defaults
	v.SetDefault("server.port", "8080")

	// Elasticsearch defaults
	v.SetDefault("elasticsearch.url", "http://localhost:9200")

	// Buffer defaults
	v.SetDefault("buffer.flushinterval", "3s")
	v.SetDefault("buffer.maxbatchsize", 5242880)   // 5MB
	v.SetDefault("buffer.maxbuffersize", 52428800) // 50MB

	// Retry defaults
	v.SetDefault("retry.attempts", 3)
	v.SetDefault("retry.backoffmin", "100ms")
}

// bindEnvVars binds environment variables to config keys
func bindEnvVars(v *viper.Viper) {
	// Use uppercase env vars for compatibility
	v.BindEnv("server.port", "PORT")
	v.BindEnv("elasticsearch.url", "ES_URL")
	v.BindEnv("buffer.flushinterval", "FLUSH_INTERVAL")
	v.BindEnv("buffer.maxbatchsize", "MAX_BATCH_SIZE")
	v.BindEnv("buffer.maxbuffersize", "MAX_BUFFER_SIZE")
	v.BindEnv("retry.attempts", "RETRY_ATTEMPTS")
	v.BindEnv("retry.backoffmin", "RETRY_BACKOFF_MIN")
}
