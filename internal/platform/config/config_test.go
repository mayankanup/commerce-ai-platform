package config

import "testing"

func TestValidate(t *testing.T) {

	cfg := &Config{}

	cfg.Server.Port = 8080
	cfg.Database.Path = "data/ecommerce.db"
	cfg.Database.SchemaPath = "data/schema.sql"
	cfg.Database.SeedPath = "data/seed"
	/*cfg.Ollama.Endpoint = "http://localhost:11434"
	cfg.Ollama.Model = "llama3.1"*/

	if err := Validate(cfg); err != nil {
		t.Fatal(err)
	}
}
