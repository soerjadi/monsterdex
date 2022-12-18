package user

import "github.com/soerjadi/monsterdex/internal/repository/user"

func GetUsecase(repo user.Repository) Usecase {
	return &userUsecase{
		repository: repo,
	}
}
