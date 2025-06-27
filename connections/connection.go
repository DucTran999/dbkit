package connections

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

// connection implements the Connection interface
type connection struct {
	db *gorm.DB
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
