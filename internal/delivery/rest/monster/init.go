package monster

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/soerjadi/monsterdex/internal/delivery/rest"
	"github.com/soerjadi/monsterdex/internal/delivery/rest/middleware"
	"github.com/soerjadi/monsterdex/internal/usecase/access_token"
	"github.com/soerjadi/monsterdex/internal/usecase/monster"
	"github.com/soerjadi/monsterdex/internal/usecase/user"
)

func NewHandler(usecase monster.Usecase, authUsecase access_token.Usecase, userUsecase user.Usecase) rest.API {
	return &Handler{
		usecase: usecase,
		auth:    authUsecase,
		user:    userUsecase,
	}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	monsters := r.PathPrefix("/monsters").Subrouter()
	detailMonster := r.PathPrefix("/monster").Subrouter()

	monsters.Use(middleware.CheckLoggedUser(h.auth, h.user))
	detailMonster.Use(middleware.CheckLoggedUser(h.auth, h.user))

	monsters.HandleFunc("/", rest.HandlerFunc(h.GetList).Serve).Methods(http.MethodGet)
	detailMonster.HandleFunc("/{monster_id}", rest.HandlerFunc(h.GetDetailMonster).Serve).Methods(http.MethodGet)
	detailMonster.HandleFunc("/{monster_id}/capture", rest.HandlerFunc(h.CapturedMonster).Serve).Methods(http.MethodPost)

	// admin section
	internal := r.PathPrefix("/admin/monster").Subrouter()

	internal.Use(middleware.OnlyAdminLoggedIn(h.auth, h.user))

	internal.HandleFunc("/", rest.HandlerFunc(h.InsertMonster).Serve).Methods(http.MethodPost)
	internal.HandleFunc("/{monster_id}", rest.HandlerFunc(h.UpdateMonster).Serve).Methods(http.MethodPut)
	internal.HandleFunc("/{monster_id}", rest.HandlerFunc(h.DeleteMonster).Serve).Methods(http.MethodDelete)

}
