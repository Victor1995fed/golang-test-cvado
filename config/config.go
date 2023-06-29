package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"log"
)

// Config Конфиги приложения
type Config struct {
	App      App
	Database Database
}

type Database struct {
	Host     string `env:"MYSQL_HOST"`
	Port     string `env:"MYSQL_TCP_PORT"`
	Database string `env:"MYSQL_DATABASE"`
	Username string `env:"DB_USERNAME"`
	Password string `env:"MYSQL_ROOT_PASSWORD"`
}

type App struct {
	Port int `env:"APP_PORT"`
}

func ParseConfig(path string) Config {
	err := godotenv.Load(path)
	if err != nil {
		log.Fatalf("Failed to load env file: %e", err)
	}

	cfg := Config{}
	err = env.Parse(&cfg)
	if err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}
	return cfg
}
