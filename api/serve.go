package api

import (
	"log"
	"net/http"
	"profile/consts"

	"github.com/gorilla/handlers"
)

const (
	SERVERPREFIX = "SERVERPREFIX"
	SERVERPORT   = "SERVERPORT"
)

// Serve : Run api server
func (a *App) Serve(env map[string]string) {
	router := a.Router(env[SERVERPREFIX])
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"})
	log.Println(consts.App + " : " + "Server is listening...")
	http.ListenAndServe(env[SERVERPORT], handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(router))
}
