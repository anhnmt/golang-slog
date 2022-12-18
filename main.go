package main

import (
	"io"
	"os"

	"github.com/natefinch/lumberjack/v3"
	"golang.org/x/exp/slog"
)

func main() {
	// Multi Writer
	writer := []io.Writer{
		getLogWriter("./logs/data.log"),
		os.Stdout,
	}

	textHandler := slog.NewTextHandler(io.MultiWriter(writer...))
	logger := slog.New(textHandler)
	slog.SetDefault(logger)

	slog.Info("Go is the best language!")

	slog.Info("Usage Statistics", slog.Int("current-memory", 50))
}

// getLogWriter returns a lumberjack.logger
func getLogWriter(logFileUrl string) *lumberjack.Roller {
	options := &lumberjack.Options{
		MaxBackups: 5,  // Files
		MaxAge:     30, // 30 days
		Compress:   false,
	}

	var maxSize int64 = 100 * 1024 * 1024 // 100 MB
	roller, err := lumberjack.NewRoller(logFileUrl, maxSize, options)

	if err != nil {
		panic(err)
	}

	return roller
}
