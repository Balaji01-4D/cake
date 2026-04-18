package cli

import (
	"github.com/balaji01-4d/cake/internal/app"
	"github.com/balaji01-4d/cake/internal/logger"
)

type CLIContext struct {
	App    app.App
	Logger *logger.CakeLogger
}
