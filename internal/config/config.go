package config

import (
	"github.com/spf13/viper"
)

const (
	defaultServerPort         = 8081
	defaultJWTExpirationHours = 72
)

type Config struct {
	// the server port. Defaults to 8081
	ServerPort int `yaml:"server_port" env:"SERVER_PORT"`
	// the data source name (DSN) for connecting to the database. required.
	DSN string `yaml:"dsn" env:"DSN,secret"`
	// JWT signing key. required.
	JWTSigningKey string `yaml:"jwt_signing_key" env:"JWT_SIGNING_KEY,secret"`
	// JWT expiration in hours. Defaults to 72 hours (3 days)
	JWTExpiration int `yaml:"jwt_expiration" env:"JWT_EXPIRATION"`
}

func Load(file string) (*Config, error) {
	c := Config{
		ServerPort:    defaultServerPort,
		JWTExpiration: defaultJWTExpirationHours,
	}
	viper.AutomaticEnv()
	viper.SetConfigFile(file)
	_ = viper.ReadInConfig()
	c.DSN = viper.GetString("DSN")
	c.JWTSigningKey = viper.GetString("JWT_SIGNING_KEY")
	return &c, nil
}
