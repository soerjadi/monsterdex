package model

import "time"

type User struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Role      int       `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

const (
	USER_ROLE  = 0
	ADMIN_ROLE = 1
)

var SUPPORTED_ROLE = map[int]bool{
	USER_ROLE:  true,
	ADMIN_ROLE: true,
}

func (u *User) GetRole() string {
	return ""
}
