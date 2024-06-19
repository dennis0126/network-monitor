package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"log"
)

type config struct {
	Port     int    `env:"PORT"`
	DbString string `env:"DB_STRING"`
}

func InitConfig() config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}
