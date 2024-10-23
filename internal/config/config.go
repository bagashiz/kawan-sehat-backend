package config

import (
	"strings"

	"github.com/joho/godotenv"
)

type (
	// Config holds the application configuration.
	Config struct {
		App   *App
		DB    *DB
		Token *Token
	}

	// App holds the application configuration.
	App struct {
		Name string
		Env  string
		Host string
		Port string
	}

	// DB holds the database configuration.
	DB struct {
		Type     string
		Host     string
		Port     string
		User     string
		Password string
		Name     string
		URI      string
	}

	// Token holds the token configuration.
	Token struct {
		Type     string
		Secret   string
		Duration string
	}
)

// file is the name of the configuration file.
const file = ".env"

// New returns a new Config instance.
func New(getEnv func(string) string) (*Config, error) {
	env := strings.ToLower(getEnv("APP_ENV"))
	if env != "production" {
		if err := godotenv.Load(file); err != nil {
			return nil, err
		}
	}

	return &Config{
		App: &App{
			Name: getEnv("APP_NAME"),
			Env:  getEnv("APP_ENV"),
			Host: getEnv("APP_HOST"),
			Port: getEnv("APP_PORT"),
		},
		DB: &DB{
			Type:     getEnv("DB_TYPE"),
			Host:     getEnv("DB_HOST"),
			Port:     getEnv("DB_PORT"),
			User:     getEnv("DB_USER"),
			Password: getEnv("DB_PASSWORD"),
			Name:     getEnv("DB_NAME"),
			URI:      getEnv("DB_URI"),
		},
		Token: &Token{
			Type:     getEnv("TOKEN_TYPE"),
			Secret:   getEnv("TOKEN_SECRET"),
			Duration: getEnv("TOKEN_DURATION"),
		},
	}, nil
}
