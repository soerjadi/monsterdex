package user

import (
	"context"

	"github.com/soerjadi/monsterdex/internal/log"
	"github.com/soerjadi/monsterdex/internal/model"
)

func (r *userRepository) GetUserByID(ctx context.Context, id int64) (model.User, error) {
	var (
		res model.User
		err error
	)

	err = r.query.getUserByID.GetContext(ctx, &res, id)
	if err != nil {
		log.ErrorWithFields("repository.User.GetUserByID fail get context", log.KV{
			"err": err,
		})

		return model.User{}, err
	}

	return res, nil
}
