// Package dbkit provides a unified database abstraction layer for Go applications.
// It supports multiple database drivers with consistent interfaces and advanced features.
package dbkit

import (
	"context"
	"fmt"

	"github.com/DucTran999/dbkit/config"
	"github.com/DucTran999/dbkit/drivers"
	"gorm.io/gorm"
)

// Connection represents a database connection with advanced features
type Connection interface {
	// Core database operations
	DB() *gorm.DB
	Close(ctx context.Context) error
}

// connection implements the Connection interface
type connection struct {
	db     *gorm.DB
	config config.Config
	driver drivers.DriverInterface
}

// NewConnection creates a new database connection based on the provided configuration
func NewConnection(driver drivers.Driver, cfg config.Config) (Connection, error) {
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	// Create driver instance
	driverInst, err := drivers.NewDriver(driver)
	if err != nil {
		return nil, fmt.Errorf("failed to create driver: %w", err)
	}

	// Open database connection
	db, err := driverInst.Open(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	conn := &connection{
		db:     db,
		config: cfg,
		driver: driverInst,
	}

	return conn, nil
}

// DB returns the underlying GORM database instance
func (c *connection) DB() *gorm.DB {
	return c.db
}

// Close closes the database connection
func (c *connection) Close(ctx context.Context) error {
	if c.db == nil {
		return nil
	}

	sqlDB, err := c.db.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB: %w", err)
	}

	if err := sqlDB.Close(); err != nil {
		return fmt.Errorf("failed to close connection: %w", err)
	}

	return nil
}
