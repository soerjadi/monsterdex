package monster

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

type prepareQueryMock struct {
	getListMonster   *sqlmock.ExpectedPrepare
	getDetailMonster *sqlmock.ExpectedPrepare
	captureMonster   *sqlmock.ExpectedPrepare
	insertMonster    *sqlmock.ExpectedPrepare
	updateMonster    *sqlmock.ExpectedPrepare
	deleteMonster    *sqlmock.ExpectedPrepare
}

func expectPrepareMock(mock sqlmock.Sqlmock) prepareQueryMock {
	prepareQueryMock := prepareQueryMock{}

	prepareQueryMock.getListMonster = mock.ExpectPrepare(`
	SELECT 
			id,
			name,
			tag_name,
			description,
			height,
			weight,
			image,
			type,
			hit_point_stat,
			attack_stat,
			defence_stat,
			speed_stat,
			COALESCE\(\(
				SELECT 
					capture_status 
				FROM user_monster_link uml 
				WHERE
					user_id = (.+)
			\), false\) as captured,
			created_at,
			updated_at
		FROM 
			monster m
	`)

	prepareQueryMock.getDetailMonster = mock.ExpectPrepare(`
	SELECT 
		m.id,
		m.name,
		m.tag_name,
		m.description,
		m.height,
		m.weight,
		m.image,
		m.type,
		m.hit_point_stat,
		m.attack_stat,
		m.defence_stat,
		m.speed_stat,
		COALESCE\(\(
			SELECT 
				capture_status 
			FROM user_monster_link uml 
			WHERE
				user_id = (.+)
		\), false\) as captured,
		m.created_at,
		m.updated_at
	FROM 
		monster m
	WHERE
		id = (.+)
	`)

	prepareQueryMock.captureMonster = mock.ExpectPrepare(`
	INSERT INTO user_monster_link \(
		monster_id,
		user_id,
		capture_status,
		created_at
	\) VALUES \(
		(.+),
		(.+),
		(.+),
		NOW\(\)
	\) ON CONFLICT \(monster_id, user_id\) DO 
	UPDATE SET
		capture_status=(.+)
	RETURNING id, monster_id, user_id, capture_status
	`)

	prepareQueryMock.insertMonster = mock.ExpectPrepare(`
	INSERT INTO monster\(
		name,
		tag_name,
		description,
		height,
		weight,
		image,
		type,
		hit_point_stat,
		attack_stat,
		defence_stat,
		speed_stat,
		created_at,
		updated_at
	\) VALUES \(
		(.+),
		(.+),
		(.+),
		(.+),
		(.+),
		(.+),
		(.+),
		(.+),
		(.+),
		(.+),
		(.+),
		NOW\(\),
		NOW\(\)
	\) RETURNING 
		id, 
		name,
		tag_name,
		description,
		height,
		weight,
		image,
		type,
		hit_point_stat,
		attack_stat,
		defence_stat,
		speed_stat,
		created_at,
		updated_at
	`)

	prepareQueryMock.updateMonster = mock.ExpectPrepare(`
	UPDATE monster 
	SET
		name=(.+),
		tag_name=(.+),
		description=(.+),
		height=(.+),
		weight=(.+),
		image=(.+),
		type=(.+),
		hit_point_stat=(.+),
		attack_stat=(.+),
		defence_stat=(.+),
		speed_stat=(.+),
		updated_at=NOW\(\)
	WHERE
		id=(.+)
	`)

	prepareQueryMock.deleteMonster = mock.ExpectPrepare(`
	DELETE FROM monster WHERE id=(.+)
	`)

	return prepareQueryMock
}

func TestGetRepository(t *testing.T) {
	tests := []struct {
		name     string
		initMock func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock)
		want     func(db *sqlx.DB) Repository
		wantErr  bool
	}{
		{
			name: "success",
			initMock: func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				expectPrepareMock(mock)
				expectPrepareMock(mock)
				return sqlx.NewDb(db, "postgres"), db, mock
			},
			want: func(db *sqlx.DB) Repository {
				q, _ := prepareQueries(db)

				return &monsterRepository{
					query: q,
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, dbMock, mock := tt.initMock()
			defer dbMock.Close()

			got, err := GetRepository(db)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRepository() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			want := tt.want(db)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("GetRepository() = %v, want %v", got, want)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err.Error())
			}
		})
	}

}
