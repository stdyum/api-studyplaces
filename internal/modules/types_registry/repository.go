package types_registry

import (
	"context"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/models"
	"github.com/stdyum/api-common/proto/impl/types_registry"
	"github.com/stdyum/api-common/umaps"
	"github.com/stdyum/api-common/uslices"
)

type iRepository interface {
	GetTypesByIds(ctx context.Context, token string, studyPlaceId uuid.UUID, ids models.TypesIds) (models.TypesModels, error)
}

type repository struct {
	client types_registry.TypesRegistryClient
}

func newRepository(client types_registry.TypesRegistryClient) iRepository {
	return &repository{client: client}
}

func (r *repository) GetTypesByIds(ctx context.Context, token string, studyPlaceId uuid.UUID, ids models.TypesIds) (models.TypesModels, error) {
	typeIds := types_registry.TypesIds{
		Token:        token,
		StudyPlaceId: studyPlaceId.String(),
		GroupsIds:    uuidsToString(ids.GroupsIds),
		RoomsIds:     uuidsToString(ids.RoomsIds),
		StudentIds:   uuidsToString(ids.StudentIds),
		SubjectsIds:  uuidsToString(ids.SubjectsIds),
		TeachersIds:  uuidsToString(ids.TeachersIds),
	}

	types, err := r.client.GetTypesByIds(ctx, &typeIds)
	if err != nil {
		return models.TypesModels{}, err
	}

	groupsIds, err := mapType(types.Groups, func(id uuid.UUID, item *types_registry.Group) models.Group {
		return models.Group{
			ID:   id,
			Name: item.Name,
		}
	})
	if err != nil {
		return models.TypesModels{}, nil
	}

	roomsIds, err := mapType(types.Rooms, func(id uuid.UUID, item *types_registry.Room) models.Room {
		return models.Room{
			ID:   id,
			Name: item.Name,
		}
	})
	if err != nil {
		return models.TypesModels{}, nil
	}

	studentIds, err := mapType(types.Students, func(id uuid.UUID, item *types_registry.Student) models.Student {
		return models.Student{
			ID:   id,
			Name: item.Name,
		}
	})
	if err != nil {
		return models.TypesModels{}, nil
	}

	subjectsIds, err := mapType(types.Subjects, func(id uuid.UUID, item *types_registry.Subject) models.Subject {
		return models.Subject{
			ID:   id,
			Name: item.Name,
		}
	})
	if err != nil {
		return models.TypesModels{}, nil
	}

	teachersIds, err := mapType(types.Teachers, func(id uuid.UUID, item *types_registry.Teacher) models.Teacher {
		return models.Teacher{
			ID:   id,
			Name: item.Name,
		}
	})
	if err != nil {
		return models.TypesModels{}, nil
	}

	return models.TypesModels{
		GroupsIds:   groupsIds,
		RoomsIds:    roomsIds,
		StudentIds:  studentIds,
		SubjectsIds: subjectsIds,
		TeachersIds: teachersIds,
	}, nil
}

func uuidsToString(uuids []uuid.UUID) []string {
	return uslices.MapFunc(uuids, func(item uuid.UUID) string {
		return item.String()
	})
}

func mapType[T any, R any](t map[string]T, converter func(id uuid.UUID, item T) R) (map[uuid.UUID]R, error) {
	return umaps.MapFuncErr(t, func(key string, value T) (uuid.UUID, R, error) {
		var nilResult R

		id, err := uuid.Parse(key)
		if err != nil {
			return uuid.Nil, nilResult, err
		}

		return id, converter(id, value), nil
	})
}
