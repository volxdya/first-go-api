package api

import (
	"log"
	"net/http"
)

type api struct {
	addr string
	r    *http.ServeMux
}

func New(addr string, r *http.ServeMux) *api {
	return &api{addr: addr, r: r}
}

func (api *api) FillEndpoints() {
	api.r.HandleFunc("/api/hello", hello)
	api.r.HandleFunc("/api/goodbye", goodbye)
}

func (api *api) ListenAndServe(serverUrl string) error {
	log.Println("server started on", serverUrl)
	return http.ListenAndServe(api.addr, api.r)
}
