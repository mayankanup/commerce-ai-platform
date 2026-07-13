package app

import (
	"context"

	"github.com/mayankanup/commerce-ai-platform/internal/agent"
	"github.com/mayankanup/commerce-ai-platform/internal/platform/config"
	"github.com/mayankanup/commerce-ai-platform/internal/platform/logging"
	"github.com/mayankanup/commerce-ai-platform/internal/platform/server"
	"github.com/mayankanup/commerce-ai-platform/internal/storage/sqlite"
)

type Application struct {
	Config *config.Config
	Logger *logging.Logger

	DB     *sqlite.Database
	Server *server.Server
	Agent  *agent.Agent
}

func (a *Application) Start() error {
	return a.Server.Start()
}

func (a *Application) Shutdown(ctx context.Context) error {

	if a.Server != nil {
		if err := a.Server.Shutdown(ctx); err != nil {
			return err
		}
	}

	if a.DB != nil {
		if err := a.DB.Close(); err != nil {
			return err
		}
	}

	return nil
}
