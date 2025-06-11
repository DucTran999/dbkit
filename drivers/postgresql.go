package drivers

import (
	"fmt"

	"github.com/DucTran999/dbkit/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgreSQLDriver implements the DriverInterface for PostgreSQL
type PostgreSQLDriver struct{}

// NewPostgreSQLDriver creates a new PostgreSQL driver instance
func NewPostgreSQLDriver() *PostgreSQLDriver {
	return &PostgreSQLDriver{}
}

// Open opens a PostgreSQL database connection
func (d *PostgreSQLDriver) Open(config config.Config) (*gorm.DB, error) {
	dsn := ""

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open PostgreSQL connection: %w", err)
	}

	return db, nil
}

func (d *PostgreSQLDriver) Name() string {
	return PostgreSQL.String()
}
