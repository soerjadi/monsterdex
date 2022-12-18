package user

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/soerjadi/monsterdex/internal/model"
)

func TestGetUserByID(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name     string
		args     args
		initMock func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock)
		want     model.User
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
				p.getUserByID.ExpectQuery().WillReturnRows(sqlmock.NewRows([]string{
					"id",
					"name",
					"email",
					"role",
					"created_at",
					"updated_at",
				}).
					AddRow(1, "user", "user@main.com", 0, time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC), time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC)))

				return sqlx.NewDb(db, "postgres"), db, mock
			},
			want: model.User{
				ID:        1,
				Name:      "user",
				Email:     "user@main.com",
				Role:      0,
				CreatedAt: time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC),
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
				p.getUserByID.ExpectQuery().WillReturnError(errors.New("mock error"))

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

			re := userRepository{
				query: q,
			}
			got, err := re.GetUserByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetUserByID() = %v, want %v", got, tt.want)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err.Error())
			}
		})
	}
}
