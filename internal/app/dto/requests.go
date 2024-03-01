package dto

import (
	"github.com/google/uuid"
	"github.com/stdyum/api-common/models"
)

type AuthRequestDTO struct {
	Token        string    `json:"token"`
	StudyPlaceId uuid.UUID `json:"studyPlaceId"`
}

type CreateStudyPlaceRequestDTO struct {
	Title         string `json:"title"`
	AdminUserName string `json:"adminUserName"`
}

type EnrollRequestDTO struct {
	StudyPlaceId uuid.UUID   `json:"studyPlaceId"`
	UserName     string      `json:"userName"`
	Role         models.Role `json:"role"`
	TypeId       uuid.UUID   `json:"typeId"`
}

type SetEnrollmentAcceptanceRequestDTO struct {
	Accepted bool `json:"accepted"`
}

type SetEnrollmentBlockedRequestDTO struct {
	Blocked bool `json:"blocked"`
}

type UpdatePreferencesRequestDTO struct {
	Group       models.PreferenceGroup `json:"group"`
	Preferences any                    `json:"preferences"`
}
