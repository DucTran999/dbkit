package dbkit_test

import (
	"testing"

	"github.com/DucTran999/dbkit"
	"github.com/DucTran999/dbkit/config"
	"github.com/stretchr/testify/require"
)

func TestPostgreSQLConnection(t *testing.T) {
	pgConf := config.PostgreSQLConfig{
		Config: config.Config{
			Host:     "localhost",
			Port:     5432,
			Username: "test",
			Password: "test",
			Database: "dbkit_test",
			Timezone: "Asia/Ho_Chi_Minh",
		},
		SSLMode: config.PgSSLDisable, // or whatever the correct SSLMode constant is
	}

	conn, err := dbkit.NewPostgreSQLConnection(pgConf)
	require.NoError(t, err)

	// Test Ping to DB
	err = conn.Ping()
	require.NoError(t, err)

	// verify db instance
	db := conn.DB()
	require.NotNil(t, db)

	conn.Close()

	// Test Ping to DB
	err = conn.Ping()
	require.ErrorContains(t, err, "database is closed")
}

func TestPostgreSQLConnectionFailed(t *testing.T) {
	pgConf := config.PostgreSQLConfig{
		Config: config.Config{
			Port:     5433,
			Username: "test",
			Password: "test",
			Database: "dbkit_test",
			Timezone: "Asia/Ho_Chi_Minh",
		},
		SSLMode: config.PgSSLDisable, // or whatever the correct SSLMode constant is
	}

	// test config missing host
	conn, err := dbkit.NewPostgreSQLConnection(pgConf)
	require.ErrorIs(t, err, config.ErrMissingHost)

	// Test connection failed cause wrong port
	pgConf.Host = "localhost"
	conn, err = dbkit.NewPostgreSQLConnection(pgConf)
	require.ErrorContains(t, err, "connection refused")
	require.Nil(t, conn)
}
