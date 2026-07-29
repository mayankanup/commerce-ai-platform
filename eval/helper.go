package eval

import (
	"testing"

	"github.com/mayankanup/commerce-ai-platform/internal/app"
)

func NewTestApplication(t *testing.T) *app.Application {

	t.Helper()

	application, err := app.Bootstrap(
		app.Options{
			ConfigFile: "../config/config.yaml",
		},
	)

	if err != nil {
		t.Fatalf("bootstrap failed: %v", err)
	}

	return application
}
