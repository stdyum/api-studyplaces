package entities

import (
	"github.com/google/uuid"
	"github.com/stdyum/api-common/entities"
	"github.com/stdyum/api-common/models"
)

type Enrollment struct {
	entities.Timed

	ID           uuid.UUID
	UserId       uuid.UUID
	StudyPlaceId uuid.UUID
	UserName     string
	Role         models.Role
	TypeId       uuid.UUID
	Permissions  []string
	Accepted     bool
	Blocked      bool
}
