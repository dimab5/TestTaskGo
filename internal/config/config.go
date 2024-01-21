package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env              string `yaml:"env"`
	Database         DatabaseConfig
	HttpServerConfig `yaml:"http_server"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
}

type HttpServerConfig struct {
	Address     string        `yaml:"address"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

func MustLoad() *Config {
	configPath := "C:\\prog\\TestTaskGo\\config\\local.yaml"
	os.Setenv("CONFIG_PATH", configPath)

	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {

		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}

func (config *Config) DatabaseConnectionString() string {
	return fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%d sslmode=%s",
		config.Database.User,
		config.Database.Password,
		config.Database.DBName,
		config.Database.Host,
		config.Database.Port,
		config.Database.SSLMode,
	)
}
