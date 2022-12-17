package monster

import (
	"context"

	"github.com/soerjadi/monsterdex/internal/model"
	"github.com/soerjadi/monsterdex/internal/repository/monster"
)

type Usecase interface {
	GetListMonster(ctx context.Context, req model.MonsterListRequest, userID int64) ([]model.Monster, error)
	GetDetailMonster(ctx context.Context, id, userID int64) (model.Monster, error)
	CaptureMonster(ctx context.Context, req model.CaptureMonsterReq) (model.CapturedMonster, error)

	InsertMonster(ctx context.Context, req model.MonsterRequest) (model.MonsterData, error)
	UpdateMonster(ctx context.Context, req model.MonsterRequest) error
	DeleteMonster(ctx context.Context, id int64) error
}

type monsterUsecase struct {
	repository monster.Repository
}
