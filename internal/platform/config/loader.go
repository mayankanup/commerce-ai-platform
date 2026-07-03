package config

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

func Load(path string) (*Config, error) {

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read config: %w", err)
	}

	var cfg Config

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("unable to parse yaml: %w", err)
	}

	overrideWithEnv(&cfg)

	if err := Validate(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func overrideWithEnv(cfg *Config) {

	if port := os.Getenv("SERVER_PORT"); port != "" {

		if p, err := strconv.Atoi(port); err == nil {
			cfg.Server.Port = p
		}
	}

	if model := os.Getenv("OLLAMA_MODEL"); model != "" {
		cfg.Ollama.Model = model
	}

	if endpoint := os.Getenv("OLLAMA_ENDPOINT"); endpoint != "" {
		cfg.Ollama.Endpoint = endpoint
	}
}
