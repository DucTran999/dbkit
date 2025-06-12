package tests

import (
	"testing"

	"github.com/DucTran999/dbkit"
	"github.com/DucTran999/dbkit/config"
	"github.com/stretchr/testify/require"
)

func TestPostgreSQL(t *testing.T) {
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

	defer conn.Close()
}
