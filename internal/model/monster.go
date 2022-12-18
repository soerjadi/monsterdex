package model

import (
	"time"

	"github.com/lib/pq"
)

type Monster struct {
	ID           int64         `db:"id"`
	Name         string        `db:"name"`
	TagName      string        `db:"tag_name"`
	Description  string        `db:"description"`
	Height       float32       `db:"height"`
	Weight       int           `db:"weight"`
	Image        string        `db:"image"`
	Type         pq.Int64Array `db:"type"`
	HitPoint     int           `db:"hit_point_stat"`
	AttackPoint  int           `db:"attack_stat"`
	DefencePoint int           `db:"defence_stat"`
	SpeedPoint   int           `db:"speed_stat"`
	CreatedAt    time.Time     `db:"created_at"`
	UpdatedAt    time.Time     `db:"updated_at"`
	Captured     bool          `db:"captured"`
}

type MonsterData struct {
	ID           int64         `db:"id"`
	Name         string        `db:"name"`
	TagName      string        `db:"tag_name"`
	Description  string        `db:"description"`
	Height       float32       `db:"height"`
	Weight       int           `db:"weight"`
	Image        string        `db:"image"`
	Type         pq.Int64Array `db:"type"`
	HitPoint     int           `db:"hit_point_stat"`
	AttackPoint  int           `db:"attack_stat"`
	DefencePoint int           `db:"defence_stat"`
	SpeedPoint   int           `db:"speed_stat"`
	CreatedAt    time.Time     `db:"created_at"`
	UpdatedAt    time.Time     `db:"updated_at"`
}

type MonsterType struct {
	TypeID   int64  `json:"type_id"`
	TypeName string `json:"type_name"`
}

type MonsterListRequest struct {
	Query         *string `json:"query"`
	QueryType     *[]int  `json:"query_type"`
	Sort          *string `json:"sort"`
	SortDirection *int    `json:"sort_direction"`
}

type MonsterRequest struct {
	Name         string  `json:"name"`
	TagName      string  `json:"tag_name"`
	Description  string  `json:"description"`
	Height       float32 `json:"height"`
	Weight       int     `json:"weight"`
	Image        string  `json:"image"`
	Type         []int   `json:"type"`
	HitPoint     int     `json:"hit_point_stat"`
	AttackPoint  int     `json:"attack_point_stat"`
	DefencePoint int     `json:"defence_point_stat"`
	SpeedPoint   int     `json:"speed_point_stat"`
	ID           int64   `json:"id,omitempty"`
}

type CapturedMonster struct {
	MonsterID     int64 `db:"monster_id"`
	UserID        int64 `db:"user_id"`
	CaptureStatus bool  `db:"capture_status"`
}

type CaptureMonsterReq struct {
	UserID        int64 `json:"user_id,omitempty"`
	MonsterID     int64 `json:"monster_id,omitempty"`
	CaptureStatus bool  `json:"capture_status"`
}
