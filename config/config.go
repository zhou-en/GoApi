package config

import (
	"os"
)

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Charset  string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Username: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			Charset:  "utf8",
		},
	}
}
