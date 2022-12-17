package monster

import (
	"reflect"
	"testing"

	"github.com/soerjadi/monsterdex/internal/repository/monster"
)

func TestGetUsecase(t *testing.T) {
	type args struct {
		repository monster.Repository
	}
	tests := []struct {
		name string
		args args
		want Usecase
	}{
		{
			name: "soerja",
			want: &monsterUsecase{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUsecase(tt.args.repository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}
