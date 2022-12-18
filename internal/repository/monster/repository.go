package monster

import (
	"context"

	"github.com/lib/pq"
	"github.com/soerjadi/monsterdex/internal/log"
	"github.com/soerjadi/monsterdex/internal/model"
)

func (r *monsterRepository) GetListMonster(ctx context.Context, req model.MonsterListRequest, userID int64) ([]model.Monster, error) {
	var dbRes []model.Monster

	param := buildQueryParam{
		query:          getListMonster,
		monsterListReq: req,
		userID:         userID,
	}

	q, err := buildMonsterListQuery(ctx, param)
	if err != nil {
		log.ErrorWithFields("repository.monster.GetListMonster fail buildMonsterListQuery", log.KV{
			"err":   err,
			"args":  q.args,
			"query": q.query,
		})
		return dbRes, err
	}

	err = r.DB.SelectContext(ctx, &dbRes, q.query, q.args...)
	if err != nil {
		log.ErrorWithFields("repository.monster.GetListMonster fail select context", log.KV{
			"err":   err,
			"args":  q.args,
			"query": q.query,
		})
		return []model.Monster{}, err
	}

	return dbRes, nil
}

func (r *monsterRepository) GetDetailMonster(ctx context.Context, id int64, userID int64) (model.Monster, error) {
	var res model.Monster

	err := r.query.getDetailMonster.GetContext(ctx, &res, userID, id)
	if err != nil {
		log.ErrorWithFields("repository.monster.GetDetailMonster fail get context", log.KV{
			"err":  err,
			"args": id,
		})

		return model.Monster{}, err
	}

	return res, nil
}

func (r *monsterRepository) CaptureMonster(ctx context.Context, req model.CaptureMonsterReq) (model.CapturedMonster, error) {
	var (
		err error
		res model.CapturedMonster
	)

	if err = r.query.captureMonster.GetContext(ctx, &res, req.MonsterID, req.UserID, req.CaptureStatus); err != nil {
		log.ErrorWithFields("repository.monster.CaptureMonster fail upsert capture monster", log.KV{
			"err":     err,
			"request": req,
		})
		return model.CapturedMonster{}, err
	}

	return res, nil
}

func (r *monsterRepository) InsertMonster(ctx context.Context, req model.MonsterRequest) (model.MonsterData, error) {
	var (
		err error
		res model.MonsterData
	)

	if err = r.query.insertMonster.GetContext(
		ctx,
		&res,
		req.Name,
		req.TagName,
		req.Description,
		req.Height,
		req.Weight,
		req.Image,
		pq.Array(req.Type),
		req.HitPoint,
		req.AttackPoint,
		req.DefencePoint,
		req.SpeedPoint,
	); err != nil {
		log.ErrorWithFields("repository.monster.InsertMonster fail insert monster", log.KV{
			"err":     err,
			"request": req,
		})
		return model.MonsterData{}, err
	}

	return res, nil
}

func (r *monsterRepository) UpdateMonster(ctx context.Context, req model.MonsterRequest) error {
	var err error

	if _, err = r.query.updateMonster.ExecContext(
		ctx,
		req.Name,
		req.TagName,
		req.Description,
		req.Height,
		req.Weight,
		req.Image,
		pq.Array(req.Type),
		req.HitPoint,
		req.AttackPoint,
		req.DefencePoint,
		req.SpeedPoint,
		req.ID,
	); err != nil {
		log.ErrorWithFields("repository.monter.UpdateMonster fail update monster", log.KV{
			"err":     err,
			"request": req,
		})

		return err
	}

	return nil
}

func (r *monsterRepository) DeleteMonster(ctx context.Context, id int64) error {
	var err error

	if _, err = r.query.deleteMonster.ExecContext(ctx, id); err != nil {
		log.ErrorWithFields("repository.monster.DeleteMonster fail delete monster", log.KV{
			"err": err,
			"id":  id,
		})
		return err
	}

	return nil
}
