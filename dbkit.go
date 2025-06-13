// Package dbkit provides a unified database abstraction layer for Go applications.
// It supports multiple database drivers with consistent interfaces and advanced features.
package dbkit

import (
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
	Ping() error
}

// connection implements the Connection interface
type connection struct {
	db  *gorm.DB
	cfg config.PostgreSQLConfig
}

// NewPostgreSQLConnection initializes and returns a new PostgreSQL database connection.
func NewPostgreSQLConnection(cfg config.PostgreSQLConfig) (Connection, error) {
	// Validate the configuration values before proceeding.
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	// Open a new GORM DB instance using the PostgreSQL dialect and provided config.
	db, err := dialects.NewPostgreSQLDialect().Open(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to open PostgreSQL connection: %w", err)
	}

	// Wrap the GORM DB instance and config in a Connection implementation.
	conn := &connection{
		db:  db,
		cfg: cfg,
	}

	// Initialize the connection pool (e.g., max open/idle connections).
	if err := conn.initPool(); err != nil {
		return nil, fmt.Errorf("failed to initialize connection pool: %w", err)
	}

	// Return the fully initialized connection.
	return conn, nil
}

// DB returns the underlying GORM database instance
func (c *connection) DB() *gorm.DB {
	return c.db
}

// Ping tests the database connection
func (c *connection) Ping() error {
	sqlDB, err := c.db.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	return nil
}

// Close closes the database connection
func (c *connection) Close() error {
	sqlDB, err := c.db.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB: %w", err)
	}

	if err := sqlDB.Close(); err != nil {
		return fmt.Errorf("failed to close connection: %w", err)
	}

	return nil
}

func (c *connection) initPool() error {
	sqlDB, err := c.db.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB: %w", err)
	}

	// Apply connection pool settings
	sqlDB.SetMaxIdleConns(c.cfg.MaxIdleConnection)
	sqlDB.SetMaxOpenConns(c.cfg.MaxOpenConnection)
	sqlDB.SetConnMaxLifetime(c.cfg.ConnMaxLifetime)
	sqlDB.SetConnMaxIdleTime(c.cfg.ConnMaxIdleTime)

	return nil
}
