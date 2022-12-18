package user

import (
	"context"

	"github.com/soerjadi/monsterdex/internal/model"
)

func (u *userUsecase) GetUserByID(ctx context.Context, id int64) (model.User, error) {

	res, err := u.repository.GetUserByID(ctx, id)
	if err != nil {
		return model.User{}, err
	}

	return res, nil
}
