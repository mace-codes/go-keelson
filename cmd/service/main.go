package service

import (
	"github.com/mace-codes/go-keelson/internal/app"
	"go.uber.org/zap"
)

// main is the entry point for the service. It initializes the application and starts it.
func main() {
	if err := app.Run(); err != nil {
		logger, _ := zap.NewProduction()
		defer logger.Sync()
		logger.Fatal("Failed to run the service", zap.Error(err))
	}
}
