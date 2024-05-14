package types_registry

import (
	"github.com/stdyum/api-common/proto/impl/types_registry"
)

func New(client types_registry.TypesRegistryClient) Controller {
	repo := newRepository(client)
	ctrl := newController(repo)

	return ctrl
}
