package monster

import (
	"context"

	"github.com/soerjadi/monsterdex/internal/model"
)

func (u *monsterUsecase) GetListMonster(ctx context.Context, req model.MonsterListRequest, userID int64) ([]model.Monster, error) {

	res, err := u.repository.GetListMonster(ctx, req, userID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *monsterUsecase) GetDetailMonster(ctx context.Context, id, userID int64) (model.Monster, error) {

	res, err := u.repository.GetDetailMonster(ctx, id, userID)
	if err != nil {
		return model.Monster{}, err
	}

	return res, nil

}

func (u *monsterUsecase) CaptureMonster(ctx context.Context, req model.CaptureMonsterReq) (model.CapturedMonster, error) {

	res, err := u.repository.CaptureMonster(ctx, req)
	if err != nil {
		return model.CapturedMonster{}, err
	}

	return res, nil
}

func (u *monsterUsecase) InsertMonster(ctx context.Context, req model.MonsterRequest) (model.MonsterData, error) {

	res, err := u.repository.InsertMonster(ctx, req)
	if err != nil {
		return model.MonsterData{}, err
	}

	return res, nil

}

func (u *monsterUsecase) UpdateMonster(ctx context.Context, req model.MonsterRequest) error {

	err := u.repository.UpdateMonster(ctx, req)
	if err != nil {
		return err
	}

	return nil

}

func (u *monsterUsecase) DeleteMonster(ctx context.Context, id int64) error {

	err := u.repository.DeleteMonster(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
