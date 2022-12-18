package monster

import (
	"reflect"
	"testing"

	"github.com/gorilla/mux"
	"github.com/soerjadi/monsterdex/internal/delivery/rest"
	"github.com/soerjadi/monsterdex/internal/usecase/access_token"
	"github.com/soerjadi/monsterdex/internal/usecase/monster"
	"github.com/soerjadi/monsterdex/internal/usecase/user"
)

func TestNewHandler(t *testing.T) {
	type args struct {
		usecase monster.Usecase
		auth    access_token.Usecase
		user    user.Usecase
	}
	tests := []struct {
		name string
		args args
		want rest.API
	}{
		{
			name: "success",
			args: args{
				usecase: nil,
				auth:    nil,
				user:    nil,
			},
			want: &Handler{
				usecase: nil,
				user:    nil,
				auth:    nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHandler(tt.args.usecase, tt.args.auth, tt.args.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandler_RegisterRoutes(t *testing.T) {
	type fields struct {
		usecase monster.Usecase
		auth    access_token.Usecase
		user    user.Usecase
	}
	type args struct {
		r *mux.Router
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "pass",
			args: args{
				r: mux.NewRouter(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				usecase: tt.fields.usecase,
				auth:    tt.fields.auth,
				user:    tt.fields.user,
			}
			h.RegisterRoutes(tt.args.r)
		})
	}
}
