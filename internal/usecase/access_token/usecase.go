package access_token

import (
	"context"

	"github.com/soerjadi/monsterdex/internal/model"
)

func (u *tokenUsecase) GetUserIDByToken(ctx context.Context, token string) (model.AccessToken, error) {

	res, err := u.repository.GetUserIDByToken(ctx, token)
	if err != nil {
		return model.AccessToken{}, err
	}

	return res, nil
}
