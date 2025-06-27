package config

import (
	"fmt"

	"github.com/pkg/errors"
)

var (
	ErrPostgresqlSSLMode = errors.New("invalid ssl mode")
)

type PgSSLConfig string

const (
	PgSSLDisable    PgSSLConfig = "disable"
	PgSSLVerifyFull PgSSLConfig = "verify-full"
)

func (p PgSSLConfig) Validate() error {
	switch p {
	case PgSSLDisable:
		return nil
	case PgSSLVerifyFull:
		return nil
	default:
		return fmt.Errorf("unsupported SSL mode: %s", p)
	}
}

// PostgreSQLConfig holds PostgreSQL-specific configuration
type PostgreSQLConfig struct {
	Config

	PoolConfig

	SSLMode PgSSLConfig
}

func (pc *PostgreSQLConfig) Validate() error {
	if err := pc.Config.Validate(); err != nil {
		return err
	}

	pc.PoolConfig.SetDefaults()

	if err := pc.SSLMode.Validate(); err != nil {
		pc.SSLMode = PgSSLDisable
	}

	return nil
}
