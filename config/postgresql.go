package config

// PostgreSQLConfig creates a PostgreSQL-specific configuration
func PostgreSQLConfig(host string, port int, username, password, database string) Config {
	cfg := &config{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		Database: database,
	}

	return cfg
}
