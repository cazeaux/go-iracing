package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type User struct {
	ID        int    `mapstructure:"id"`
	Timestamp string `mapstructure:"timestamp"`
	Enabled   bool   `mapstructure:"enabled"`
	DiscordID string `mapstructure:"discordId"`
	Color     int    `mapstructure:"color"`
}

type Notification struct {
	DiscordChannel string `mapstructure:"discordChannel"`
}

type Config struct {
	Users        []User       `mapstructure:"users"`
	Notification Notification `mapstructure:"notification"`
	Env          ConfigEnv
}

type ConfigEnv struct {
	Email        string
	Password     string
	ClientID     string
	ClientSecret string
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigName("config") // "config.yaml"
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path) // folder where config.yaml is located

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode config: %w", err)
	}
	cfg.LoadEnv()

	return &cfg, nil
}

func (c *Config) LoadEnv() {
	configEnv := ConfigEnv{
		Email:        os.Getenv("IR_EMAIL"),
		Password:     os.Getenv("IR_PASSWORD"),
		ClientID:     os.Getenv("IR_CLIENT_ID"),
		ClientSecret: os.Getenv("IR_CLIENT_SECRET"),
	}

	c.Env = configEnv
}
