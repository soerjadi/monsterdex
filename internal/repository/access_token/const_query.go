package access_token

import "github.com/jmoiron/sqlx"

type prepareQuery struct {
	getUserIDByToken *sqlx.Stmt
}

const (
	getUserIDByToken = `
	SELECT
		token,
		user_id,
		created_at,
		valid_thru
	FROM
		access_token
	WHERE
		token = $1
	LIMIT 1
	`
)
