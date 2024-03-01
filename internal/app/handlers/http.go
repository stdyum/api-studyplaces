package handlers

import (
	"github.com/stdyum/api-common/hc"
	confHttp "github.com/stdyum/api-common/http"
	"github.com/stdyum/api-studyplaces/internal/app/controllers"
)

type HTTP interface {
	confHttp.Routes

	GetStudyPlaces(ctx *hc.Context)
	GetStudyPlaceById(ctx *hc.Context)
	RegisterStudyPlace(ctx *hc.Context)
	CloseStudyPlaceById(ctx *hc.Context)

	GetUserEnrollments(ctx *hc.Context)
	Enroll(ctx *hc.Context)
	WithdrawEnrollmentById(ctx *hc.Context)
	SetEnrollmentAcceptance(ctx *hc.Context)
	SetEnrollmentBlocked(ctx *hc.Context)

	GetEnrollmentPreferences(ctx *hc.Context)
	UpdateEnrollmentPreferences(ctx *hc.Context)
}

type http struct {
	controller controllers.Controller
}

func NewHTTP(controller controllers.Controller) HTTP {
	return &http{
		controller: controller,
	}
}
