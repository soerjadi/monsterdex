package access_token

import (
	"context"

	"github.com/soerjadi/monsterdex/internal/model"
)

//go:generate `mockgen -package=mocks -mock_names=Repository=MockAccessTokenRepository -destination=../../mocks/access_token_repo_mock.go -source=type.go`

type Repository interface {
	GetUserIDByToken(ctx context.Context, token string) (model.AccessToken, error)
}

type tokenRepository struct {
	query prepareQuery
}
