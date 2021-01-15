package config

import (
	"github.com/joho/godotenv"
	"os"
)

const (
	defaultServerPort         = 8081
	defaultJWTExpirationHours = 72
)

type Config struct {
	// the server port. Defaults to 8080
	ServerPort int `yaml:"server_port" env:"SERVER_PORT"`
	// the data source name (DSN) for connecting to the database. required.
	DSN string `yaml:"dsn" env:"DSN,secret"`
	// JWT signing key. required.
	JWTSigningKey string `yaml:"jwt_signing_key" env:"JWT_SIGNING_KEY,secret"`
	// JWT expiration in hours. Defaults to 72 hours (3 days)
	JWTExpiration int `yaml:"jwt_expiration" env:"JWT_EXPIRATION"`
}

func Load(filename string) (*Config, error) {
	c := Config{
		ServerPort:    defaultServerPort,
		JWTExpiration: defaultJWTExpirationHours,
	}
	_ = godotenv.Load(filename)
	c.DSN = os.Getenv("DSN")
	c.JWTSigningKey = os.Getenv("JWT_SIGNING_KEY")

	return &c, nil
}
