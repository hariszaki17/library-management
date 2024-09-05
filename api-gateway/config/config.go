package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	UserRPCAddress string
	GRPCPort       string
	APPPort        string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	return &Config{
		GRPCPort:       getEnv("GRPC_PORT", "50051"),
		APPPort:        getEnv("APP_PORT", "8080"),
		UserRPCAddress: getEnv("USER_RPC_ADDRESS", "user-service:50051"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
