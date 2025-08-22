package web

import (
	"github.com/gorilla/mux"
)

func SetupRouter(router *mux.Router) {
	router.
		PathPrefix("/admin").
		HandlerFunc(indexPage)
}
