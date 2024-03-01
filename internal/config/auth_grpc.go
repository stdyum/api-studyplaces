package config

import (
	"github.com/stdyum/api-common/proto/impl/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthGRpcConfig struct {
	URL string `env:"URL"`
}

func ConnectToAuthServer(config AuthGRpcConfig) (auth.AuthClient, error) {
	conn, err := grpc.Dial(config.URL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return auth.NewAuthClient(conn), nil
}
