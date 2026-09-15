package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/mace-codes/go-keelson/internal/app/config"
)

func Run() error {
	godotenv.Load() // Load environment variables from .env file

	cfg, err := config.Load()
	if err != nil {
		return err
	}

	deps, err := initializeDependencies(cfg)
	if err != nil {
		return err
	}

	for _, s := range deps.startables {
		if err := s.Start(); err != nil {
			return err
		}
	}

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

	// serve http, routes, middleware, etc. here

	return nil
}
