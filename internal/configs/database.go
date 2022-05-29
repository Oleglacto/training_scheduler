package configs

import (
	"fmt"
	"net/url"
)

type Database struct {
	Host     string `env:"DB_HOST"`
	Port     int    `env:"DB_PORT"`
	Schema   string `env:"DB_SCHEMA"`
	Database string `env:"DB_NAME"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
}

func (cfg Database) Dsn() string {
	return (&url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(cfg.User, cfg.Password),
		Host:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Path:     cfg.Database,
		RawQuery: fmt.Sprintf("sslmode=disable&?search_path=%s", cfg.Schema),
	}).String()
}

// postgresql://zebra:password@localhost/training_scheduler?search_path=main
