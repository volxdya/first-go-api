package api

import (
	"myapi/pkg/repository"
	"net/http"

	"github.com/gorilla/mux"
)

type api struct {
	addr string
	r    *mux.Router
	psql *repository.PGRepo
}

func New(addr string, r *mux.Router, psql *repository.PGRepo) *api {
	return &api{addr: addr, r: r, psql: psql}
}

func (api *api) FillEndpoints() {
	api.r.HandleFunc("/api/users", api.users).Queries("id", "{id}")
	api.r.HandleFunc("/api/users", api.users)
}

func (api *api) ListenAndServe() error {
	return http.ListenAndServe(api.addr, api.r)
}
