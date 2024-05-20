package handlers

import (
	netHttp "net/http"

	"github.com/stdyum/api-common/hc"
	"github.com/stdyum/api-studyplaces/internal/app/dto"
)

func (h *http) GetStudyPlaces(ctx *hc.Context) {
	query := ctx.PaginationQuery()

	studyPlaces, err := h.controller.GetStudyPlaces(ctx, query)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(netHttp.StatusOK, studyPlaces)
}

func (h *http) GetStudyPlaceById(ctx *hc.Context) {
	id, err := ctx.UUIDParam("id")
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	studyPlace, err := h.controller.GetStudyPlaceById(ctx, id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(netHttp.StatusOK, studyPlace)
}

func (h *http) RegisterStudyPlace(ctx *hc.Context) {
	user := ctx.User()

	var requestDTO dto.CreateStudyPlaceRequestDTO
	if err := ctx.BindJSON(&requestDTO); err != nil {
		_ = ctx.Error(err)
		return
	}

	studyPlace, err := h.controller.RegisterStudyPlace(ctx, user, requestDTO)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(netHttp.StatusCreated, studyPlace)
}

func (h *http) CloseStudyPlaceById(ctx *hc.Context) {
	user := ctx.User()

	id, err := ctx.UUIDParam("id")
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	if err = h.controller.CloseStudyPlaceById(ctx, user, id); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(netHttp.StatusNoContent)
}

func (h *http) GetUserEnrollments(ctx *hc.Context) {
	user := ctx.User()
	query := ctx.PaginationQuery()

	enrollments, err := h.controller.GetUserEnrollments(ctx, user, query)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(netHttp.StatusOK, enrollments)
}

func (h *http) GetStudyPlaceEnrollments(ctx *hc.Context) {
	user := ctx.User()
	query := ctx.PaginationQuery()
	studyPlaceId := ctx.StudyPlaceId()
	accepted, err := ctx.QueryBool("accepted")
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	token := ctx.GetHeader("Authorization")

	enrollments, err := h.controller.GetEnrollmentRequests(ctx, studyPlaceId, token, user, query, accepted)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(netHttp.StatusOK, enrollments)
}

func (h *http) GetUserEnrollmentById(ctx *hc.Context) {
	user := ctx.User()
	id, err := ctx.UUIDParam("id")
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	enrollments, err := h.controller.GetUserEnrollmentById(ctx, user, id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(netHttp.StatusOK, enrollments)
}

func (h *http) Enroll(ctx *hc.Context) {
	user := ctx.User()

	var requestDTO dto.EnrollRequestDTO
	if err := ctx.BindJSON(&requestDTO); err != nil {
		_ = ctx.Error(err)
		return
	}

	enroll, err := h.controller.Enroll(ctx, user, requestDTO)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(netHttp.StatusCreated, enroll)
}

func (h *http) SetEnrollmentAcceptance(ctx *hc.Context) {
	user := ctx.User()
	id, err := ctx.UUIDParam("id")
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	var requestDTO dto.SetEnrollmentAcceptanceRequestDTO
	if err = ctx.BindJSON(&requestDTO); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err = h.controller.SetEnrollmentAcceptance(ctx, user, id, requestDTO); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(netHttp.StatusNoContent)
}

func (h *http) WithdrawEnrollmentById(ctx *hc.Context) {
	user := ctx.User()
	id, err := ctx.UUIDParam("id")
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	if err = h.controller.WithdrawEnrollmentById(ctx, user, id); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(netHttp.StatusNoContent)
}

func (h *http) PatchStudyPlaceEnrollment(ctx *hc.Context) {
	user := ctx.User()
	id, err := ctx.UUIDParam("id")
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	var request dto.UpdateStudyPlaceEnrollmentRequestDTO
	if err = ctx.BindJSON(&request); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err = h.controller.UpdateStudyPlaceEnrollment(ctx, user, id, request); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(netHttp.StatusNoContent)
}

func (h *http) SetEnrollmentBlocked(ctx *hc.Context) {
	user := ctx.User()
	id, err := ctx.UUIDParam("id")
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	var requestDTO dto.SetEnrollmentBlockedRequestDTO
	if err = ctx.BindJSON(&requestDTO); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err = h.controller.SetEnrollmentBlocked(ctx, user, id, requestDTO); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(netHttp.StatusNoContent)
}

func (h *http) GetEnrollmentPreferences(ctx *hc.Context) {
	user := ctx.User()
	id := ctx.StudyPlaceId()

	preferences, err := h.controller.GetEnrollmentPreferences(ctx, user, id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(netHttp.StatusOK, preferences)
}

func (h *http) UpdateEnrollmentPreferences(ctx *hc.Context) {
	user := ctx.User()
	id := ctx.StudyPlaceId()

	var requestDTO dto.UpdatePreferencesRequestDTO
	if err := ctx.BindJSON(&requestDTO); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err := h.controller.UpdateEnrollmentPreferences(ctx, user, id, requestDTO); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(netHttp.StatusNoContent)
}
