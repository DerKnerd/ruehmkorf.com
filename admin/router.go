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
	mux.HandleFunc("/admin/news/new/", middleware.CheckLoginMiddleware(routes.NewsNew))
	mux.HandleFunc("/admin/news/edit/", middleware.CheckLoginMiddleware(routes.NewsEdit))
	mux.HandleFunc("/admin/news/delete/", middleware.CheckLoginMiddleware(routes.NewsDelete))

	mux.HandleFunc("/admin/", middleware.CheckLoginMiddleware(routes.Home))
	return nil
}
