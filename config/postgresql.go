package config

import (
	"strings"

	"github.com/pkg/errors"
)

var (
	ErrPostgresqlSSLMode = errors.New("invalid ssl mode")
)

type PgSSLConfig int

func (p PgSSLConfig) String() string {
	switch p {
	case PgSSLDisable:
		return "disable"
	case PgSSLVerifyFull:
		return "verify-full"
	default:
		return ""
	}
}

const (
	PgSSLDisable PgSSLConfig = iota
	PgSSLVerifyFull
)

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

	if strings.TrimSpace(pc.SSLMode.String()) == "" {
		return ErrPostgresqlSSLMode
	}

	return nil
}
