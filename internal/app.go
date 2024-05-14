package internal

import (
	"github.com/stdyum/api-common/grpc/clients"
	"github.com/stdyum/api-studyplaces/internal/app"
	"github.com/stdyum/api-studyplaces/internal/config"
	"github.com/stdyum/api-studyplaces/internal/modules/types_registry"
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

	typesRegistryClient, err := config.ConnectToTypesRegistryServer(config.Config.TypesRegistryGRpc)
	if err != nil {
		return err
	}

	typesRegistry := types_registry.New(typesRegistryClient)

	routes, err := app.New(db, authServer, typesRegistry)
	if err != nil {
		return err
	}

	routes.Ports = config.Config.Ports
	return routes.Run()
}
