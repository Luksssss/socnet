package main

import (
	_ "github.com/joho/godotenv/autoload" // load env vars from .env file

	"github.com/kelseyhightower/envconfig"
)

// Config represents application configuration.
type Config struct {
	// Development mode enables dev-only features
	Development bool `envconfig:"DEVELOPMENT" default:"false"`

	// Port to listen on
	Port     int `envconfig:"PORT" default:"8080"`
	HttpPort int `envconfig:"PORT" default:"8081"`

	// PostgreSQL server
	PGHost     string `envconfig:"POSTGRES_HOST" required:"true"`
	PGPort     uint16 `envconfig:"POSTGRES_PORT" required:"true"`
	PGDatabase string `envconfig:"POSTGRES_DATABASE" required:"true"`
	PGParams   string `envconfig:"POSTGRES_PARAMS"`
	PGUsername string `envconfig:"POSTGRES_USERNAME" required:"true"`
	PGPassword string `envconfig:"POSTGRES_PASSWORD" required:"true"`
}

func readConfig() (*Config, error) {
	cfg := &Config{}
	//godotenv.Load("../build/.env")
	//godotenv.Load("/.env")
	err := envconfig.Process("", cfg)
	return cfg, err
}
