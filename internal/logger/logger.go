package logger

import (
	"io"
	"log/slog"
	"os"
)

type CakeLogger struct {
	*slog.Logger
}

func New(debug bool) (*CakeLogger, error) {
	var handler slog.Handler

	if debug {
		f, err := os.OpenFile("cake.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
		if err != nil {
			return nil, err
		}

		handler = slog.NewTextHandler(f, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	} else {
		handler = slog.NewTextHandler(io.Discard, nil)
	}

	return &CakeLogger{
		Logger: slog.New(handler),
	}, nil
}
