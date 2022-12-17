package access_token

import "github.com/jmoiron/sqlx"

func prepareQueries(db *sqlx.DB) (prepareQuery, error) {
	var (
		err error
		q   prepareQuery
	)

	q.getUserIDByToken, err = db.Preparex(getUserIDByToken)
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

	return &tokenRepository{
		query: query,
	}, nil
}
