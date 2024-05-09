package controllers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/databases/pagination"
	timed "github.com/stdyum/api-common/entities"
	"github.com/stdyum/api-common/models"
	"github.com/stdyum/api-studyplaces/internal/app/dto"
	"github.com/stdyum/api-studyplaces/internal/app/entities"
)

func (c *controller) Auth(ctx context.Context, requestDTO dto.AuthRequestDTO) (dto.AuthResponseDTO, error) {
	user, err := c.authRepository.Auth(ctx, requestDTO.Token)
	if err != nil {
		return dto.AuthResponseDTO{}, err
	}

	enrollmentAuth, err := c.enrollmentAuth(ctx, user.ID, requestDTO.StudyPlaceId)
	if err != nil {
		return dto.AuthResponseDTO{}, err
	}

	return dto.AuthResponseDTO{
		Id:            user.ID,
		Login:         user.Login,
		PictureUrl:    user.PictureUrl,
		Email:         user.Email,
		VerifiedEmail: user.VerifiedEmail,
		Enrollment: dto.EnrollmentsResponseItemDTO{
			ID:           enrollmentAuth.ID,
			UserId:       enrollmentAuth.UserId,
			StudyPlaceId: enrollmentAuth.StudyPlaceId,
			UserName:     enrollmentAuth.UserName,
			Role:         enrollmentAuth.Role,
			TypeId:       enrollmentAuth.TypeId,
			Permissions:  enrollmentAuth.Permissions,
			Accepted:     enrollmentAuth.Accepted,
		},
	}, nil
}

func (c *controller) GetStudyPlaces(ctx context.Context, query pagination.CreatedAtPageQuery) (dto.StudyPlacesResponseDTO, error) {
	rows, total, err := c.repository.GetStudyPlacesPaginated(ctx, &query)
	if err != nil {
		return dto.StudyPlacesResponseDTO{}, err
	}

	timedRows := make([]timed.Timed, len(rows))
	studyPlaces := make([]dto.StudyPlacesResponseItemDTO, len(rows))
	for i, row := range rows {
		timedRows[i] = row.Timed
		studyPlaces[i] = dto.StudyPlacesResponseItemDTO{
			ID:    row.ID,
			Title: row.Title,
		}
	}

	return dto.StudyPlacesResponseDTO{
		Items:    studyPlaces,
		PerPage:  query.QPerPage,
		Total:    total,
		Page:     query.QPage,
		Next:     pagination.GetNextCreatedAtQuery(timedRows, total, &query),
		Previous: pagination.GetPreviousCreatedAtQuery(timedRows, &query),
	}, nil
}

func (c *controller) GetStudyPlaceById(ctx context.Context, id uuid.UUID) (dto.StudyPlacesResponseItemDTO, error) {
	studyPlace, err := c.repository.GetStudyPlaceById(ctx, id)
	if err != nil {
		return dto.StudyPlacesResponseItemDTO{}, err
	}

	return dto.StudyPlacesResponseItemDTO{
		ID:    studyPlace.ID,
		Title: studyPlace.Title,
	}, nil
}

