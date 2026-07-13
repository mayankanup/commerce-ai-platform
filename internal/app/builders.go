package app

import (
	"context"

	"github.com/mayankanup/commerce-ai-platform/internal/agent"
	inventoryservice "github.com/mayankanup/commerce-ai-platform/internal/inventory/service"
	inventorytool "github.com/mayankanup/commerce-ai-platform/internal/inventory/tool"
	"github.com/mayankanup/commerce-ai-platform/internal/llm/mock"
	"github.com/mayankanup/commerce-ai-platform/internal/storage/sqlite"
)

func buildAgent(db *sqlite.Database) (*agent.Agent, error) {

	inventoryRepo := sqlite.NewInventoryRepository(db)

	inventorySvc := inventoryservice.New(
		inventoryRepo,
	)

	checkInventoryTool := inventorytool.New(
		inventorySvc,
	)

	registry := agent.NewRegistry()

	registry.Register(checkInventoryTool)

	llmClient := mock.New()

	return agent.New(
		llmClient,
		registry,
		agent.Options{
			MaxToolRounds: 5,
		},
	), nil
}

func initializeDatabase(
	ctx context.Context,
	db *sqlite.Database,
) error {

	if err := db.Migrate(ctx); err != nil {
		return err
	}

	return db.Seed(ctx)
}
