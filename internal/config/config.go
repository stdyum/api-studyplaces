package config

import (
	"github.com/stdyum/api-common/env"
	"github.com/stdyum/api-common/server"
)

var Config Model

type Model struct {
	Ports    server.PortConfig `env:"PORT"`
	Database DatabaseConfig    `env:"DATABASE"`
	AuthGRpc AuthGRpcConfig    `env:"GRPC_AUTH"`
}

func init() {
	err := env.Fill(&Config)
	if err != nil {
		panic("cannot fill config: " + err.Error())
	}
}
