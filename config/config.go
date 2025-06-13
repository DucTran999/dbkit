// Package config provides configuration structures and validation for dbkit
package config

import (
	"errors"
	"strings"
	"time"
)

var (
	ErrMissingHost     = errors.New("host is required")
	ErrInvalidPort     = errors.New("port must be between 1 and 65535")
	ErrMissingUsername = errors.New("username is required")
	ErrMissingDatabase = errors.New("database name is required")
)

const (
	DefaultMaxIdleConnection     = 10
	DefaultMaxOpenConnection     = 100
	DefaultConnectionMaxLifetime = time.Hour
	DefaultConnectionMaxIdleTime = 10 * time.Minute
)

// Config represents the complete database configuration
type Config struct {
	// Address
	Host string
	Port int

	// Authentication information
	Username string
	Password string
	Database string

	// Connection options
	Timezone string
}

// Validate validates the configuration
func (c *Config) Validate() error {
	// Validate required fields
	if strings.TrimSpace(c.Host) == "" {
		return ErrMissingHost
	}

	if c.Port <= 0 || c.Port > 65535 {
		return ErrInvalidPort
	}

	if strings.TrimSpace(c.Username) == "" {
		return ErrMissingUsername
	}

	if strings.TrimSpace(c.Database) == "" {
		return ErrMissingDatabase
	}

	return nil
}

type PoolConfig struct {
	// MaxIdleConnection sets the maximum number of connections in the idle connection pool.
	// Default: 10, Recommended: 2-10 depending on load
	MaxIdleConnection int

	// MaxOpenConnection sets the maximum number of open connections to the database.
	// Default: 100, Should be tuned based on your database server capacity
	MaxOpenConnection int

	// ConnMaxLifetime sets the maximum amount of time a connection may be reused.
	// Default: 1 hour, Recommended: 5m-1h to prevent stale connections
	ConnMaxLifetime time.Duration

	// ConnMaxIdleTime sets the maximum amount of time a connection may be idle.
	// Default: 10 minutes, helps clean up unused connections
	ConnMaxIdleTime time.Duration
}

func (pc *PoolConfig) SetDefaults() {
	if pc.MaxIdleConnection <= 0 {
		pc.MaxIdleConnection = DefaultMaxIdleConnection
	}

	if pc.MaxOpenConnection <= 0 {
		pc.MaxOpenConnection = DefaultMaxOpenConnection
	}

	// Ensure MaxOpenConnection >= MaxIdleConnection
	if pc.MaxOpenConnection < pc.MaxIdleConnection {
		pc.MaxOpenConnection = pc.MaxIdleConnection
	}

	if pc.ConnMaxLifetime <= 0 {
		pc.ConnMaxLifetime = DefaultConnectionMaxLifetime
	} else if pc.ConnMaxLifetime > 24*time.Hour {
		pc.ConnMaxLifetime = 24 * time.Hour // Cap at 24 hours
	}

	if pc.ConnMaxIdleTime <= 0 {
		pc.ConnMaxIdleTime = DefaultConnectionMaxIdleTime
	} else if pc.ConnMaxIdleTime > time.Hour {
		pc.ConnMaxIdleTime = time.Hour // Cap at 1 hour
	}
}
