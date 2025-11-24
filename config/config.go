package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Config struct {
	HTTP struct {
		Port string `env:"HTTP_PORT" env-default:"8080"`
	}
	DB struct {
		Host     string `env:"DB_HOST" env-default:"localhost"`
		Port     string `env:"DB_PORT" env-default:"5432"`
		User     string `env:"DB_USER" env-default:"postgres"`
		Password string `env:"DB_PASSWORD" env-default:"postgres"`
		Name     string `env:"DB_NAME" env-default:"edu_adapt"`
	}
}

func MustLoad() *Config {
	var cfg Config
	if err := cleanenv.ReadConfig(".env", &cfg); err != nil {
		if err := cleanenv.ReadEnv(&cfg); err != nil {
			log.Fatalf("cannot read config: %s", err)
		}
	}
	return &cfg
}
