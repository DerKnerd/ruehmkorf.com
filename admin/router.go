package admin

import (
	"net/http"
	"ruehmkorf.com/admin/routes"
)

func InitRouting(mux *http.ServeMux) error {
	mux.HandleFunc("/admin/", routes.Startpage)
	return nil
}
