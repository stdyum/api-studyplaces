package repositories

import (
	"github.com/lib/pq"
	"github.com/stdyum/api-common/databases"
	"github.com/stdyum/api-studyplaces/internal/app/entities"
)

func (r *repository) scanStudyPlace(row databases.Scan) (studyPlace entities.StudyPlace, err error) {
	err = row.Scan(
		&studyPlace.ID,
		&studyPlace.Title,
		&studyPlace.CreatedAt,
		&studyPlace.UpdatedAt,
	)
	return
}

func (r *repository) scanEnrollment(row databases.Scan) (enrollment entities.Enrollment, err error) {
	err = row.Scan(
		&enrollment.ID,
		&enrollment.UserId,
		&enrollment.StudyPlaceId,
		&enrollment.UserName,
		&enrollment.Role,
		&enrollment.TypeId,
		pq.Array(&enrollment.Permissions),
		&enrollment.Accepted,
		&enrollment.Blocked,
		&enrollment.CreatedAt,
		&enrollment.UpdatedAt,
	)
	return
}

func (r *repository) scanPreferences(row databases.Scan) (preferences entities.Preferences, err error) {
	err = row.Scan(
		&preferences.EnrollmentId,
		&preferences.Website,
		&preferences.Schedule,
		&preferences.Journal,
		&preferences.CreatedAt,
		&preferences.UpdatedAt,
	)
	return
}
