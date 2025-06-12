package config

import (
	"log"
	"os"
)

type Config struct {
	URL          string
	PORT         string
	TOKEN_SECRET string
	ORIGIN_URL   string
}

func NewConfig() *Config {
	databaseURL, ok := getEnvVar("DATABASE_URL")
	if !ok {
		log.Fatal("DATABASE_URL not found in environments")
	}

	port, ok := getEnvVar("PORT")
	if !ok {
		log.Fatal("PORT not found in environments")
	}

	tokenSecret, ok := getEnvVar("TOKEN_SECRET")
	if !ok {
		log.Fatal("TOKEN_SECRET not found in environments")
	}

	originURL, ok := getEnvVar("ORIGIN_URL")
	if !ok {
		log.Fatal("ORIGIN_URL not found in environments")
	}

	return &Config{
		URL:          databaseURL,
		PORT:         port,
		TOKEN_SECRET: tokenSecret,
		ORIGIN_URL:   originURL,
	}
}

func getEnvVar(key string) (string, bool) {
	value := os.Getenv(key)
	if value == "" {
		return "", false
	}

	return value, true
}
