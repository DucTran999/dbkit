# DBKit - Database Abstraction Layer for Go

[![Go Report Card](https://goreportcard.com/badge/github.com/DucTran999/dbkit)](https://goreportcard.com/report/github.com/DucTran999/dbkit)
[![Go](https://img.shields.io/badge/Go-1.23-blue?logo=go)](https://golang.org)
[![codecov](https://codecov.io/gh/DucTran999/dbkit/graph/badge.svg?token=5XBMMBKCPD)](https://codecov.io/gh/DucTran999/dbkit)
[![Known Vulnerabilities](https://snyk.io/test/github/ductran999/dbkit/badge.svg)](https://snyk.io/test/github/ductran999/dbkit)
[![License](https://img.shields.io/github/license/DucTran999/dbkit)](LICENSE)

## Overview

**DBKit** is a lightweight and extensible database abstraction layer for Go. It simplifies working with multiple SQL database dialects (e.g. PostgreSQL, ClickHouse, MySQL) using a unified interface and configuration-driven setup. Built on top of [GORM](https://gorm.io), DBKit is ideal for applications that require flexible, decoupled database initialization and switching.

## Features

- ✅ Unified interface for multiple SQL databases
- 🔌 Supports PostgreSQL, MySQL, and ClickHouse
- 🧪 Easy to extend with custom dialects
- 📦 Designed for modular use in microservices or monoliths
- 🧹 Clean, testable code with coverage reporting

## Installation

```bash
go get github.com/DucTran999/dbkit
```

## Quick start

### Basic Connection

Here's how to establish a connection to different databases:

### PostgreSQL

```go
package main

import (
	"log"

	"github.com/DucTran999/dbkit"
	"github.com/DucTran999/dbkit/config"
)

func main() {
	pgConf := config.PostgreSQLConfig{
        Config: config.Config{
            Host:     "your_host",
            Port:     5432,
            Username: "your_username",
            Password: "your_password",
            Database: "your_database",
            Timezone: "UTC",
        },
        SSLMode: config.PgSSLDisable,
    }

	conn, err := dbkit.NewPostgreSQLConnection(pgConf)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()


    // Test the connection
    if err := conn.Ping(); err != nil {
        log.Fatalf("Failed to ping database: %v", err)
    }

    log.Println("Successfully connected to PostgreSQL!")
}
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License.
