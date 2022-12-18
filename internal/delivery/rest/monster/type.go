package monster

import (
	"github.com/soerjadi/monsterdex/internal/usecase/access_token"
	"github.com/soerjadi/monsterdex/internal/usecase/monster"
	"github.com/soerjadi/monsterdex/internal/usecase/user"
)

type Handler struct {
	usecase monster.Usecase
	auth    access_token.Usecase
	user    user.Usecase
}
