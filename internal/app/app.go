package app

import (
	"database/sql"

	"github.com/stdyum/api-common/proto/impl/auth"
	"github.com/stdyum/api-common/server"
	"github.com/stdyum/api-studyplaces/internal/app/controllers"
	"github.com/stdyum/api-studyplaces/internal/app/errors"
	"github.com/stdyum/api-studyplaces/internal/app/handlers"
	"github.com/stdyum/api-studyplaces/internal/app/repositories"
	"github.com/stdyum/api-studyplaces/internal/modules/types_registry"
)

func New(database *sql.DB, authServer auth.AuthClient, registry types_registry.Controller) (server.Routes, error) {
	repo := repositories.New(database)
	authRepo := repositories.NewAuth(authServer)

	ctrl := controllers.New(repo, authRepo, registry)

	errors.Register()

	httpHndl := handlers.NewHTTP(ctrl)
	grpcHndl := handlers.NewGRPC(ctrl)

	routes := server.Routes{
		GRPC: grpcHndl,
		HTTP: httpHndl,
	}

	return routes, nil
}
