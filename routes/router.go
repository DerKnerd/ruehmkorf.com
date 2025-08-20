package routes

import (
	"ruehmkorf/routes/admin"

	"github.com/gorilla/mux"
)

func SetupRouter(router *mux.Router) {
	admin.SetupRouter(router)
}
