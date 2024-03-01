package entities

import (
	"github.com/google/uuid"
	"github.com/stdyum/api-common/entities"
)

type StudyPlace struct {
	entities.Timed

	ID    uuid.UUID
	Title string
}
