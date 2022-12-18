package user

import "github.com/jmoiron/sqlx"

type prepareQuery struct {
	getUserByID *sqlx.Stmt
}

const (
	getUserByID = `
	SELECT
		id,
		name,
		email,
		role,
		created_at,
		updated_at
	FROM 
		users
	WHERE 
		id=$1
	`
)
