package controllers

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/databases/pagination"
	"github.com/stdyum/api-common/models"
	"github.com/stdyum/api-studyplaces/internal/app/dto"
	"github.com/stdyum/api-studyplaces/internal/app/repositories"
)

var (
	ErrNoPermissions = errors.New("no permission")
	ErrValidation    = errors.New("validation")
)

type Controller interface {
	Auth(ctx context.Context, requestDTO dto.AuthRequestDTO) (dto.AuthResponseDTO, error)

	GetStudyPlaces(ctx context.Context, query pagination.CreatedAtPageQuery) (dto.StudyPlacesResponseDTO, error)
	GetStudyPlaceById(ctx context.Context, id uuid.UUID) (dto.StudyPlacesResponseItemDTO, error)
	RegisterStudyPlace(ctx context.Context, user models.User, studyPlace dto.CreateStudyPlaceRequestDTO) (dto.CreateStudyPlaceResponseDTO, error)
	CloseStudyPlaceById(ctx context.Context, user models.User, id uuid.UUID) error

	GetUserEnrollments(ctx context.Context, user models.User, query pagination.CreatedAtPageQuery) (dto.EnrollmentsResponseDTO, error)
	Enroll(ctx context.Context, user models.User, request dto.EnrollRequestDTO) (dto.EnrollmentsResponseItemDTO, error)
	WithdrawEnrollmentById(ctx context.Context, user models.User, enrollmentId uuid.UUID) error
	SetEnrollmentAcceptance(ctx context.Context, user models.User, enrollmentId uuid.UUID, requestDTO dto.SetEnrollmentAcceptanceRequestDTO) error
	SetEnrollmentBlocked(ctx context.Context, user models.User, enrollmentId uuid.UUID, requestDTO dto.SetEnrollmentBlockedRequestDTO) error

	GetEnrollmentPreferences(ctx context.Context, user models.User, enrollmentId uuid.UUID) (dto.PreferencesResponseDTO, error)
	UpdateEnrollmentPreferences(ctx context.Context, user models.User, enrollmentId uuid.UUID, preferences dto.UpdatePreferencesRequestDTO) error
}

type controller struct {
	repository     repositories.Repository
	authRepository repositories.AuthRepository
}

func New(repository repositories.Repository, authRepository repositories.AuthRepository) Controller {
	return &controller{
		repository:     repository,
		authRepository: authRepository,
	}
}
