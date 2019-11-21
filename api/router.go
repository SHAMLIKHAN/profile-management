package api

import "github.com/gorilla/mux"

// Router : Basic router
func (a *App) Router(prefix string) *mux.Router {
	r := mux.NewRouter()
	return r
}
