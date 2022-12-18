package monster

import "github.com/jmoiron/sqlx"

func prepareQueries(db *sqlx.DB) (prepareQuery, error) {
	var (
		err error
		q   prepareQuery
	)

	q.getListMonster, err = db.Preparex(getListMonster)
	if err != nil {
		return q, err
	}

	q.getDetailMonster, err = db.Preparex(getDetailMonster)
	if err != nil {
		return q, err
	}

	q.captureMonster, err = db.Preparex(captureMonster)
	if err != nil {
		return q, err
	}

	q.insertMonster, err = db.Preparex(insertMonster)
	if err != nil {
		return q, err
	}

	q.updateMonster, err = db.Preparex(updateMonster)
	if err != nil {
		return q, err
	}

	q.deleteMonster, err = db.Preparex(deleteMonster)
	if err != nil {
		return q, err
	}

	return q, nil

}

func GetRepository(db *sqlx.DB) (Repository, error) {
	query, err := prepareQueries(db)
	if err != nil {
		return nil, err
	}

	return &monsterRepository{
		query: query,
		DB:    db,
	}, nil
}
