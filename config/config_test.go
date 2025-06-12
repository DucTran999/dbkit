package config_test

import (
	"testing"

	"github.com/DucTran999/dbkit/config"
	"github.com/stretchr/testify/require"
)

func TestConfigValidate(t *testing.T) {
	tests := []struct {
		name        string
		config      config.Config
		expectedErr error
	}{
		{
			name: "valid config",
			config: config.Config{
				Host:     "localhost",
				Port:     5432,
				Username: "testuser",
				Password: "testpass",
				Database: "testdb",
				Timezone: "UTC",
			},
			expectedErr: nil,
		},
		{
			name: "valid config with minimal fields",
			config: config.Config{
				Host:     "127.0.0.1",
				Port:     3306,
				Username: "user",
				Database: "db",
			},
			expectedErr: nil,
		},
		{
			name: "missing host",
			config: config.Config{
				Port:     5432,
				Username: "testuser",
				Database: "testdb",
			},
			expectedErr: config.ErrMissingHost,
		},
		{
			name: "empty host",
			config: config.Config{
				Host:     "",
				Port:     5432,
				Username: "testuser",
				Database: "testdb",
			},
			expectedErr: config.ErrMissingHost,
		},
		{
			name: "whitespace only host",
			config: config.Config{
				Host:     "   ",
				Port:     5432,
				Username: "testuser",
				Database: "testdb",
			},
			expectedErr: config.ErrMissingHost,
		},
		{
			name: "invalid port - zero",
			config: config.Config{
				Host:     "localhost",
				Port:     0,
				Username: "testuser",
				Database: "testdb",
			},
			expectedErr: config.ErrInvalidPort,
		},
		{
			name: "invalid port - negative",
			config: config.Config{
				Host:     "localhost",
				Port:     -1,
				Username: "testuser",
				Database: "testdb",
			},
			expectedErr: config.ErrInvalidPort,
		},
		{
			name: "invalid port - too high",
			config: config.Config{
				Host:     "localhost",
				Port:     65536,
				Username: "testuser",
				Database: "testdb",
			},
			expectedErr: config.ErrInvalidPort,
		},
		{
			name: "missing username",
			config: config.Config{
				Host:     "localhost",
				Port:     5432,
				Database: "testdb",
			},
			expectedErr: config.ErrMissingUsername,
		},
		{
			name: "empty username",
			config: config.Config{
				Host:     "localhost",
				Port:     5432,
				Username: "",
				Database: "testdb",
			},
			expectedErr: config.ErrMissingUsername,
		},
		{
			name: "whitespace only username",
			config: config.Config{
				Host:     "localhost",
				Port:     5432,
				Username: "   ",
				Database: "testdb",
			},
			expectedErr: config.ErrMissingUsername,
		},
		{
			name: "missing database",
			config: config.Config{
				Host:     "localhost",
				Port:     5432,
				Username: "testuser",
			},
			expectedErr: config.ErrMissingDatabase,
		},
		{
			name: "empty database",
			config: config.Config{
				Host:     "localhost",
				Port:     5432,
				Username: "testuser",
				Database: "",
			},
			expectedErr: config.ErrMissingDatabase,
		},
		{
			name: "whitespace only database",
			config: config.Config{
				Host:     "localhost",
				Port:     5432,
				Username: "testuser",
				Database: "   ",
			},
			expectedErr: config.ErrMissingDatabase,
		},
		{
			name: "valid port boundaries",
			config: config.Config{
				Host:     "localhost",
				Port:     1,
				Username: "testuser",
				Database: "testdb",
			},
			expectedErr: nil,
		},
		{
			name: "valid port upper boundary",
			config: config.Config{
				Host:     "localhost",
				Port:     65535,
				Username: "testuser",
				Database: "testdb",
			},
			expectedErr: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.config.Validate()

			require.ErrorIs(t, err, tc.expectedErr)
		})
	}
}

// Benchmark test for validation performance
func BenchmarkConfig_Validate(b *testing.B) {
	config := config.Config{
		Host:     "localhost",
		Port:     5432,
		Username: "testuser",
		Password: "testpass",
		Database: "testdb",
		Timezone: "UTC",
	}

	b.ResetTimer()
	for range b.N {
		_ = config.Validate()
	}
}

func BenchmarkConfigValidateWithErrors(b *testing.B) {
	config := config.Config{
		Host:     "",
		Port:     0,
		Username: "",
		Database: "",
	}

	b.ResetTimer()

	for range b.N {
		_ = config.Validate()
	}
}
