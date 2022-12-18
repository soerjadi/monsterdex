package user

import (
	"context"

	"github.com/soerjadi/monsterdex/internal/model"
	"github.com/soerjadi/monsterdex/internal/repository/user"
)

type Usecase interface {
	GetUserByID(ctx context.Context, id int64) (model.User, error)
}

type userUsecase struct {
	repository user.Repository
}
