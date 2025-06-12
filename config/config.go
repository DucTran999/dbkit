// Package config provides configuration structures and validation for dbkit
package config

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrMissingHost     = errors.New("host is required")
	ErrInvalidPort     = errors.New("port must be between 1 and 65535")
	ErrMissingUsername = errors.New("username is required")
	ErrMissingDatabase = errors.New("database name is required")
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
	var errs []error

	// Validate required fields
	if strings.TrimSpace(c.Host) == "" {
		errs = append(errs, ErrMissingHost)
	}

	if c.Port <= 0 || c.Port > 65535 {
		errs = append(errs, ErrInvalidPort)
	}

	if strings.TrimSpace(c.Username) == "" {
		errs = append(errs, ErrMissingUsername)
	}

	if strings.TrimSpace(c.Database) == "" {
		errs = append(errs, ErrMissingDatabase)
	}

	if len(errs) > 0 {
		return fmt.Errorf("configuration validation failed: %v", errs)
	}

	return nil
}
