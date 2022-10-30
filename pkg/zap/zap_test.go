// @author moqi
// On 2022/10/30 18:32:24
package main

import (
	"go.uber.org/zap"
	"testing"
	"time"
)

const url = "https://www.uber.com"

func TestSugaredLogger(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger) // flushes buffer, if any

	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}

func TestZapLogger(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}
