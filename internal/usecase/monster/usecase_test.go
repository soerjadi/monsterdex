package monster

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/soerjadi/monsterdex/internal/mocks"
	"github.com/soerjadi/monsterdex/internal/model"
	"github.com/soerjadi/monsterdex/internal/repository/monster"
)

func TestGetListMonster(t *testing.T) {
	type fields struct {
		repository monster.Repository
	}

	type args struct {
		ctx    context.Context
		req    model.MonsterListRequest
		userID int64
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.Monster
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx:    context.TODO(),
				req:    model.MonsterListRequest{},
				userID: 1,
			},
			fields: fields{
				repository: func() monster.Repository {
					ctrl := gomock.NewController(t)
					mock := mocks.NewMockMonsterRepository(ctrl)

					mock.EXPECT().
						GetListMonster(gomock.Any(), gomock.Any(), gomock.Any()).
						Return([]model.Monster{
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
						}, nil)

					return mock
				}(),
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
		},
		{
			name: "fail",
			args: args{
				ctx:    context.TODO(),
				req:    model.MonsterListRequest{},
				userID: 1,
			},
			fields: fields{
				repository: func() monster.Repository {
					ctrl := gomock.NewController(t)
					mock := mocks.NewMockMonsterRepository(ctrl)

					mock.EXPECT().
						GetListMonster(gomock.Any(), gomock.Any(), gomock.Any()).
						Return(nil, errors.New("mock error"))

					return mock
				}(),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			au := &monsterUsecase{
				repository: tt.fields.repository,
			}
			got, err := au.GetListMonster(tt.args.ctx, tt.args.req, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.GetListMonster() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usecase.GetListMonster() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDetailMonster(t *testing.T) {
	type fields struct {
		repository monster.Repository
	}

	type args struct {
		ctx    context.Context
		id     int64
		userID int64
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.Monster
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx:    context.TODO(),
				id:     1,
				userID: 1,
			},
			fields: fields{
				repository: func() monster.Repository {
					ctrl := gomock.NewController(t)
					mock := mocks.NewMockMonsterRepository(ctrl)

					mock.EXPECT().
						GetDetailMonster(gomock.Any(), gomock.Any(), gomock.Any()).
						Return(model.Monster{
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
						}, nil)

					return mock
				}(),
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
		},
		{
			name: "fail",
			args: args{
				ctx:    context.TODO(),
				id:     1,
				userID: 1,
			},
			fields: fields{
				repository: func() monster.Repository {
					ctrl := gomock.NewController(t)
					mock := mocks.NewMockMonsterRepository(ctrl)

					mock.EXPECT().
						GetDetailMonster(gomock.Any(), gomock.Any(), gomock.Any()).
						Return(model.Monster{}, errors.New("mock error"))

					return mock
				}(),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			au := &monsterUsecase{
				repository: tt.fields.repository,
			}
			got, err := au.GetDetailMonster(tt.args.ctx, tt.args.id, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.GetDetailMonster() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usecase.GetDetailMonster() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCaptureMonster(t *testing.T) {
	type fields struct {
		repository monster.Repository
	}

	type args struct {
		ctx context.Context
		req model.CaptureMonsterReq
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.CapturedMonster
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.TODO(),
				req: model.CaptureMonsterReq{},
			},
			fields: fields{
				repository: func() monster.Repository {
					ctrl := gomock.NewController(t)
					mock := mocks.NewMockMonsterRepository(ctrl)

					mock.EXPECT().
						CaptureMonster(gomock.Any(), gomock.Any()).
						Return(model.CapturedMonster{
							ID:            1,
							UserID:        1,
							MonsterID:     1,
							CaptureStatus: true,
						}, nil)

					return mock
				}(),
			},
			want: model.CapturedMonster{
				ID:            1,
				UserID:        1,
				MonsterID:     1,
				CaptureStatus: true,
			},
		},
		{
			name: "fail",
			args: args{
				ctx: context.TODO(),
				req: model.CaptureMonsterReq{},
			},
			fields: fields{
				repository: func() monster.Repository {
					ctrl := gomock.NewController(t)
					mock := mocks.NewMockMonsterRepository(ctrl)

					mock.EXPECT().
						CaptureMonster(gomock.Any(), gomock.Any()).
						Return(model.CapturedMonster{}, errors.New("mock error"))

					return mock
				}(),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			au := &monsterUsecase{
				repository: tt.fields.repository,
			}
			got, err := au.CaptureMonster(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.CaptureMonster() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usecase.CaptureMonster() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertMonster(t *testing.T) {
	type fields struct {
		repository monster.Repository
	}

	type args struct {
		ctx context.Context
		req model.MonsterRequest
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.MonsterData
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.TODO(),
				req: model.MonsterRequest{},
			},
			fields: fields{
				repository: func() monster.Repository {
					ctrl := gomock.NewController(t)
					mock := mocks.NewMockMonsterRepository(ctrl)

					mock.EXPECT().
						InsertMonster(gomock.Any(), gomock.Any()).
						Return(model.MonsterData{
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
						}, nil)

					return mock
				}(),
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
		},
		{
			name: "fail",
			args: args{
				ctx: context.TODO(),
				req: model.MonsterRequest{},
			},
			fields: fields{
				repository: func() monster.Repository {
					ctrl := gomock.NewController(t)
					mock := mocks.NewMockMonsterRepository(ctrl)

					mock.EXPECT().
						InsertMonster(gomock.Any(), gomock.Any()).
						Return(model.MonsterData{}, errors.New("mock error"))

					return mock
				}(),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			au := &monsterUsecase{
				repository: tt.fields.repository,
			}
			got, err := au.InsertMonster(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.InsertMonster() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usecase.InsertMonster() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateMonster(t *testing.T) {
	type fields struct {
		repository monster.Repository
	}

	type args struct {
		ctx context.Context
		req model.MonsterRequest
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.TODO(),
				req: model.MonsterRequest{},
			},
			fields: fields{
				repository: func() monster.Repository {
					ctrl := gomock.NewController(t)
					mock := mocks.NewMockMonsterRepository(ctrl)

					mock.EXPECT().
						UpdateMonster(gomock.Any(), gomock.Any()).
						Return(nil)

					return mock
				}(),
			},
		},
		{
			name: "fail",
			args: args{
				ctx: context.TODO(),
				req: model.MonsterRequest{},
			},
			fields: fields{
				repository: func() monster.Repository {
					ctrl := gomock.NewController(t)
					mock := mocks.NewMockMonsterRepository(ctrl)

					mock.EXPECT().
						UpdateMonster(gomock.Any(), gomock.Any()).
						Return(errors.New("mock error"))

					return mock
				}(),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			au := &monsterUsecase{
				repository: tt.fields.repository,
			}
			err := au.UpdateMonster(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.UpdateMonster() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDeleteMonster(t *testing.T) {
	type fields struct {
		repository monster.Repository
	}

	type args struct {
		ctx context.Context
		id  int64
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.TODO(),
				id:  1,
			},
			fields: fields{
				repository: func() monster.Repository {
					ctrl := gomock.NewController(t)
					mock := mocks.NewMockMonsterRepository(ctrl)

					mock.EXPECT().
						DeleteMonster(gomock.Any(), gomock.Any()).
						Return(nil)

					return mock
				}(),
			},
		},
		{
			name: "fail",
			args: args{
				ctx: context.TODO(),
				id:  1,
			},
			fields: fields{
				repository: func() monster.Repository {
					ctrl := gomock.NewController(t)
					mock := mocks.NewMockMonsterRepository(ctrl)

					mock.EXPECT().
						DeleteMonster(gomock.Any(), gomock.Any()).
						Return(errors.New("mock error"))

					return mock
				}(),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			au := &monsterUsecase{
				repository: tt.fields.repository,
			}
			err := au.DeleteMonster(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.DeleteMonster() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
