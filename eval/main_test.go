package eval

import (
	"os"
	"testing"

	"github.com/mayankanup/commerce-ai-platform/internal/app"
)

var testApp *app.Application

func TestMain(m *testing.M) {

	var err error

	testApp, err = app.Bootstrap(
		app.Options{
			ConfigFile: "../config/config.yaml",
		},
	)

	if err != nil {
		panic(err)
	}

	code := m.Run()

	os.Exit(code)
}
