// Package dbkit provides a unified database abstraction layer for Go applications.
// It supports multiple database drivers with consistent interfaces and advanced features.
package dbkit

import (
	"context"
	"fmt"

	"github.com/DucTran999/dbkit/config"
	"github.com/DucTran999/dbkit/dialects"
	"gorm.io/gorm"
)

// Connection represents a database connection with advanced features
type Connection interface {
	// Core database operations
	DB() *gorm.DB
	Close() error
	Ping(ctx context.Context) error
}

// connection implements the Connection interface
type connection struct {
	db *gorm.DB
}

// NewPostgreSQLConnection initializes and returns a new PostgreSQL database connection.
func NewPostgreSQLConnection(cfg config.PostgreSQLConfig) (Connection, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	db, err := dialects.NewPostgreSQLDialect(cfg).Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open PostgreSQL connection: %w", err)
	}

	// Return the fully initialized connection.
	return &connection{db: db}, nil
}

// DB returns the underlying GORM database instance
func (c *connection) DB() *gorm.DB {
	return c.db
}

// Ping tests the database connection
func (c *connection) Ping(ctx context.Context) error {
	sqlDB, err := c.db.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB: %w", err)
	}

	return sqlDB.PingContext(ctx)
}

// Close closes the database connection
func (c *connection) Close() error {
	sqlDB, err := c.db.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB: %w", err)
	}

	return sqlDB.Close()
}
