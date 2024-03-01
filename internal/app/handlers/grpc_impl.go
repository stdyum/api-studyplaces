package handlers

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/grpc"
	"github.com/stdyum/api-common/proto/impl/studyplaces"
	"github.com/stdyum/api-studyplaces/internal/app/dto"
)

func (h *gRPC) Auth(ctx context.Context, token *studyplaces.EnrollmentToken) (*studyplaces.EnrollmentUser, error) {
	studyPlaceId, err := uuid.Parse(token.StudyPlaceId)
	if err != nil {
		return nil, grpc.ConvertError(err)
	}

	request := dto.AuthRequestDTO{
		Token:        token.Token,
		StudyPlaceId: studyPlaceId,
	}
	response, err := h.controller.Auth(ctx, request)
	if err != nil {
		return nil, grpc.ConvertError(err)
	}

	role := studyplaces.EnrollmentRole_value[strings.ToUpper(string(response.Enrollment.Role))]

	return &studyplaces.EnrollmentUser{
		Id:            response.Id.String(),
		Login:         response.Login,
		PictureUrl:    response.PictureUrl,
		Email:         response.Email,
		VerifiedEmail: response.VerifiedEmail,
		Enrollment: &studyplaces.Enrollment{
			Id:           response.Enrollment.ID.String(),
			UserId:       response.Enrollment.UserId.String(),
			StudyPlaceId: response.Enrollment.StudyPlaceId.String(),
			UserName:     response.Enrollment.UserName,
			Role:         studyplaces.EnrollmentRole(role),
			TypeId:       response.Enrollment.TypeId.String(),
			Permissions:  response.Enrollment.Permissions,
			Accepted:     response.Enrollment.Accepted,
		},
	}, nil
}
