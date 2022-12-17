package access_token

import (
	"context"

	"github.com/soerjadi/monsterdex/internal/log"
	"github.com/soerjadi/monsterdex/internal/model"
)

func (r *tokenRepository) GetUserIDByToken(ctx context.Context, token string) (model.AccessToken, error) {
	var (
		res model.AccessToken
		err error
	)

	err = r.query.getUserIDByToken.GetContext(ctx, &res, token)
	if err != nil {
		log.ErrorWithFields("repository.access_token.GetUserIDByToken fail get context", log.KV{
			"err": err,
		})
		return model.AccessToken{}, err
	}

	return res, nil

}
