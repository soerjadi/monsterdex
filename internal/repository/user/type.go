package user

import (
	"context"

	"github.com/soerjadi/monsterdex/internal/model"
)

type Repository interface {
	GetUserByID(ctx context.Context, id int64) (model.User, error)
}

type userRepository struct {
	query prepareQuery
}