func (c *controller) RegisterStudyPlace(ctx context.Context, user models.User, request dto.CreateStudyPlaceRequestDTO) (dto.CreateStudyPlaceResponseDTO, error) {
	studyPlace := entities.StudyPlace{
		ID:    uuid.New(),
		Title: request.Title,
	}

	if err := c.repository.CreateStudyPlace(ctx, studyPlace); err != nil {
		return dto.CreateStudyPlaceResponseDTO{}, err
	}

	enrollment := entities.Enrollment{
		ID:           uuid.New(),
		UserId:       user.ID,
		StudyPlaceId: studyPlace.ID,
		UserName:     request.AdminUserName,
		Role:         models.RoleAdmin,
		TypeId:       uuid.Nil,
		Permissions:  []string{string(models.PermissionAdmin)},
		Accepted:     true,
	}

	if err := c.repository.CreateEnrollment(ctx, enrollment); err != nil {
		return dto.CreateStudyPlaceResponseDTO{}, err
	}

	preferences := entities.Preferences{
		EnrollmentId: enrollment.ID,
	}
	if err := c.repository.CreatePreferences(ctx, preferences); err != nil {
		return dto.CreateStudyPlaceResponseDTO{}, err
	}

	return dto.CreateStudyPlaceResponseDTO{
		StudyPlace: dto.StudyPlacesResponseItemDTO{
			ID:    studyPlace.ID,
			Title: studyPlace.Title,
		},
		Enrollment: dto.EnrollmentsResponseItemDTO{
			ID:           enrollment.ID,
			UserId:       enrollment.UserId,
			StudyPlaceId: enrollment.StudyPlaceId,
			UserName:     enrollment.UserName,
			Role:         enrollment.Role,
			TypeId:       enrollment.TypeId,
			Permissions:  enrollment.Permissions,
			Accepted:     enrollment.Accepted,
		},
	}, nil
}

func (c *controller) CloseStudyPlaceById(ctx context.Context, user models.User, id uuid.UUID) error {
	if _, err := c.enrollmentAuth(ctx, user.ID, id, models.PermissionAdmin); err != nil {
		return err
	}

	return c.repository.DeleteStudyPlaceById(ctx, id)
}

func (c *controller) GetUserEnrollments(ctx context.Context, user models.User, query pagination.CreatedAtPageQuery) (dto.EnrollmentsResponseDTO, error) {
	rows, total, err := c.repository.GetUserEnrollmentsPaginated(ctx, &query, user.ID)
	if err != nil {
		return dto.EnrollmentsResponseDTO{}, err
	}

	timedRows := make([]timed.Timed, len(rows))
	enrollments := make([]dto.EnrollmentsResponseItemDTO, len(rows))
	for i, row := range rows {
		timedRows[i] = row.Timed
		enrollments[i] = dto.EnrollmentsResponseItemDTO{
			ID:              row.ID,
			UserId:          row.UserId,
			StudyPlaceId:    row.StudyPlaceId,
			StudyPlaceTitle: row.StudyPlaceTitle,
			UserName:        row.UserName,
			Role:            row.Role,
			TypeId:          row.TypeId,
			Permissions:     row.Permissions,
			Accepted:        row.Accepted,
		}
	}

	return dto.EnrollmentsResponseDTO{
		Items:    enrollments,
		PerPage:  query.QPerPage,
		Total:    total,
		Page:     query.QPage,
		Next:     pagination.GetNextCreatedAtQuery(timedRows, total, &query),
		Previous: pagination.GetPreviousCreatedAtQuery(timedRows, &query),
	}, nil
}

func (c *controller) GetUserEnrollmentById(ctx context.Context, user models.User, enrollmentId uuid.UUID) (dto.EnrollmentsResponseItemDTO, error) {
	enrollment, err := c.repository.GetUserEnrollmentByIdAndUserId(ctx, user.ID, enrollmentId)
	if err != nil {
		return dto.EnrollmentsResponseItemDTO{}, err
	}

	return dto.EnrollmentsResponseItemDTO{
		ID:              enrollment.ID,
		UserId:          enrollment.UserId,
		StudyPlaceId:    enrollment.StudyPlaceId,
		StudyPlaceTitle: enrollment.StudyPlaceTitle,
		UserName:        enrollment.UserName,
		Role:            enrollment.Role,
		TypeId:          enrollment.TypeId,
		Permissions:     enrollment.Permissions,
		Accepted:        enrollment.Accepted,
	}, nil
}

