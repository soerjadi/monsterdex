package monster

import (
	"github.com/soerjadi/monsterdex/internal/repository/monster"
)

func GetUsecase(repo monster.Repository) Usecase {
	return &monsterUsecase{
		repository: repo,
	}
}
