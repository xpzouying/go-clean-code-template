package cli

import (
	cli "github.com/urfave/cli/v2"
)

var (
	// listenAddr is the address to bind the local server to.
	listenAddr string

	// logLevel is the log level. debug, info, ...
	logLevel string

	// db dsn: default: sqlite3://:memory:
	dbConnStr string

	// configPath is the path to the config file. default: empty string, not start from config file.
	configPath string
)

var RootFlags = []cli.Flag{
	&cli.StringFlag{
		Name:        "listen",
		Aliases:     []string{"l"},
		Usage:       "Bind the local server to this address.",
		Destination: &listenAddr,
		Required:    false,
		Value:       "127.0.0.1:8080",
	},

	&cli.StringFlag{
		Name:        "log_level",
		Aliases:     []string{"ll"},
		Usage:       "Set the log level.",
		Destination: &logLevel,
		Required:    false,
		Value:       "debug",
	},

	&cli.StringFlag{
		Name:        "db",
		Usage:       "Set the database connection string.",
		Destination: &dbConnStr,
		Required:    false,
		Value:       "sqlite3://:memory:",
	},

	&cli.StringFlag{
		Name:        "config",
		Aliases:     []string{"c"},
		Usage:       "Set the path to the config file.",
		Destination: &configPath,
		Required:    false,
		Value:       "",
	},
}
