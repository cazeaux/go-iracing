package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type User struct {
	ID        string `mapstructure:"id"`
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

	return &cfg, nil
}
