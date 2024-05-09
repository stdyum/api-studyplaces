package repositories

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/stdyum/api-common/databases/pagination"

	"github.com/stdyum/api-common/databases"
	"github.com/stdyum/api-studyplaces/internal/app/entities"
)

func (r *repository) GetStudyPlacesPaginated(ctx context.Context, paginationQuery pagination.Query) ([]entities.StudyPlace, int, error) {
	result, total, err := pagination.QueryPaginationContext(
		ctx, r.database,
		"SELECT id, title, created_at, updated_at FROM study_places",
		"SELECT count(*) FROM study_places",
		paginationQuery,
	)
	rows, err := databases.ScanArrayErr(result, r.scanStudyPlace, err)
	if err != nil {
		return nil, 0, err
	}

	return rows, total, nil
}

func (r *repository) GetStudyPlaceById(ctx context.Context, studyPlaceId uuid.UUID) (studyPlace entities.StudyPlace, err error) {
	row := r.database.QueryRowContext(ctx,
		"SELECT id, title, created_at, updated_at FROM study_places WHERE id = $1",
		studyPlaceId,
	)

	return r.scanStudyPlace(row)
}

func (r *repository) CreateStudyPlace(ctx context.Context, studyPlace entities.StudyPlace) error {
	_, err := r.database.ExecContext(ctx, "INSERT INTO study_places (id, title) VALUES ($1, $2)", studyPlace.ID, studyPlace.Title)

	return err
}

func (r *repository) DeleteStudyPlaceById(ctx context.Context, studyPlaceId uuid.UUID) error {
	result, err := r.database.ExecContext(ctx, "DELETE FROM study_places WHERE id = $1", studyPlaceId)

	return databases.AssertRowAffectedErr(result, err)
}

func (r *repository) GetUserEnrollmentsPaginated(ctx context.Context, paginationQuery pagination.Query, userId uuid.UUID) ([]entities.EnrollmentWithStudyPlace, int, error) {
	paginationQuery.SetField("enrollments.created_at")
	result, total, err := pagination.QueryPaginationContext(
		ctx, r.database,
		`
SELECT enrollments.id,
       user_id,
       study_place_id,
       study_places.title,
       user_name,
       role,
       type_id,
       permissions,
       accepted,
       blocked,
       enrollments.created_at,
       enrollments.updated_at
FROM enrollments
INNER JOIN public.study_places on study_places.id = enrollments.study_place_id
WHERE user_id = $1`,
		"SELECT count(*) FROM enrollments",
		paginationQuery,
		userId,
	)
	rows, err := databases.ScanArrayErr(result, r.scanEnrollmentWithStudyPlace, err)
	if err != nil {
		return nil, 0, err
	}

	return rows, total, nil
}

func (r *repository) GetUserEnrollmentById(ctx context.Context, id uuid.UUID) (entities.Enrollment, error) {
	row := r.database.QueryRowContext(ctx,
		"SELECT id, user_id, study_place_id, user_name, role, type_id, permissions, accepted, blocked, created_at, updated_at FROM enrollments WHERE id = $1",
		id,
	)

	return r.scanEnrollment(row)
}

func (r *repository) GetUserEnrollmentByUserIdAndStudyPlaceId(ctx context.Context, userId, studyPlaceId uuid.UUID) (entities.Enrollment, error) {
	row := r.database.QueryRowContext(ctx,
		"SELECT id, user_id, study_place_id, user_name, role, type_id, permissions, accepted, blocked, created_at, updated_at FROM enrollments WHERE user_id = $1 AND study_place_id = $2",
		userId, studyPlaceId,
	)

	return r.scanEnrollment(row)
}

func (r *repository) GetUserEnrollmentByIdAndUserId(ctx context.Context, userId, id uuid.UUID) (entities.EnrollmentWithStudyPlace, error) {
	row := r.database.QueryRowContext(ctx,
		`
SELECT enrollments.id,
       user_id,
       study_place_id,
       study_places.title,
       user_name,
       role,
       type_id,
       permissions,
       accepted,
       blocked,
       enrollments.created_at,
       enrollments.updated_at
FROM enrollments
INNER JOIN public.study_places on study_places.id = enrollments.study_place_id
WHERE enrollments.id = $1 AND user_id = $2`,
		id, userId,
	)

	return r.scanEnrollmentWithStudyPlace(row)
}

func (r *repository) CreateEnrollment(ctx context.Context, enrollment entities.Enrollment) error {
	_, err := r.database.ExecContext(ctx,
		"INSERT INTO enrollments (id, user_id, study_place_id, user_name, role, type_id, permissions, accepted, blocked) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		enrollment.ID, enrollment.UserId, enrollment.StudyPlaceId, enrollment.UserName, enrollment.Role, enrollment.TypeId, pq.Array(enrollment.Permissions), enrollment.Accepted, enrollment.Blocked,
	)

	return err
}

func (r *repository) SetEnrollmentAcceptance(ctx context.Context, enrollmentId uuid.UUID, accepted bool) error {
	result, err := r.database.ExecContext(ctx,
		"UPDATE enrollments SET accepted = $1 WHERE id = $2",
		accepted, enrollmentId,
	)

	return databases.AssertRowAffectedErr(result, err)
}

func (r *repository) SetEnrollmentBlocked(ctx context.Context, enrollmentId uuid.UUID, accepted bool) error {
	result, err := r.database.ExecContext(ctx,
		"UPDATE enrollments SET blocked = $1 WHERE id = $2",
		accepted, enrollmentId,
	)

	return databases.AssertRowAffectedErr(result, err)
}

func (r *repository) DeleteEnrollment(ctx context.Context, userId uuid.UUID, enrollmentId uuid.UUID) error {
	result, err := r.database.ExecContext(ctx,
		"DELETE FROM enrollments WHERE id=$1 and user_id=$2",
		enrollmentId, userId,
	)

	return databases.AssertRowAffectedErr(result, err)
}

func (r *repository) CreatePreferences(ctx context.Context, preferences entities.Preferences) error {
	result, err := r.database.ExecContext(ctx,
		"INSERT INTO preferences (enrollment_id, website, schedule, journal) VALUES ($1, $2, $3, $4)",
		preferences.EnrollmentId, preferences.Website, preferences.Schedule, preferences.Journal,
	)

	return databases.AssertRowAffectedErr(result, err)
}

func (r *repository) UpdatePreferences(ctx context.Context, enrollmentId uuid.UUID, group string, preferences []byte) error {
	result, err := r.database.ExecContext(ctx,
		fmt.Sprintf("UPDATE preferences SET %s=$1 WHERE enrollment_id=$2", group),
		preferences, enrollmentId,
	)

	return databases.AssertRowAffectedErr(result, err)
}

func (r *repository) GetPreferences(ctx context.Context, enrollmentId uuid.UUID) (entities.Preferences, error) {
	row := r.database.QueryRowContext(ctx,
		"SELECT enrollment_id, website, schedule, journal, created_at, updated_at FROM preferences WHERE enrollment_id = $1",
		enrollmentId,
	)

	return r.scanPreferences(row)
}
