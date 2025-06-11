// Package drivers provides database driver implementations for dbkit
package drivers

import (
	"fmt"

	"github.com/DucTran999/dbkit/config"
	"gorm.io/gorm"
)

// Driver represents a database driver type
type Driver int

const (
	// PostgreSQL driver
	PostgreSQL Driver = iota
)

// String returns the string representation of the driver
func (d Driver) String() string {
	switch d {
	case PostgreSQL:
		return "postgresql"
	default:
		return "unknown"
	}
}

// Validate validates the driver
func (d Driver) Validate() error {
	switch d {
	case PostgreSQL:
		return nil
	default:
		return fmt.Errorf("unsupported driver: %d", d)
	}
}

// DriverInterface defines the interface that all database drivers must implement
type DriverInterface interface {
	// Open opens a database connection
	Open(config config.Config) (*gorm.DB, error)

	// Name returns the driver name
	Name() string
}

// NewDriver creates a new driver instance based on the driver type
func NewDriver(driverType Driver) (DriverInterface, error) {
	switch driverType {
	case PostgreSQL:
		return NewPostgreSQLDriver(), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %s", driverType.String())
	}
}
