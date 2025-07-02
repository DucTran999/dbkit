package connections

import (
	"fmt"

	"github.com/DucTran999/dbkit/config"
	"github.com/DucTran999/dbkit/dialects"
)

// NewClickHouseConnection initializes and returns a new ClickHouse database connection.
func NewClickHouseConnection(cfg config.ClickHouseConfig) (*connection, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	db, err := dialects.NewClickHouseDialect(cfg).Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open MySQL connection: %w", err)
	}

	// Return the fully initialized connection.
	return &connection{db: db}, nil
}
