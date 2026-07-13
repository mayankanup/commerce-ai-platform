package app

import (
	"context"
	"fmt"

	"github.com/mayankanup/commerce-ai-platform/internal/platform/config"
	"github.com/mayankanup/commerce-ai-platform/internal/platform/logging"
	"github.com/mayankanup/commerce-ai-platform/internal/platform/server"
	"github.com/mayankanup/commerce-ai-platform/internal/platform/web"
	"github.com/mayankanup/commerce-ai-platform/internal/storage/sqlite"
)

func Bootstrap(options Options) (*Application, error) {

	cfg, err := config.Load(options.ConfigFile)
	if err != nil {
		return nil, err
	}

	logger := logging.New(
		logging.Options{
			Level:  logging.ParseLevel(cfg.Logging.Level),
			Format: cfg.Logging.Format,
		},
	)

	logger.Info(
		"Commerce AI Platform starting",
		"version", cfg.App.Version,
		"environment", cfg.App.Environment,
	)

	db, err := sqlite.New(
		sqlite.Options{
			Path:            cfg.Database.Path,
			SchemaPath:      cfg.Database.SchemaPath,
			SeedPath:        cfg.Database.SeedPath,
			MaxOpenConns:    cfg.Database.MaxOpenConns,
			MaxIdleConns:    cfg.Database.MaxIdleConns,
			ConnMaxLifetime: cfg.Database.ConnMaxLifetime,
			Logger:          logger,
		},
	)
	if err != nil {
		return nil, err
	}

	logger.Info(
		"SQLite connected",
		"path", cfg.Database.Path,
		"schemaPath", cfg.Database.SchemaPath,
		"seedPath", cfg.Database.SeedPath,
	)

	ctx := context.Background()

	if err := initializeDatabase(ctx, db); err != nil {
		return nil, err
	}

	aiAgent, err := buildAgent(cfg, db)
	if err != nil {
		return nil, err
	}

	router := web.NewRouter(
		web.ApplicationInfo{
			Name:        cfg.App.Name,
			Version:     cfg.App.Version,
			Environment: cfg.App.Environment,
		},
		logger,
	)

	srv := server.New(
		server.Options{
			Address: fmt.Sprintf(
				"%s:%d",
				cfg.Server.Host,
				cfg.Server.Port,
			),
			Logger:  logger,
			Handler: router,
		},
	)

	return &Application{
		Config: cfg,
		Logger: logger,
		DB:     db,
		Server: srv,
		Agent:  aiAgent,
	}, nil
}
