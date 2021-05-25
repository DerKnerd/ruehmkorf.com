package routes

import (
	"net/http"
	httpUtils "ruehmkorf.com/utils/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	httpUtils.RenderAdmin("admin/templates/home/home.gohtml", nil, w)
}
