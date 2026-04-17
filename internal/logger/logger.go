package logger

import (
	"log/slog"
	"os"
)

type CakeLogger struct {
	*slog.Logger
}

func NewLogger(debug bool) *CakeLogger {
	level := slog.LevelInfo
	if debug {
		level = slog.LevelDebug
	}

	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: level,
	})
	return &CakeLogger{
		Logger: slog.New(handler),
	}
}
