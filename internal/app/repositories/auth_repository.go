package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/models"
	"github.com/stdyum/api-common/proto/impl/auth"
)

type AuthRepository interface {
	Auth(ctx context.Context, token string) (models.User, error)
}

type authRepository struct {
	client auth.AuthClient
}

func (r *authRepository) Auth(ctx context.Context, token string) (models.User, error) {
	user, err := r.client.Auth(ctx, &auth.Token{Token: token})
	if err != nil {
		return models.User{}, err
	}

	userId, err := uuid.Parse(user.Id)
	if err != nil {
		return models.User{}, err
	}

	return models.User{
		ID:            userId,
		Login:         user.Login,
		PictureUrl:    user.PictureUrl,
		Email:         user.Email,
		VerifiedEmail: user.VerifiedEmail,
	}, nil
}

func NewAuth(client auth.AuthClient) AuthRepository {
	return &authRepository{client: client}
}
