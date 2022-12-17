package access_token

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/soerjadi/monsterdex/internal/mocks"
	"github.com/soerjadi/monsterdex/internal/model"
	"github.com/soerjadi/monsterdex/internal/repository/access_token"
)

func TestGetUserIDByToken(t *testing.T) {
	type fields struct {
		repository access_token.Repository
	}

	type args struct {
		ctx   context.Context
		token string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.AccessToken
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx:   context.TODO(),
				token: "token",
			},
			fields: fields{
				repository: func() access_token.Repository {
					ctrl := gomock.NewController(t)
					mock := mocks.NewMockAccessTokenRepository(ctrl)

					mock.EXPECT().
						GetUserIDByToken(gomock.Any(), gomock.Any()).
						Return(model.AccessToken{
							Token:     "token",
							UserID:    1,
							CreatedAt: time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC),
							ValidThru: time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC),
						}, nil)

					return mock
				}(),
			},
			want: model.AccessToken{
				Token:     "token",
				UserID:    1,
				CreatedAt: time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC),
				ValidThru: time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "fail",
			args: args{
				ctx:   context.TODO(),
				token: "token",
			},
			fields: fields{
				repository: func() access_token.Repository {
					ctrl := gomock.NewController(t)
					mock := mocks.NewMockAccessTokenRepository(ctrl)

					mock.EXPECT().
						GetUserIDByToken(gomock.Any(), gomock.Any()).
						Return(model.AccessToken{}, errors.New("mock error"))

					return mock
				}(),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			au := &tokenUsecase{
				repository: tt.fields.repository,
			}
			got, err := au.GetUserIDByToken(tt.args.ctx, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.GetUserIDByToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usecase.GetDetailMonster() = %v, want %v", got, tt.want)
			}
		})
	}
}
