package dto

import (
	"github.com/google/uuid"
	"github.com/stdyum/api-common/databases/pagination"
	"github.com/stdyum/api-common/models"
)

type AuthResponseDTO struct {
	Id            uuid.UUID                  `json:"id"`
	Login         string                     `json:"login"`
	PictureUrl    string                     `json:"pictureUrl"`
	Email         string                     `json:"email"`
	VerifiedEmail bool                       `json:"verifiedEmail"`
	Enrollment    EnrollmentsResponseItemDTO `json:"enrollment"`
}

type StudyPlacesResponseItemDTO struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
}

type StudyPlacesResponseDTO struct {
	Items    []StudyPlacesResponseItemDTO   `json:"items"`
	PerPage  int                            `json:"perPage"`
	Total    int                            `json:"total"`
	Page     int                            `json:"page"`
	Next     *pagination.CreatedAtPageQuery `json:"next"`
	Previous *pagination.CreatedAtPageQuery `json:"previous"`
}

type EnrollmentsResponseItemDTO struct {
	ID              uuid.UUID   `json:"id"`
	UserId          uuid.UUID   `json:"userId"`
	StudyPlaceId    uuid.UUID   `json:"studyPlaceId"`
	StudyPlaceTitle string      `json:"studyPlaceTitle"`
	UserName        string      `json:"userName"`
	Role            models.Role `json:"role"`
	TypeId          uuid.UUID   `json:"typeId"`
	Permissions     []string    `json:"permissions"`
	Accepted        bool        `json:"accepted"`
}

type CreateStudyPlaceResponseDTO struct {
	StudyPlace StudyPlacesResponseItemDTO `json:"studyPlace"`
	Enrollment EnrollmentsResponseItemDTO `json:"enrollment"`
}

type EnrollmentsResponseDTO struct {
	Items    []EnrollmentsResponseItemDTO   `json:"items"`
	PerPage  int                            `json:"perPage"`
	Total    int                            `json:"total"`
	Page     int                            `json:"page"`
	Next     *pagination.CreatedAtPageQuery `json:"next"`
	Previous *pagination.CreatedAtPageQuery `json:"previous"`
}

type PreferencesResponseDTO struct {
	EnrollmentId uuid.UUID `json:"enrollmentId"`
	Website      any       `json:"website"`
	Schedule     any       `json:"schedule"`
	Journal      any       `json:"journal"`
}
