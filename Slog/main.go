package main

import (
	"log/slog"
	"os"
)

func main() {
	handler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(handler)
	slog.SetDefault(logger)

	logger.Info("error", "errid", 1, slog.Group("grp1", "id", 1, "parameter", "elite"), "context", "far")
	child := logger.With("childid", 1)
	child.Info("issuee", "type", "active")
}
