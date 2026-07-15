package app

import (
	"context"

	"github.com/mayankanup/commerce-ai-platform/internal/agent"
	catalogservice "github.com/mayankanup/commerce-ai-platform/internal/catalog/service"
	catalogtool "github.com/mayankanup/commerce-ai-platform/internal/catalog/tool"
	inventoryservice "github.com/mayankanup/commerce-ai-platform/internal/inventory/service"
	inventorytool "github.com/mayankanup/commerce-ai-platform/internal/inventory/tool"
	llmfactory "github.com/mayankanup/commerce-ai-platform/internal/llm/factory"
	"github.com/mayankanup/commerce-ai-platform/internal/platform/config"
	"github.com/mayankanup/commerce-ai-platform/internal/storage/sqlite"
)

func buildAgent(cfg *config.Config, db *sqlite.Database) (*agent.Agent, error) {

	registry := agent.NewRegistry()
	createAndRegisterInventoryTool(db, registry)
	createAndRegisterCatalogTool(db, registry)

	llmClient, err := llmfactory.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return agent.New(
		llmClient,
		registry,
		agent.Options{
			MaxToolRounds: 5,
		},
	), nil
}

func createAndRegisterInventoryTool(db *sqlite.Database, registry *agent.Registry) {
	inventoryRepo := sqlite.NewInventoryRepository(db)

	inventorySvc := inventoryservice.New(
		inventoryRepo,
	)

	checkInventoryTool := inventorytool.New(
		inventorySvc,
	)

	registry.Register(checkInventoryTool)
}

func createAndRegisterCatalogTool(db *sqlite.Database, registry *agent.Registry) {
	catalogRepo := sqlite.NewCatalogRepository(db)

	catalogService := catalogservice.New(
		catalogRepo,
	)

	searchProductsTool := catalogtool.New(
		catalogService,
	)

	registry.Register(searchProductsTool)
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
