package dialects

import (
	"fmt"

	"github.com/DucTran999/dbkit/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgreSQLDialect struct{}

func NewPostgreSQLDialect() *postgreSQLDialect {
	return &postgreSQLDialect{}
}

// Open opens a PostgreSQL database connection
func (d *postgreSQLDialect) Open(cfg config.PostgreSQLConfig) (*gorm.DB, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		cfg.Host, cfg.Username, cfg.Password, cfg.Database, cfg.Port, cfg.SSLMode, cfg.Timezone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open PostgreSQL connection: %w", err)
	}

	return db, nil
}
