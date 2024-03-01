package internal

import (
	"github.com/stdyum/api-common/grpc/clients"
	"github.com/stdyum/api-studyplaces/internal/app"
	"github.com/stdyum/api-studyplaces/internal/config"
)

func App() error {
	db, err := config.ConnectToDatabase(config.Config.Database)
	if err != nil {
		return err
	}

	authServer, err := config.ConnectToAuthServer(config.Config.AuthGRpc)
	if err != nil {
		return err
	}
	clients.AuthGRpcClient = authServer

	routes, err := app.New(db, authServer)
	if err != nil {
		return err
	}

	routes.Ports = config.Config.Ports
	return routes.Run()
}
