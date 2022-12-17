package access_token

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

func TestGetUserIDByToken(t *testing.T) {
	type args struct {
		ctx   context.Context
		token string
	}
	tests := []struct {
		name     string
		args     args
		initMock func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock)
		want     model.AccessToken
		wantErr  bool
	}{
		{
			name: "success",
			args: args{
				ctx:   context.TODO(),
				token: "token",
			},
			initMock: func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				p := expectPrepareMock(mock)
				p.getUserIDByToken.ExpectQuery().WillReturnRows(sqlmock.NewRows([]string{
					"token",
					"user_id",
					"created_at",
					"valid_thru",
				}).
					AddRow("token", 1, time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC), time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC)))

				return sqlx.NewDb(db, "postgres"), db, mock
			},
			want: model.AccessToken{
				Token:     "token",
				UserID:    1,
				CreatedAt: time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC),
				ValidThru: time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name: "fail",
			args: args{
				ctx:   context.TODO(),
				token: "token",
			},
			initMock: func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				p := expectPrepareMock(mock)
				p.getUserIDByToken.ExpectQuery().WillReturnError(errors.New("mock error"))

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

			re := tokenRepository{
				query: q,
			}
			got, err := re.GetUserIDByToken(tt.args.ctx, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetUserIDByToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetUserIDByToken() = %v, want %v", got, tt.want)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err.Error())
			}
		})
	}
}
