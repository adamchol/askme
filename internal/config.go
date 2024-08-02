package internal

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	OpenAI struct {
		APIKey string `yaml:"api_key"`
	} `yaml:"openai"`
	Claude struct {
		APIKey string `yaml:"api_key"`
	} `yaml:"claude"`
}

func NewConfigService() (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(fmt.Sprintf("%s/askme/askme.yml", homeDir))
	if err != nil {
		return nil, err
	}

	var config Config

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
