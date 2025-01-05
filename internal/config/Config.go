package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	ServerAddress string
	WorkerCount   int
}

func Load() *Config {
	WorkerCount, err := strconv.Atoi(getEnv("WORKER_COUNT", "5"))
	if err != nil {
		log.Fatalf("Invalid WORKER_COUNT: %v", err)
	}
	return &Config{
		ServerAddress: getEnv("SERVER_PORT", ":8080"),
		WorkerCount:   WorkerCount,
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
