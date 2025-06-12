package config

import (
	"strings"

	"github.com/pkg/errors"
)

type PgSSLConfig int

func (p PgSSLConfig) String() string {
	switch p {
	case PgSSLDisable:
		return "disable"
	case PgSSLVerifyFull:
		return "verify-full"
	default:
		return "unknown"
	}
}

const (
	PgSSLDisable PgSSLConfig = iota
	PgSSLVerifyFull
)

// PostgreSQLConfig holds PostgreSQL-specific configuration
type PostgreSQLConfig struct {
	Config

	SSLMode PgSSLConfig
}

func (pc *PostgreSQLConfig) Validate() error {
	if err := pc.Config.Validate(); err != nil {
		return err
	}

	if strings.TrimSpace(pc.SSLMode.String()) == "" {
		return errors.New("sslmode cannot be empty")
	}

	return nil
}
