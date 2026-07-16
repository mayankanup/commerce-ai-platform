package app

import (
	"testing"

	"github.com/mayankanup/commerce-ai-platform/internal/platform/config"
)

func TestBuildMemoryRepository(t *testing.T) {

	cfg := &config.Config{}

	cfg.RAG.VectorDB.Provider = "memory"

	repo, err := buildVectorRepository(cfg)
	if err != nil {
		t.Fatal(err)
	}

	if repo == nil {
		t.Fatal("repository is nil")
	}
}

func TestBuildUnknownRepository(t *testing.T) {

	cfg := &config.Config{}

	cfg.RAG.VectorDB.Provider = "invalid"

	_, err := buildVectorRepository(cfg)

	if err == nil {
		t.Fatal("expected error")
	}
}
