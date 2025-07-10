package server

import (
	"diffinlist/internal/auth"
	db "diffinlist/internal/db/generated"
	"net/http"
)

func setupAuthRoutes(mux *http.ServeMux, queries *db.Queries) {
	mux.HandleFunc("POST /sign-up/{$}", func(w http.ResponseWriter, r *http.Request) {
		auth.HandleSignup(queries, w, r)
	})

	mux.HandleFunc("POST /login/{$}", func(w http.ResponseWriter, r *http.Request) {
		auth.HandleLogin(queries, w, r)
	})

	mux.HandleFunc("POST /logout/{$}", func(w http.ResponseWriter, r *http.Request) {
		auth.HandleLogout(queries, w, r)
	})
}
