package monster

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/soerjadi/monsterdex/internal/model"
)

func TestGetListMonster(t *testing.T) {
	type args struct {
		ctx context.Context
		req model.MonsterListRequest
		id  int64
	}
	tests := []struct {
		name     string
		args     args
		initMock func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock)
		want     []model.Monster
		wantErr  bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.TODO(),
				req: model.MonsterListRequest{},
				id:  1,
			},
			initMock: func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				p := expectPrepareMock(mock)
				p.getListMonster.ExpectQuery().WillReturnRows(sqlmock.NewRows([]string{
					"id",
					"name",
					"tag_name",
					"description",
					"height",
					"weight",
					"image",
					"type",
					"hit_point_stat",
					"attack_stat",
					"defence_stat",
					"speed_stat",
					"captured",
					"created_at",
					"updated_at",
				}).
					AddRow(1, "monster1", "monster1 tag", "description", 21.5, 216, "image", pq.Array([]int64{1, 2}), 250, 300, 400, 320, false, time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC), time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC)).
					AddRow(2, "monster2", "monster2 tag", "description", 21.5, 216, "image", pq.Array([]int64{1, 2}), 250, 300, 400, 320, false, time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC), time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC)))

				return sqlx.NewDb(db, "postgres"), db, mock
			},
			want: []model.Monster{
				{
					ID:           1,
					Name:         "monster1",
					TagName:      "monster1 tag",
					Description:  "description",
					Height:       21.5,
					Weight:       216,
					Image:        "image",
					Type:         []int64{1, 2},
					HitPoint:     250,
					AttackPoint:  300,
					DefencePoint: 400,
					SpeedPoint:   320,
					Captured:     false,
					CreatedAt:    time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC),
					UpdatedAt:    time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC),
				},
				{
					ID:           2,
					Name:         "monster2",
					TagName:      "monster2 tag",
					Description:  "description",
					Height:       21.5,
					Weight:       216,
					Image:        "image",
					Type:         []int64{1, 2},
					HitPoint:     250,
					AttackPoint:  300,
					DefencePoint: 400,
					SpeedPoint:   320,
					Captured:     false,
					CreatedAt:    time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC),
					UpdatedAt:    time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC),
				},
			},
			wantErr: false,
		},
		{
			name: "fail",
			args: args{
				ctx: context.TODO(),
				req: model.MonsterListRequest{},
				id:  1,
			},
			initMock: func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				p := expectPrepareMock(mock)
				p.getListMonster.ExpectQuery().WillReturnError(errors.New("mock error"))

				return sqlx.NewDb(db, "postgres"), db, mock
			},
			want:    []model.Monster{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, dbMock, mock := tt.initMock()
			defer dbMock.Close()
			q, _ := prepareQueries(db)

			re := monsterRepository{
				query: q,
				DB:    db,
			}
			got, err := re.GetListMonster(tt.args.ctx, tt.args.req, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetListMonster() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetListMonster() = %v, want %v", got, tt.want)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err.Error())
			}
		})
	}
}

func TestGetDetailMonster(t *testing.T) {
	type args struct {
		ctx    context.Context
		id     int64
		userID int64
	}
	tests := []struct {
		name     string
		args     args
		initMock func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock)
		want     model.Monster
		wantErr  bool
	}{
		{
			name: "success",
			args: args{
				ctx:    context.TODO(),
				id:     1,
				userID: 1,
			},
			initMock: func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				p := expectPrepareMock(mock)
				p.getDetailMonster.ExpectQuery().WillReturnRows(sqlmock.NewRows([]string{
					"id",
					"name",
					"tag_name",
					"description",
					"height",
					"weight",
					"image",
					"type",
					"hit_point_stat",
					"attack_stat",
					"defence_stat",
					"speed_stat",
					"captured",
					"created_at",
					"updated_at",
				}).
					AddRow(1, "monster1", "monster1 tag", "description", 21.5, 216, "image", pq.Array([]int{1, 2}), 250, 300, 400, 320, false, time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC), time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC)))

				return sqlx.NewDb(db, "postgres"), db, mock
			},
			want: model.Monster{
				ID:           1,
				Name:         "monster1",
				TagName:      "monster1 tag",
				Description:  "description",
				Height:       21.5,
				Weight:       216,
				Image:        "image",
				Type:         []int64{1, 2},
				HitPoint:     250,
				AttackPoint:  300,
				DefencePoint: 400,
				SpeedPoint:   320,
				Captured:     false,
				CreatedAt:    time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC),
				UpdatedAt:    time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name: "fail",
			args: args{
				ctx:    context.TODO(),
				id:     1,
				userID: 1,
			},
			initMock: func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				p := expectPrepareMock(mock)
				p.getDetailMonster.ExpectQuery().WillReturnError(errors.New("mock error"))

				return sqlx.NewDb(db, "postgres"), db, mock
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, dbMock, mock := tt.initMock()
			defer dbMock.Close()
			q, _ := prepareQueries(db)

			re := monsterRepository{
				query: q,
				DB:    db,
			}
			got, err := re.GetDetailMonster(tt.args.ctx, tt.args.id, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetDetailMonster() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetDetailMonster() = %v, want %v", got, tt.want)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err.Error())
			}
		})
	}
}

