package routes

import (
	"ruehmkorf/routes/admin"
	"ruehmkorf/routes/frontend"

	"github.com/gorilla/mux"
)

func SetupRouter(router *mux.Router) {
	admin.SetupRouter(router)
	frontend.SetupRouter(router)
}
