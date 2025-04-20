package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const configFileName = "config.yaml"

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"dbname"`
	} `yaml:"database"`
	Token struct {
		SecretKey string `yaml:"secret_key"`
		ExpiresIn int    `yaml:"expiration"`
	} `yaml:"token"`
}

func NewConfig() (*Config, error) {
	file, err := os.Open(configFileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file %s: %w", configFileName, err)
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to decode config file %s: %w", configFileName, err)
	}

	return &config, nil
}
