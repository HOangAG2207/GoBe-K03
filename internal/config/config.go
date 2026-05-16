package config

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// ===== Config structs =====

type Config struct {
	App   AppConfig   `mapstructure:"app"`
	Redis RedisConfig `mapstructure:"redis"`
}

type AppConfig struct {
	Port        string `mapstructure:"port"`
	ServiceName string `mapstructure:"service_name"`
	InstanceID  string `mapstructure:"instance_id"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	Prefix   string `mapstructure:"prefix"`
}

// ===== Load config =====

func Load() *Config {
	v := viper.New()

	// Load .env
	v.SetConfigFile(findEnvFile())
	v.SetConfigType("env")

	// Hỗ trợ ENV: APP.PORT ↔ APP_PORT
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		log.Printf("skip .env: %v", err)
	}

	// Default values
	// v.SetDefault("app.port", "8080")

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		log.Fatalf("cannot parse config: %v", err)
	}

	return &cfg
}

// ===== Find .env =====

func findEnvFile() string {
	dir, err := os.Getwd()
	if err != nil {
		return ".env"
	}

	for {
		path := filepath.Join(dir, ".env")
		if _, err := os.Stat(path); err == nil {
			return path
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	return ".env"
}
