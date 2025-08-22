package admin

import (
	"ruehmkorf/routes/admin/api"
	"ruehmkorf/routes/admin/web"

	"github.com/gorilla/mux"
)

func SetupRouter(router *mux.Router) {
	api.SetupRouter(router)
	web.SetupRouter(router)
}
