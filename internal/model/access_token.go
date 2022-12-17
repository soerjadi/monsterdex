package model

import "time"

type AccessToken struct {
	Token     string    `db:"token"`
	UserID    int64     `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	ValidThru time.Time `db:"valid_thru"`
}
