package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DBHost      string `env:"DB_HOST" env-default:"postgres"`
	DBUser      string `env:"DB_USER" env-default:"postgres"`
	DBPassword  string `env:"DB_PASSWORD" env-default:""`
	DBName      string `env:"DB_NAME" env-default:"postgres"`
	DBPort      string `env:"DB_PORT" env-default:"5432"`
	GRPCPort    string `env:"GRPC_PORT" env-default:"50051"`
	RedistPort  string `env:"REDIS_PORT" env-default:"6379"`
	RedistHost  string `env:"REDIS_HOST" env-default:"redis"`
	ServiceName string `env:"SERVICE_NAME" env-default:"book-service"`
}

var (
	Data Config
)

func Load(path string) error {
	err := cleanenv.ReadConfig(path, &Data)
	if err != nil {
		return err
	}
	return nil
}