func (c *controller) Enroll(ctx context.Context, user models.User, request dto.EnrollRequestDTO) (dto.EnrollmentsResponseItemDTO, error) {
	enrollment := entities.Enrollment{
		ID:           uuid.New(),
		UserId:       user.ID,
		StudyPlaceId: request.StudyPlaceId,
		UserName:     request.UserName,
		Role:         request.Role,
		TypeId:       request.TypeId,
		Permissions:  nil,
		Accepted:     false,
	}

	if err := c.repository.CreateEnrollment(ctx, enrollment); err != nil {
		return dto.EnrollmentsResponseItemDTO{}, err
	}

	preferences := entities.Preferences{
		EnrollmentId: enrollment.ID,
	}
	if err := c.repository.CreatePreferences(ctx, preferences); err != nil {
		return dto.EnrollmentsResponseItemDTO{}, err
	}

	return dto.EnrollmentsResponseItemDTO{
		ID:           enrollment.ID,
		UserId:       enrollment.UserId,
		StudyPlaceId: enrollment.StudyPlaceId,
		UserName:     enrollment.UserName,
		Role:         enrollment.Role,
		TypeId:       enrollment.TypeId,
		Permissions:  enrollment.Permissions,
		Accepted:     enrollment.Accepted,
	}, nil
}

func (c *controller) WithdrawEnrollmentById(ctx context.Context, user models.User, enrollmentId uuid.UUID) error {
	return c.repository.DeleteEnrollment(ctx, user.ID, enrollmentId)
}

func (c *controller) SetEnrollmentAcceptance(ctx context.Context, user models.User, enrollmentId uuid.UUID, requestDTO dto.SetEnrollmentAcceptanceRequestDTO) error {
	enrollment, err := c.repository.GetUserEnrollmentById(ctx, enrollmentId)
	if err != nil {
		return err
	}

	if _, err = c.enrollmentAuth(ctx, user.ID, enrollment.StudyPlaceId, models.PermissionEnrollments); err != nil {
		return err
	}

	return c.repository.SetEnrollmentAcceptance(ctx, enrollmentId, requestDTO.Accepted)
}

func (c *controller) SetEnrollmentBlocked(ctx context.Context, user models.User, enrollmentId uuid.UUID, requestDTO dto.SetEnrollmentBlockedRequestDTO) error {
	enrollment, err := c.repository.GetUserEnrollmentById(ctx, enrollmentId)
	if err != nil {
		return err
	}

	if _, err = c.enrollmentAuth(ctx, user.ID, enrollment.StudyPlaceId, models.PermissionEnrollments); err != nil {
		return err
	}

	return c.repository.SetEnrollmentAcceptance(ctx, enrollmentId, requestDTO.Blocked)
}

func (c *controller) GetEnrollmentPreferences(ctx context.Context, user models.User, enrollmentId uuid.UUID) (dto.PreferencesResponseDTO, error) {
	enrollment, err := c.repository.GetUserEnrollmentByIdAndUserId(ctx, user.ID, enrollmentId)
	if err != nil {
		return dto.PreferencesResponseDTO{}, err
	}

	preferences, err := c.repository.GetPreferences(ctx, enrollment.ID)
	if err != nil {
		return dto.PreferencesResponseDTO{}, err
	}

	return dto.PreferencesResponseDTO{
		EnrollmentId: preferences.EnrollmentId,
		Website:      preferences.Website,
		Schedule:     preferences.Schedule,
		Journal:      preferences.Journal,
	}, nil
}

func (c *controller) UpdateEnrollmentPreferences(ctx context.Context, user models.User, enrollmentId uuid.UUID, requestDTO dto.UpdatePreferencesRequestDTO) error {
	enrollment, err := c.repository.GetUserEnrollmentByIdAndUserId(ctx, user.ID, enrollmentId)
	if err != nil {
		return err
	}

	preferencesBytes, err := json.Marshal(requestDTO.Preferences)
	if err != nil {
		return err
	}

	if !requestDTO.Group.IsValid() {
		return fmt.Errorf("no such group: %s: %w", requestDTO.Group, ErrValidation)
	}

	return c.repository.UpdatePreferences(ctx, enrollment.ID, string(requestDTO.Group), preferencesBytes)
}
