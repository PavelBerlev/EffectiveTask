package config

import (
	"os"
)

type Config struct {
	AppPort string
	DBHost string
	DBUser string
	DBPassword string
	DBPort string
	DBName string
}

func LoadConfig() (*Config, error){
	return &Config{
		AppPort: os.Getenv("APP_PORT"),
		DBHost: os.Getenv("DB_HOST"),
		DBUser: os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBPort: os.Getenv("DB_PORT"),
		DBName: os.Getenv("DB_NAME"),
	}, nil
}