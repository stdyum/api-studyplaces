package handlers

import (
	"github.com/stdyum/api-common/hc"
	"github.com/stdyum/api-common/http/middlewares"
	"github.com/stdyum/api-common/proto/impl/studyplaces"
	"google.golang.org/grpc"
)

func (h *http) ConfigureRoutes() *hc.Engine {
	engine := hc.New()
	engine.Use(hc.Recovery())

	group := engine.Group("api/v1", hc.Logger(), middlewares.ErrorMiddleware())
	{
		group.GET("studyplaces", middlewares.PaginationMiddleware(10), h.GetStudyPlaces)
		group.GET("studyplaces/:id", h.GetStudyPlaceById)

		withAuth := group.Group("", middlewares.AuthMiddleware())
		{
			studyplacesGroup := withAuth.Group("studyplaces")
			{
				studyplacesGroup.POST("", middlewares.AuthMiddleware(), h.RegisterStudyPlace)
				studyplacesGroup.DELETE(":id", middlewares.AuthMiddleware(), h.CloseStudyPlaceById)
			}

			enrollmentsGroup := withAuth.Group("enrollments")
			{
				enrollmentsGroup.GET("", middlewares.PaginationMiddleware(10), h.GetUserEnrollments)
				enrollmentsGroup.GET(":id", h.GetUserEnrollmentById)
				enrollmentsGroup.POST("", h.Enroll)
				enrollmentsGroup.DELETE(":id", h.WithdrawEnrollmentById)
				enrollmentsGroup.PUT("acceptance", h.SetEnrollmentAcceptance)
				enrollmentsGroup.PUT("blocked", h.SetEnrollmentBlocked)
			}

			preferencesGroup := withAuth.Group("preferences")
			{
				preferencesGroup.GET(":id", h.GetEnrollmentPreferences)
				preferencesGroup.PUT(":id", h.UpdateEnrollmentPreferences)
			}
		}
	}

	return engine
}

func (h *gRPC) ConfigureRoutes() *grpc.Server {
	grpcServer := grpc.NewServer()
	studyplaces.RegisterStudyplacesServer(grpcServer, h)
	return grpcServer
}
