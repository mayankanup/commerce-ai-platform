package config

import "time"

type Config struct {
	App        AppConfig        `yaml:"app"`
	Server     ServerConfig     `yaml:"server"`
	Database   DatabaseConfig   `yaml:"database"`
	Ollama     OllamaConfig     `yaml:"ollama"`
	Logging    LoggingConfig    `yaml:"logging"`
	Telemetry  TelemetryConfig  `yaml:"telemetry"`
	Evaluation EvaluationConfig `yaml:"evaluation"`
	LLM        LLMConfig        `yaml:"llm"`
	Embedding  EmbeddingConfig  `yaml:"embedding"`
}

type LoggingConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}

type AppConfig struct {
	Name        string `yaml:"name"`
	Version     string `yaml:"version"`
	Environment string `yaml:"environment"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type DatabaseConfig struct {
	Path            string        `yaml:"path"`
	SchemaPath      string        `yaml:"schemaPath"`
	SeedPath        string        `yaml:"seedPath"`
	MaxOpenConns    int           `yaml:"maxOpenConns"`
	MaxIdleConns    int           `yaml:"maxIdleConns"`
	ConnMaxLifetime time.Duration `yaml:"connMaxLifetime"`
}

type OllamaConfig struct {
	Endpoint string `yaml:"endpoint"`
	Model    string `yaml:"model"`
}

type TelemetryConfig struct {
	Enabled     bool   `yaml:"enabled"`
	ServiceName string `yaml:"service_name"`
}

type EvaluationConfig struct {
	Enabled bool `yaml:"enabled"`
}

type LLMConfig struct {
	Provider string `yaml:"provider"`
}

type EmbeddingConfig struct {
	Provider string `yaml:"provider"`
	Model    string `yaml:"model"`
	Endpoint string `yaml:"endpoint"`
}