func TestCaptureMonster(t *testing.T) {
	type args struct {
		ctx context.Context
		req model.CaptureMonsterReq
	}
	tests := []struct {
		name     string
		args     args
		initMock func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock)
		want     model.CapturedMonster
		wantErr  bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.TODO(),
				req: model.CaptureMonsterReq{
					UserID:        1,
					MonsterID:     1,
					CaptureStatus: true,
				},
			},
			initMock: func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				p := expectPrepareMock(mock)
				p.captureMonster.ExpectQuery().WillReturnRows(sqlmock.NewRows([]string{
					"id",
					"user_id",
					"monster_id",
					"capture_status",
				}).
					AddRow(1, 1, 1, true))

				return sqlx.NewDb(db, "postgres"), db, mock
			},
			want: model.CapturedMonster{
				UserID:        1,
				MonsterID:     1,
				CaptureStatus: true,
			},
			wantErr: false,
		},
		{
			name: "fail",
			args: args{
				ctx: context.TODO(),
				req: model.CaptureMonsterReq{
					UserID:        1,
					MonsterID:     1,
					CaptureStatus: true,
				},
			},
			initMock: func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				p := expectPrepareMock(mock)
				p.captureMonster.ExpectQuery().WillReturnError(errors.New("mock error"))

				return sqlx.NewDb(db, "postgres"), db, mock
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, dbMock, mock := tt.initMock()
			defer dbMock.Close()
			q, _ := prepareQueries(db)

			re := monsterRepository{
				query: q,
			}
			got, err := re.CaptureMonster(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.CaptureMonster() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.CaptureMonster() = %v, want %v", got, tt.want)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err.Error())
			}
		})
	}
}

func TestInsertMonster(t *testing.T) {
	type args struct {
		ctx context.Context
		req model.MonsterRequest
	}
	tests := []struct {
		name     string
		args     args
		initMock func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock)
		want     model.MonsterData
		wantErr  bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.TODO(),
				req: model.MonsterRequest{
					Name:         "monster1",
					TagName:      "monster1 tag",
					Description:  "description",
					Height:       21.5,
					Weight:       216,
					Image:        "image",
					Type:         []int{1, 2},
					HitPoint:     250,
					AttackPoint:  300,
					DefencePoint: 400,
					SpeedPoint:   320,
				},
			},
			initMock: func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				p := expectPrepareMock(mock)
				p.insertMonster.ExpectQuery().WillReturnRows(sqlmock.NewRows([]string{
					"id",
					"name",
					"tag_name",
					"description",
					"height",
					"weight",
					"image",
					"type",
					"hit_point_stat",
					"attack_stat",
					"defence_stat",
					"speed_stat",
					"created_at",
					"updated_at",
				}).
					AddRow(1, "monster1", "monster1 tag", "description", 21.5, 216, "image", pq.Array([]int{1, 2}), 250, 300, 400, 320, time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC), time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC)))

				return sqlx.NewDb(db, "postgres"), db, mock
			},
			want: model.MonsterData{
				ID:           1,
				Name:         "monster1",
				TagName:      "monster1 tag",
				Description:  "description",
				Height:       21.5,
				Weight:       216,
				Image:        "image",
				Type:         []int64{1, 2},
				HitPoint:     250,
				AttackPoint:  300,
				DefencePoint: 400,
				SpeedPoint:   320,
				CreatedAt:    time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC),
				UpdatedAt:    time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name: "fail",
			args: args{
				ctx: context.TODO(),
				req: model.MonsterRequest{
					Name:         "monster1",
					TagName:      "monster1 tag",
					Description:  "description",
					Height:       21.5,
					Weight:       216,
					Image:        "image",
					Type:         []int{1, 2},
					HitPoint:     250,
					AttackPoint:  300,
					DefencePoint: 400,
					SpeedPoint:   320,
				},
			},
			initMock: func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				p := expectPrepareMock(mock)
				p.insertMonster.ExpectQuery().WillReturnError(errors.New("mock error"))

				return sqlx.NewDb(db, "postgres"), db, mock
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, dbMock, mock := tt.initMock()
			defer dbMock.Close()
			q, _ := prepareQueries(db)

			re := monsterRepository{
				query: q,
			}
			got, err := re.InsertMonster(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.InsertMonster() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.InsertMonster() = %v, want %v", got, tt.want)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err.Error())
			}
		})
	}
}

