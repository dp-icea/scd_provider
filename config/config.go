package config

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var lock = &sync.Mutex{}

type Config struct {
	HostUrl string
	DssUrl  string
	AuthUrl string
	ApiKey  string
}

func newInstance() *Config {
	return &Config{
		HostUrl: getEnv("HOST_URL", ""),
		DssUrl:  getEnv("DSS_URL", ""),
		AuthUrl: getEnv("AUTH_URL", ""),
		ApiKey:  getEnv("API_KEY", "brutm"),
	}
}

var configInstance *Config

func initEnv() {
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
}

func GetGlobalConfig() *Config {
	if configInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if configInstance == nil {
			initEnv()
			configInstance = newInstance()
		}
	}
	return configInstance
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
