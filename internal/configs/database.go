package configs

import "fmt"

type DataBaseConfig struct {
	Host     string `env:"DB_HOST"`
	Port     int    `env:"DB_PORT"`
	Schema   string `env:"DB_SCHEMA"`
	Database string `env:"DB_NAME"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
}

func (cfg DataBaseConfig) GetConnectionUrl() string {
	return fmt.Sprintf("postgresql://%s:%s@%s/%s?search_path=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Database,
		cfg.Schema,
	)
}

// postgresql://zebra:password@localhost/training_scheduler?search_path=main
