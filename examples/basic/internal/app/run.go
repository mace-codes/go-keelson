package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/mace-codes/go-keelson-transport/utils/logger/zaplog"
	"github.com/mace-codes/go-keelson/examples/basic/internal/app/config"
	"go.uber.org/zap"

	transport "github.com/mace-codes/go-keelson-transport"
)

func Run() error {
	_ = godotenv.Load(".env") // Load environment variables from .env file

	cfg, err := config.Load() // Load configuration from environment variables
	if err != nil {
		return err
	}

	deps, err := initializeDependencies(cfg) // Initialize dependencies based on the loaded configuration
	if err != nil {
		return err
	}
	defer deps.logger.Sync()

	// Start all startable dependencies
	for _, s := range deps.startables {
		if err := s.Start(); err != nil {
			return err
		}
	}

	// Ensure that all stoppable dependencies are stopped when the application exits
	defer func() {
		for _, s := range deps.stopables {
			if err := s.Stop(); err != nil {
				// Log the error but continue stopping other dependencies
				deps.logger.Error("Failed to stop dependency", zap.Error(err))
			}
		}
	}()

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
		},
		AllowCredentials: true,
		Debug:            false,
	}))

	srv, err := transport.NewTransport(
		deps,
		deps.cfg,
		serveRoutes,
		router,
		transport.RegisterChiRoutes,
		transport.WithLogger(zaplog.New(deps.logger)),
	)
	if err != nil {
		return err
	}

	return srv.ListenAndServe()
}
