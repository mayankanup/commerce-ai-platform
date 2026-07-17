package config

import "fmt"

func Validate(cfg *Config) error {

	if cfg.Server.Port <= 0 {
		return fmt.Errorf("server.port must be greater than zero")
	}

	if cfg.Database.Path == "" {
		return fmt.Errorf("database.path is required")
	}

	if cfg.Database.SchemaPath == "" {
		return fmt.Errorf("database.schemaPath is required")
	}

	if cfg.Database.SeedPath == "" {
		return fmt.Errorf("database.seedPath is required")
	}

	/*if cfg.Ollama.Endpoint == "" {
		return fmt.Errorf("ollama.endpoint is required")
	}

	if cfg.Ollama.Model == "" {
		return fmt.Errorf("ollama.model is required")
	}*/

	return nil
}
