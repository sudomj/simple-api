package logger

import (
	"log/slog"
	"os"
)

type logger struct {
	Level string
}

func NewLogger(level string) *logger {

	return &logger{
		Level: level,
	}
}

func (l *logger) Configure() {

	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.Level(l.parseLevel()),
	})
	logger := slog.New(logHandler)
	slog.SetDefault(logger)
}

func (l *logger) parseLevel() int {

	levelMap := map[string]int{
		"DEBUG": -4,
		"INFO":  0,
		"WARN":  4,
		"ERROR": 8,
	}

	return levelMap[l.Level]
}
