package config_test

import (
	"testing"

	"github.com/DucTran999/dbkit/config"
	"github.com/stretchr/testify/require"
)

func TestPostgresqlConfig(t *testing.T) {
	tests := []struct {
		name        string
		config      config.PostgreSQLConfig
		expectedErr error
	}{
		{
			name: "valid config default disable ssl",
			config: config.PostgreSQLConfig{
				Config: config.Config{
					Host:     "localhost",
					Port:     5432,
					Username: "testuser",
					Password: "testpass",
					Database: "testdb",
					Timezone: "UTC",
				},
			},
			expectedErr: nil,
		},
		{
			name: "valid config",
			config: config.PostgreSQLConfig{
				Config: config.Config{
					Host:     "localhost",
					Port:     5432,
					Username: "testuser",
					Password: "testpass",
					Database: "testdb",
					Timezone: "UTC",
				},
				SSLMode: config.PgSSLVerifyFull,
			},
			expectedErr: nil,
		},
		{
			name: "invalid ssl mode",
			config: config.PostgreSQLConfig{
				Config: config.Config{
					Host:     "localhost",
					Port:     5432,
					Username: "testuser",
					Password: "testpass",
					Database: "testdb",
					Timezone: "UTC",
				},
				SSLMode: -1,
			},
			expectedErr: config.ErrPostgresqlSSLMode,
		},
		{
			name: "invalid base config",
			config: config.PostgreSQLConfig{
				Config: config.Config{
					Port:     5432,
					Username: "testuser",
					Password: "testpass",
					Database: "testdb",
					Timezone: "UTC",
				},
			},
			expectedErr: config.ErrMissingHost,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.config.Validate()

			require.ErrorIs(t, err, tc.expectedErr)
		})
	}
}
