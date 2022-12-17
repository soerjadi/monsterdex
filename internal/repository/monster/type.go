package monster

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/soerjadi/monsterdex/internal/model"
)

//go:generate `mockgen -package=mocks -mock_names=Repository=MockMonsterRepository -destination=../../mocks/monster_repo_mock.go -source=type.go`

type Repository interface {
	GetListMonster(ctx context.Context, req model.MonsterListRequest, userID int64) ([]model.Monster, error)
	GetDetailMonster(ctx context.Context, id int64, userID int64) (model.Monster, error)
	CaptureMonster(ctx context.Context, req model.CaptureMonsterReq) (model.CapturedMonster, error)

	InsertMonster(ctx context.Context, req model.MonsterRequest) (model.MonsterData, error)
	UpdateMonster(ctx context.Context, req model.MonsterRequest) error
	DeleteMonster(ctx context.Context, id int64) error
}

type monsterRepository struct {
	query prepareQuery
	DB    *sqlx.DB
}

type buildQueryParam struct {
	query          string
	monsterListReq model.MonsterListRequest
	userID         int64
}

type buildQueryResult struct {
	query string
	args  []interface{}
}
