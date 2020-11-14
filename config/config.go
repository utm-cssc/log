package config

import "os"

// DBConfig - Base DB Config
type DBConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

func GetDBConfig() *DBConfig {
	return &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}
}
