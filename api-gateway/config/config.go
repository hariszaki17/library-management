package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	UserRPCAddress   string `env:"USER_RPC_ADDRESS" env-default:"user-service:50051"`
	BookRPCAddress   string `env:"BOOK_RPC_ADDRESS" env-default:"book-service:50051"`
	AuthorRPCAddress string `env:"AUTHOR_RPC_ADDRESS" env-default:"author-service:50051"`
	GRPCPort         string `env:"GRPC_PORT" env-default:"50051"`
	APPPort          string `env:"APP_PORT" env-default:"8080"`
	ServiceName      string `env:"SERVICE_NAME" env-default:"api-gateway"`
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
