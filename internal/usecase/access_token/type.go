package access_token

import (
	"context"

	"github.com/soerjadi/monsterdex/internal/model"
	"github.com/soerjadi/monsterdex/internal/repository/access_token"
)

type Usecase interface {
	GetUserIDByToken(ctx context.Context, token string) (model.AccessToken, error)
}

type tokenUsecase struct {
	repository access_token.Repository
}
