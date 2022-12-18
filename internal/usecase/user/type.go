package user

import (
	"context"

	"github.com/soerjadi/monsterdex/internal/model"
	"github.com/soerjadi/monsterdex/internal/repository/user"
)

//go:generate mockgen -package=mocks -mock_names=Usecase=MockUserUsecase -destination=../../mocks/user_usecase_mock.go -source=type.go

type Usecase interface {
	GetUserByID(ctx context.Context, id int64) (model.User, error)
}

type userUsecase struct {
	repository user.Repository
}
