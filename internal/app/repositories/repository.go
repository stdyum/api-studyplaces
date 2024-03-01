package repositories

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/databases/pagination"
	"github.com/stdyum/api-studyplaces/internal/app/entities"
)

type Repository interface {
	GetStudyPlacesPaginated(ctx context.Context, paginationQuery pagination.Query) ([]entities.StudyPlace, int, error)
	GetStudyPlaceById(ctx context.Context, studyPlaceId uuid.UUID) (entities.StudyPlace, error)
	CreateStudyPlace(ctx context.Context, studyPlace entities.StudyPlace) error
	DeleteStudyPlaceById(ctx context.Context, studyPlaceId uuid.UUID) error

	GetUserEnrollmentsPaginated(ctx context.Context, paginationQuery pagination.Query, userId uuid.UUID) ([]entities.Enrollment, int, error)
	GetUserEnrollmentById(ctx context.Context, id uuid.UUID) (entities.Enrollment, error)
	GetUserEnrollmentByUserIdAndStudyPlaceId(ctx context.Context, userId, studyPlaceId uuid.UUID) (entities.Enrollment, error)
	GetUserEnrollmentByIdAndUserId(ctx context.Context, userId uuid.UUID, id uuid.UUID) (entities.Enrollment, error)
	CreateEnrollment(ctx context.Context, enrollment entities.Enrollment) error
	SetEnrollmentAcceptance(ctx context.Context, enrollmentId uuid.UUID, accepted bool) error
	SetEnrollmentBlocked(ctx context.Context, enrollmentId uuid.UUID, accepted bool) error
	DeleteEnrollment(ctx context.Context, userId uuid.UUID, enrollmentId uuid.UUID) error

	CreatePreferences(ctx context.Context, preferences entities.Preferences) error
	GetPreferences(ctx context.Context, enrollmentId uuid.UUID) (entities.Preferences, error)
	UpdatePreferences(ctx context.Context, enrollmentId uuid.UUID, group string, preferences []byte) error
}

type repository struct {
	database *sql.DB
}

func New(database *sql.DB) Repository {
	return &repository{
		database: database,
	}
}
