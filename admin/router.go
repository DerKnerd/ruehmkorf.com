package admin

import (
	"net/http"
	"ruehmkorf.com/admin/middleware"
	"ruehmkorf.com/admin/routes"
)

func InitRouting(mux *http.ServeMux) error {
	mux.HandleFunc("/admin/login", routes.Login)
	mux.HandleFunc("/admin/twofactor", routes.TwoFactor)

	mux.HandleFunc("/admin/news/", middleware.CheckLoginMiddleware(routes.NewsList))

	mux.HandleFunc("/admin/", middleware.CheckLoginMiddleware(routes.Home))
	return nil
}
