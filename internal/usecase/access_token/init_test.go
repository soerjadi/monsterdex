package access_token

import (
	"reflect"
	"testing"

	"github.com/soerjadi/monsterdex/internal/repository/access_token"
)

func TestGetUsecase(t *testing.T) {
	type args struct {
		repository access_token.Repository
	}
	tests := []struct {
		name string
		args args
		want Usecase
	}{
		{
			name: "soerja",
			want: &tokenUsecase{},
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
