package config

import (
	"encoding/json"
	"io"
)

// Config is a struct that holds the configuration for the application.
type Config struct {
	// ListenAddr is the address to bind the local server to.
	ListenAddr string `json:"listen"`

	// LogLevel is the log level for the application.
	LogLevel string `json:"log_level"`

	// DBConnStr is the connection string for the database.
	//
	// Begain with the database type, followed by the connection string. For example:
	// - sqlite3://:memory: creates an in-memory database.
	// - sqlite3:///path/to/database.db creates a database file at the specified path.
	// - postgres://user:password@localhost:5432/dbname connects to a PostgreSQL database.
	DBConnStr string `json:"db"`
}

// NewConfig creates a new Config struct from the provided io.Reader.
// The io.Reader should contain the configuration in *json* format that can be parsed by the application.
func NewConfig(r io.Reader) (*Config, error) {

	var cfg Config
	if err := json.NewDecoder(r).Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
