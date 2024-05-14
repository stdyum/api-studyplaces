package types_registry

import (
	"context"

	"github.com/stdyum/api-common/models"
)

type Controller interface {
	GetTypesByIds(ctx context.Context, enrollment models.Enrollment, ids models.TypesIds) (models.TypesModels, error)
}

type controller struct {
	repository iRepository
}

func newController(repository iRepository) Controller {
	return &controller{repository: repository}
}

func (c *controller) GetTypesByIds(ctx context.Context, enrollment models.Enrollment, ids models.TypesIds) (models.TypesModels, error) {
	return c.repository.GetTypesByIds(ctx, enrollment.Token, enrollment.StudyPlaceId, ids)
}
