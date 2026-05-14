package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	App AppConfig
}

type AppConfig struct {
	Port string
}

func Load() *Config {
	viper.SetConfigFile(findEnvFile())
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: could not read .env file: %v", err)
	}

	return &Config{
		App: AppConfig{
			Port: viper.GetString("APP_PORT"),
		},
	}
}

// findEnvFile tìm file .env từ thư mục hiện tại lên đến root project
func findEnvFile() string {
	dir, _ := os.Getwd()
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
	return ".env" // fallback
}
