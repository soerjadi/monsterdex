package access_token

import (
	"context"
	"errors"
	"time"

	"github.com/soerjadi/monsterdex/internal/model"
)

func (u *tokenUsecase) GetUserIDByToken(ctx context.Context, token string) (model.AccessToken, error) {

	res, err := u.repository.GetUserIDByToken(ctx, token)
	if err != nil {
		return model.AccessToken{}, err
	}

	if res.CreatedAt.After(time.Now()) {
		return model.AccessToken{}, errors.New("unauthorized")
	}

	return res, nil
}
