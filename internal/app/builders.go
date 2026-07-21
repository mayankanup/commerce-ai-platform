package app

import (
	"context"

	"github.com/mayankanup/commerce-ai-platform/internal/agent"
	catalogservice "github.com/mayankanup/commerce-ai-platform/internal/catalog/service"
	catalogtool "github.com/mayankanup/commerce-ai-platform/internal/catalog/tool"
	"github.com/mayankanup/commerce-ai-platform/internal/embedding"
	inventoryservice "github.com/mayankanup/commerce-ai-platform/internal/inventory/service"
	inventorytool "github.com/mayankanup/commerce-ai-platform/internal/inventory/tool"
	llmfactory "github.com/mayankanup/commerce-ai-platform/internal/llm/factory"
	"github.com/mayankanup/commerce-ai-platform/internal/platform/config"
	"github.com/mayankanup/commerce-ai-platform/internal/prompt"
	ragretriever "github.com/mayankanup/commerce-ai-platform/internal/rag/retriever"
	ragtool "github.com/mayankanup/commerce-ai-platform/internal/rag/tool"
	"github.com/mayankanup/commerce-ai-platform/internal/storage/sqlite"

	ragrepo "github.com/mayankanup/commerce-ai-platform/internal/rag/repository"
)

func buildAgent(cfg *config.Config, db *sqlite.Database, embeddingProvider embedding.Provider, vectorRepository ragrepo.Repository) (*agent.Agent, error) {

	registry := agent.NewRegistry()
	createAndRegisterInventoryTool(db, registry)
	createAndRegisterCatalogTool(db, registry)
	createAndRegisterSearchDocumentsTool(registry, embeddingProvider, vectorRepository)

	llmClient, err := llmfactory.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	promptBuilder := prompt.New(
		prompt.DefaultSystemPrompt,
	)

	return agent.New(
		llmClient,
		promptBuilder,
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

func createAndRegisterSearchDocumentsTool(registry *agent.Registry, embeddingProvider embedding.Provider, vectorRepository ragrepo.Repository) {
	retriever := ragretriever.New(
		embeddingProvider,
		vectorRepository,
	)

	searchDocumentsTool := ragtool.NewSearchDocumentsTool(retriever)

	registry.Register(searchDocumentsTool)
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
