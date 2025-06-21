// config/config.go
package config

import (
	"fmt"
	"strings"
	"sync"

	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

// Global config instance
var (
	globalConfig *Config
	configOnce   sync.Once
)

type Config struct {
	OpenAI  OpenAIConfig  `mapstructure:"openai"`
	Vikunja VikunjaConfig `mapstructure:"vikunja"`
	Server  ServerConfig  `mapstructure:"server"`
	Log     LogConfig     `mapstructure:"log"`
}

type OpenAIConfig struct {
	APIKey string `mapstructure:"api_key"`
	Model  string `mapstructure:"model"`
}

type VikunjaConfig struct {
	URL   string `mapstructure:"url"`
	Token string `mapstructure:"token"`
}

type ServerConfig struct {
	Port int `mapstructure:"port"`
}

type LogConfig struct {
	Level string `mapstructure:"level"`
}

// Load loads config once and stores it globally
func Load() (*Config, error) {
	var err error
	configOnce.Do(func() {
		globalConfig, err = loadConfig()
	})
	return globalConfig, err
}

// Get returns the global config (must call Load() first)
func Get() *Config {
	if globalConfig == nil {
		panic("config not loaded - call config.Load() first")
	}
	return globalConfig
}

// GetOpenAI returns OpenAI config
func GetOpenAI() OpenAIConfig {
	return Get().OpenAI
}

// GetVikunja returns Vikunja config
func GetVikunja() VikunjaConfig {
	return Get().Vikunja
}

func loadConfig() (*Config, error) {
	viper.SetEnvPrefix("") // No prefix, direct env var names
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Set defaults
	viper.SetDefault("openai.model", "gpt-4o-mini")
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("log.level", "info")

	// Bind specific env vars
	if err := viper.BindEnv("openai.api_key", "OPENAI_API_KEY"); err != nil {
		return nil, fmt.Errorf("failed to bind OPENAI_API_KEY: %w", err)
	}
	if err := viper.BindEnv("openai.model", "OPENAI_MODEL"); err != nil {
		return nil, fmt.Errorf("failed to bind OPENAI_MODEL: %w", err)
	}
	if err := viper.BindEnv("vikunja.url", "VIKUNJA_URL"); err != nil {
		return nil, fmt.Errorf("failed to bind VIKUNJA_URL: %w", err)
	}
	if err := viper.BindEnv("vikunja.token", "VIKUNJA_TOKEN"); err != nil {
		return nil, fmt.Errorf("failed to bind VIKUNJA_TOKEN: %w", err)
	}
	if err := viper.BindEnv("log.level", "LOG_LEVEL"); err != nil {
		return nil, fmt.Errorf("failed to bind LOG_LEVEL: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Validate required fields
	if err := validate(&config); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	// Set log level based on config
	setLogLevel(config.Log.Level)

	// Log loaded configuration (without secrets)
	logConfig(&config)

	return &config, nil
}

func validate(config *Config) error {
	if config.OpenAI.APIKey == "" {
		return fmt.Errorf("OPENAI_API_KEY is required")
	}
	if config.Vikunja.URL == "" {
		return fmt.Errorf("VIKUNJA_URL is required")
	}
	if config.Vikunja.Token == "" {
		return fmt.Errorf("VIKUNJA_TOKEN is required")
	}
	return nil
}

func setLogLevel(level string) {
	switch strings.ToLower(level) {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn", "warning":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.InfoLevel)
		log.Warn("Invalid log level, defaulting to info", "level", level)
	}
}

func logConfig(config *Config) {
	log.Info("Configuration loaded",
		"openai_model", config.OpenAI.Model,
		"vikunja_url", config.Vikunja.URL,
		"server_port", config.Server.Port,
		"log_level", config.Log.Level,
		"openai_api_key_set", config.OpenAI.APIKey != "",
		"vikunja_token_set", config.Vikunja.Token != "",
	)
}
