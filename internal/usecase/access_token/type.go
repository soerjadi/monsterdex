package access_token

import (
	"context"

	"github.com/soerjadi/monsterdex/internal/model"
	"github.com/soerjadi/monsterdex/internal/repository/access_token"
)

//go:generate mockgen -package=mocks -mock_names=Usecase=MockAccessTokenUsecase -destination=../../mocks/access_token_usecase_mock.go -source=type.go

type Usecase interface {
	GetUserIDByToken(ctx context.Context, token string) (model.AccessToken, error)
}

type tokenUsecase struct {
	repository access_token.Repository
}
