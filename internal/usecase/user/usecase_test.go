package user

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/soerjadi/monsterdex/internal/mocks"
	"github.com/soerjadi/monsterdex/internal/model"
	"github.com/soerjadi/monsterdex/internal/repository/user"
)

func TestGetUserByID(t *testing.T) {
	type fields struct {
		repository user.Repository
	}

	type args struct {
		ctx context.Context
		id  int64
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.User
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.TODO(),
				id:  1,
			},
			fields: fields{
				repository: func() user.Repository {
					ctrl := gomock.NewController(t)
					mock := mocks.NewMockUserRepository(ctrl)

					mock.EXPECT().
						GetUserByID(gomock.Any(), gomock.Any()).
						Return(model.User{
							ID:        1,
							Name:      "user",
							Email:     "user@main.com",
							Role:      0,
							CreatedAt: time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC),
							UpdatedAt: time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC),
						}, nil)

					return mock
				}(),
			},
			want: model.User{
				ID:        1,
				Name:      "user",
				Email:     "user@main.com",
				Role:      0,
				CreatedAt: time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2021, 1, 12, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "fail",
			args: args{
				ctx: context.TODO(),
				id:  1,
			},
			fields: fields{
				repository: func() user.Repository {
					ctrl := gomock.NewController(t)
					mock := mocks.NewMockUserRepository(ctrl)

					mock.EXPECT().
						GetUserByID(gomock.Any(), gomock.Any()).
						Return(model.User{}, errors.New("mock error"))

					return mock
				}(),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			au := &userUsecase{
				repository: tt.fields.repository,
			}
			got, err := au.GetUserByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usecase.GetDetailMonster() = %v, want %v", got, tt.want)
			}
		})
	}
}
