package user

import (
	"context"

	"github.com/soerjadi/monsterdex/internal/model"
)

//go:generate `mockgen -package=mocks -mock_names=Repository=MockUserRepository -destination=../../mocks/user_repo_mock.go -source=type.go`

type Repository interface {
	GetUserByID(ctx context.Context, id int64) (model.User, error)
}

type userRepository struct {
	query prepareQuery
}
