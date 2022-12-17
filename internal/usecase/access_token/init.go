package access_token

import "github.com/soerjadi/monsterdex/internal/repository/access_token"

func GetUsecase(repo access_token.Repository) Usecase {
	return &tokenUsecase{
		repository: repo,
	}
}
