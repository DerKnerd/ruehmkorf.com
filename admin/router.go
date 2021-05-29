package admin

import (
	"net/http"
	"ruehmkorf.com/admin/middleware"
	"ruehmkorf.com/admin/routes"
	middleware2 "ruehmkorf.com/middleware"
)

func InitRouting(mux *http.ServeMux) error {
	mux.HandleFunc("/admin/login", middleware2.NoIndexMiddleware(routes.Login))
	mux.HandleFunc("/admin/twofactor", middleware2.NoIndexMiddleware(routes.TwoFactor))

	mux.HandleFunc("/admin/news/", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.NewsList)))
	mux.HandleFunc("/admin/news/new/", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.NewsNew)))
	mux.HandleFunc("/admin/news/edit/", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.NewsEdit)))
	mux.HandleFunc("/admin/news/delete/", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.NewsDelete)))

	mux.HandleFunc("/admin/profile/", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.ProfileList)))
	mux.HandleFunc("/admin/profile/new/", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.ProfileNew)))
	mux.HandleFunc("/admin/profile/edit/", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.ProfileEdit)))
	mux.HandleFunc("/admin/profile/delete/", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.ProfileDelete)))

	mux.HandleFunc("/admin/download/", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.DownloadList)))
	mux.HandleFunc("/admin/download/new/", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.DownloadNew)))
	mux.HandleFunc("/admin/download/edit/", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.DownloadEdit)))
	mux.HandleFunc("/admin/download/delete/", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.DownloadDelete)))

	mux.HandleFunc("/admin/buchstabieromat/", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.Buchstabieromat)))

	mux.HandleFunc("/admin/", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.Home)))

	return nil
}
