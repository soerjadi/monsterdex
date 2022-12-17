package access_token

import (
	"context"

	"github.com/soerjadi/monsterdex/internal/model"
)

type Repository interface {
	GetUserIDByToken(ctx context.Context, token string) (model.AccessToken, error)
}

type tokenRepository struct {
	query prepareQuery
}
