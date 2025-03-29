package config

import (
	"strings"

	"github.com/joho/godotenv"
)

type (
	// Config holds the application configuration.
	Config struct {
		app   *app
		db    *db
		token *token
	}

	// app holds the application configuration.
	app struct {
		Name string
		Env  string
		Host string
		Port string
	}

	// db holds the database configuration.
	db struct {
		Type     string
		Host     string
		Port     string
		User     string
		Password string
		Name     string
		URI      string
	}

	// token holds the token configuration.
	token struct {
		Type   string
		Secret string
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
		app: &app{
			Name: getEnv("APP_NAME"),
			Env:  getEnv("APP_ENV"),
			Host: getEnv("APP_HOST"),
			Port: getEnv("APP_PORT"),
		},
		db: &db{
			Type:     getEnv("DB_TYPE"),
			Host:     getEnv("DB_HOST"),
			Port:     getEnv("DB_PORT"),
			User:     getEnv("DB_USER"),
			Password: getEnv("DB_PASSWORD"),
			Name:     getEnv("DB_NAME"),
			URI:      getEnv("DB_URI"),
		},
		token: &token{
			Type:   getEnv("TOKEN_TYPE"),
			Secret: getEnv("TOKEN_SECRET"),
		},
	}, nil
}

func (c *Config) App() app {
	return *c.app
}

func (c *Config) DB() db {
	return *c.db
}

func (c *Config) Token() token {
	return *c.token
}
