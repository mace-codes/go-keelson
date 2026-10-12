package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/mace-codes/go-keelson-transport/utils/logger/zaplog"
	"github.com/mace-codes/go-keelson/internal/app/config"

	transport "github.com/mace-codes/go-keelson-transport"
)

func Run() error {
	godotenv.Load() // Load environment variables from .env file

	cfg, err := config.Load() // Load configuration from environment variables
	if err != nil {
		return err
	}

	deps, err := initializeDependencies(cfg) // Initialize dependencies based on the loaded configuration
	if err != nil {
		return err
	}

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
				// You can use a logger here to log the error
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
