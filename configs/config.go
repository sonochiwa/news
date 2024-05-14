package configs

import (
	"errors"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var GlobalConfig Config

type Config struct {
	Mode     string
	Auth     Auth
	Server   Server
	Postgres Postgres
}

type Auth struct {
	SecretKey string
}

type Server struct {
	Protocol string
	Host     string
	Port     string
}

type Postgres struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
	SSLMode  string
}

func init() {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(fmt.Errorf("viper.ReadInConfig: %v", err))
	}

	GlobalConfig = New()
}

func New() Config {
	config := Config{
		Mode: getString("MODE"),
		Auth: Auth{
			SecretKey: getString("SECRET_KEY"),
		},
		Server: Server{
			Protocol: getString("SRV_PROTOCOL"),
			Host:     getString("SRV_HOST"),
			Port:     getString("SRV_PORT"),
		},
		Postgres: Postgres{
			Username: getString("PG_USERNAME"),
			Password: getString("PG_PASSWORD"),
			Host:     getString("PG_HOST"),
			Port:     getString("PG_PORT"),
			DBName:   getString("PG_DBNAME"),
			SSLMode:  getString("PG_SSL_MODE"),
		},
	}

	if len(notSetKeys) != 0 {
		log.Fatal(errors.New(fmt.Sprintf("%s must be in configuration", notSetKeys)))
	}

	return config
}

var notSetKeys []string

func getString(key string) string {
	value := viper.GetString(key)
	if !viper.IsSet(key) || value == "" {
		notSetKeys = append(notSetKeys, key)
		return ""
	}

	return value
}
