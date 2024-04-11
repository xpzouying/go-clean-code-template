package cli

import (
	"os"

	"github.com/xpzouying/go-clean-code-template/internal/config"
	"github.com/xpzouying/go-clean-code-template/log"

	"github.com/pkg/errors"
)

func InitConfigAndComponents() (*config.Config, error) {

	cfg, err := loadConfig()
	if err != nil {
		return nil, err
	}

	if err := initLogger(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func loadConfig() (*config.Config, error) {

	if configPath != "" {

		f, err := os.Open(configPath)
		if err != nil {
			return nil, errors.Wrap(err, "failed to open config file")
		}
		defer f.Close()

		return config.NewConfig(f)

	} else {

		cfg := &config.Config{
			ListenAddr: listenAddr,
			LogLevel:   logLevel,
			DBConnStr:  dbConnStr,
		}

		return cfg, nil
	}
}

func initLogger(cfg *config.Config) error {
	if err := log.InitGlobalLogger(cfg.LogLevel); err != nil {
		return err
	}
	return nil
}
