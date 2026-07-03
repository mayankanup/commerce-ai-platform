package config

type Config struct {
	App        AppConfig        `yaml:"app"`
	Server     ServerConfig     `yaml:"server"`
	Database   DatabaseConfig   `yaml:"database"`
	Ollama     OllamaConfig     `yaml:"ollama"`
	Telemetry  TelemetryConfig  `yaml:"telemetry"`
	Evaluation EvaluationConfig `yaml:"evaluation"`
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
	Path string `yaml:"path"`
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
