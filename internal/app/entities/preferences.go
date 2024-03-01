package entities

import (
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/stdyum/api-common/entities"
)

type Preferences struct {
	entities.Timed

	EnrollmentId uuid.UUID
	Website      pgtype.JSON
	Schedule     pgtype.JSON
	Journal      pgtype.JSON
}
