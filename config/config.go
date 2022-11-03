package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type (
	Config struct {
		App      `yaml:"app"`
		Log      `yaml:"log"`
		Database `yaml:"database"`
	}

	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	Log struct {
		Level      string `env-required:"true" yaml:"level"   env:"LOG_LEVEL"`
		RollbarEnv string `yaml:"rollbar_env"`
	}

	Database struct {
		Type     string `yaml:"type"`
		Host     string `yaml:"host"`
		PoolMax  int    `yaml:"pool_max"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Pass     string `yaml:"pass"`
		Name     string `yaml:"name"`
		Location string `yaml:"location"`
	}
)

// NewConfig returns app config
func NewConfig() (*Config, error) {
	cfg := &Config{}

	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("unable to read config file, %v", err)
	}

	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, fmt.Errorf("unable to decode into config struct, %v", err)
	}

	return cfg, nil
}
