package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	Host     string
	User     string
	Port     string
	Name     string
	Password string
	Driver   string
}

type APIConfig struct {
	APIPort string
}

type Config struct {
	DbConfig
	APIConfig
}

func (c *Config) LoadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return nil
}

func (c *Config) ReadConfig() error {
	// ini nanti jangan hardcode disini, gunakan env
	c.DbConfig = DbConfig{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Driver:   os.Getenv("DB_DRIVER"),
	}

	c.APIConfig = APIConfig{
		APIPort: "8888",
	}

	if c.DbConfig.Host == "" || c.DbConfig.User == "" || c.DbConfig.Port == "" || c.DbConfig.Name == "" {
		return fmt.Errorf("konfigurasi tidak boleh kosong")
	}
	return nil
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cfg.ReadConfig()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
