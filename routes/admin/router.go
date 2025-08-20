package admin

import (
	"ruehmkorf/routes/admin/api"

	"github.com/gorilla/mux"
)

func SetupRouter(router *mux.Router) {
	api.SetupRouter(router)
}
