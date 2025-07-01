package connections_test

import (
	"context"
	"testing"

	"github.com/DucTran999/dbkit/config"
	"github.com/DucTran999/dbkit/connections"
	"github.com/stretchr/testify/require"
)

func TestMySQLConnection(t *testing.T) {
	pgConf := config.MySQLConfig{
		Config: config.Config{
			Host:     "localhost",
			Port:     3306,
			Username: "test",
			Password: "test",
			Database: "dbkit_test",
			TimeZone: "Asia/Ho_Chi_Minh",
		},
	}

	conn, err := connections.NewMySQLConnection(pgConf)
	require.NoError(t, err)

	// Test Ping to DB
	err = conn.Ping(context.Background())
	require.NoError(t, err)

	// verify db instance
	db := conn.DB()
	require.NotNil(t, db)

	conn.Close()

	// Test Ping to DB
	err = conn.Ping(context.Background())
	require.ErrorContains(t, err, "database is closed")
}

func TestMySQLConnectionFailed(t *testing.T) {
	pgConf := config.MySQLConfig{
		Config: config.Config{
			Port:     6432,
			Username: "test",
			Password: "test",
			Database: "dbkit_test",
			TimeZone: "Asia/Ho_Chi_Minh",
		},
	}

	// test config missing host
	conn, err := connections.NewMySQLConnection(pgConf)
	require.ErrorIs(t, err, config.ErrMissingHost)

	// Test connection failed cause wrong port
	pgConf.Host = "localhost"
	conn, err = connections.NewMySQLConnection(pgConf)
	require.ErrorContains(t, err, "connection refused")
	require.Nil(t, conn)
}
