package controllers

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/errors"
	"github.com/stdyum/api-common/models"
	"github.com/stdyum/api-studyplaces/internal/app/entities"
)

func (c *controller) enrollmentAuth(ctx context.Context, userId, studyPlaceId uuid.UUID, permissions ...models.Permission) (entities.Enrollment, error) {
	enrollment, err := c.repository.GetUserEnrollmentByUserIdAndStudyPlaceId(ctx, userId, studyPlaceId)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return entities.Enrollment{}, errors.WrapString(errors.ErrValidation, "study place does not exist")
		}

		return entities.Enrollment{}, err
	}

	return enrollment, c.assertEnrollmentPermissions(enrollment, permissions...)
}

func (c *controller) assertEnrollmentPermissions(enrollment entities.Enrollment, permissions ...models.Permission) error {
	if !enrollment.Accepted {
		return fmt.Errorf("not accepted: %w", ErrNoPermissions)
	}

	for _, permission := range permissions {
		found := false
		for _, hasPermission := range enrollment.Permissions {
			if hasPermission == string(permission) {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("%s: %w", permission, ErrNoPermissions)
		}
	}

	return nil
}
