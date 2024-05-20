package controllers

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/databases/pagination"
	"github.com/stdyum/api-common/models"
	"github.com/stdyum/api-studyplaces/internal/app/dto"
	"github.com/stdyum/api-studyplaces/internal/app/repositories"
	"github.com/stdyum/api-studyplaces/internal/modules/types_registry"
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
	GetEnrollmentRequests(ctx context.Context, studyPlaceId uuid.UUID, token string, user models.User, query pagination.CreatedAtPageQuery, accepted bool) (dto.EnrollmentsResponseDTO, error)
	GetUserEnrollmentById(ctx context.Context, user models.User, id uuid.UUID) (dto.EnrollmentsResponseItemDTO, error)
	Enroll(ctx context.Context, user models.User, request dto.EnrollRequestDTO) (dto.EnrollmentsResponseItemDTO, error)
	WithdrawEnrollmentById(ctx context.Context, user models.User, enrollmentId uuid.UUID) error
	SetEnrollmentAcceptance(ctx context.Context, user models.User, enrollmentId uuid.UUID, requestDTO dto.SetEnrollmentAcceptanceRequestDTO) error
	SetEnrollmentBlocked(ctx context.Context, user models.User, enrollmentId uuid.UUID, requestDTO dto.SetEnrollmentBlockedRequestDTO) error

	GetEnrollmentPreferences(ctx context.Context, user models.User, studyPlaceId uuid.UUID) (dto.PreferencesResponseDTO, error)
	UpdateEnrollmentPreferences(ctx context.Context, user models.User, studyPlaceId uuid.UUID, preferences dto.UpdatePreferencesRequestDTO) error

	UpdateStudyPlaceEnrollment(ctx context.Context, user models.User, enrollmentId uuid.UUID, request dto.UpdateStudyPlaceEnrollmentRequestDTO) error
}

type controller struct {
	repository     repositories.Repository
	authRepository repositories.AuthRepository
	registry       types_registry.Controller
}

func New(repository repositories.Repository, authRepository repositories.AuthRepository, registry types_registry.Controller) Controller {
	return &controller{
		repository:     repository,
		authRepository: authRepository,
		registry:       registry,
	}
}
