package sql

import (
	"database/sql"

	"go.uber.org/zap"
)

type repository struct {
	logger *zap.Logger
	client *sql.DB
}

// NewRepository creates a new instance of the SQL repository.
func NewRepository(logger *zap.Logger, client *sql.DB) *repository {
	return &repository{
		logger: logger,
		client: client,
	}
}