func TestUpdateMonster(t *testing.T) {
	type args struct {
		ctx context.Context
		req model.MonsterRequest
	}
	tests := []struct {
		name     string
		args     args
		initMock func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock)
		wantErr  bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.TODO(),
				req: model.MonsterRequest{
					Name:         "monster1",
					TagName:      "monster1 tag",
					Description:  "description",
					Height:       21.5,
					Weight:       216,
					Image:        "image",
					Type:         []int{1, 2},
					HitPoint:     250,
					AttackPoint:  300,
					DefencePoint: 400,
					SpeedPoint:   320,
					ID:           1,
				},
			},
			initMock: func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				p := expectPrepareMock(mock)
				p.updateMonster.ExpectExec().WithArgs(
					"monster1",
					"monster1 tag",
					"description",
					21.5,
					216,
					"image",
					pq.Array([]int{1, 2}),
					250,
					300,
					400,
					320,
					1,
				).WillReturnResult(sqlmock.NewResult(1, 1))

				return sqlx.NewDb(db, "postgres"), db, mock
			},
			wantErr: false,
		},
		{
			name: "fail",
			args: args{
				ctx: context.TODO(),
				req: model.MonsterRequest{
					Name:         "monster1",
					TagName:      "monster1 tag",
					Description:  "description",
					Height:       21.5,
					Weight:       216,
					Image:        "image",
					Type:         []int{1, 2},
					HitPoint:     250,
					AttackPoint:  300,
					DefencePoint: 400,
					SpeedPoint:   320,
					ID:           1,
				},
			},
			initMock: func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				p := expectPrepareMock(mock)
				p.updateMonster.ExpectExec().WithArgs(
					"monster1",
					"monster1 tag",
					"description",
					21.5,
					216,
					"image",
					pq.Array([]int{1, 2}),
					250,
					300,
					400,
					320,
					1,
				).WillReturnError(errors.New("mock error"))

				return sqlx.NewDb(db, "postgres"), db, mock
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, dbMock, mock := tt.initMock()
			defer dbMock.Close()
			q, _ := prepareQueries(db)

			re := monsterRepository{
				query: q,
			}
			err := re.UpdateMonster(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.UpdateMonster() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err.Error())
			}
		})
	}
}

func TestDeleteMonster(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name     string
		args     args
		initMock func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock)
		wantErr  bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.TODO(),
				id:  1,
			},
			initMock: func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				p := expectPrepareMock(mock)
				p.deleteMonster.ExpectExec().WithArgs(
					1,
				).WillReturnResult(sqlmock.NewResult(1, 1))

				return sqlx.NewDb(db, "postgres"), db, mock
			},
			wantErr: false,
		},
		{
			name: "fail",
			args: args{
				ctx: context.TODO(),
				id:  1,
			},
			initMock: func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				p := expectPrepareMock(mock)
				p.deleteMonster.ExpectExec().WithArgs(
					1,
				).WillReturnError(errors.New("mock error"))

				return sqlx.NewDb(db, "postgres"), db, mock
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, dbMock, mock := tt.initMock()
			defer dbMock.Close()
			q, _ := prepareQueries(db)

			re := monsterRepository{
				query: q,
			}
			err := re.DeleteMonster(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.DeleteMonster() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err.Error())
			}
		})
	}
}
