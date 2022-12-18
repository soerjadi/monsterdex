package monster

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/soerjadi/monsterdex/internal/mocks"
	"github.com/soerjadi/monsterdex/internal/model"
	"github.com/soerjadi/monsterdex/internal/usecase/monster"
)

type headerMock map[string][]string

func (hm *headerMock) Set(key string, value string) {
	return
}

type responseMock struct{}

func (rm *responseMock) Header() http.Header {
	var mock headerMock = make(map[string][]string)
	return http.Header(mock)
}

func (rm *responseMock) Write(data []byte) (int, error) {
	return 0, nil
}

func (rm *responseMock) WriteHeader(statusCode int) {
	return
}

func TestGetListMonster(t *testing.T) {
	type fields struct {
		usecase monster.Usecase
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "fail",
			fields: fields{
				usecase: func() monster.Usecase {
					ctrl := gomock.NewController(t)
					mock := mocks.NewMockMonsterUsecase(ctrl)

					mock.EXPECT().
						GetListMonster(gomock.Any(), gomock.Any(), gomock.Any()).
						Return(nil, errors.New("mock error"))

					return mock
				}(),
			},
			args: args{
				w: &responseMock{},
				r: func() *http.Request {
					r, _ := http.NewRequest("GET", "/", nil)

					return r
				}(),
			},
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				usecase: func() monster.Usecase {
					ctrl := gomock.NewController(t)
					mock := mocks.NewMockMonsterUsecase(ctrl)

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
			args: args{
				w: &responseMock{},
				r: func() *http.Request {
					r, _ := http.NewRequest("GET", "/", nil)

					return r
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				usecase: tt.fields.usecase,
			}
			got, err := h.GetList(tt.args.w, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler.GetList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler.GetList() = %v, want %v", got, tt.want)
			}
		})
	}
}
