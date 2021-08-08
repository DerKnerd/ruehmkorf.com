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

	mux.HandleFunc("/admin/news", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.NewsAction)))
	mux.HandleFunc("/admin/news/hero", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.UploadHero)))

	mux.HandleFunc("/admin/profile/", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.ProfileList)))
	mux.HandleFunc("/admin/profile/new/", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.ProfileNew)))
	mux.HandleFunc("/admin/profile/edit/", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.ProfileEdit)))
	mux.HandleFunc("/admin/profile/delete/", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.ProfileDelete)))

	mux.HandleFunc("/admin/download/", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.DownloadList)))
	mux.HandleFunc("/admin/download/new/", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.DownloadNew)))
	mux.HandleFunc("/admin/download/edit/", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.DownloadEdit)))
	mux.HandleFunc("/admin/download/delete/", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.DownloadDelete)))

	mux.HandleFunc("/admin/buchstabieromat", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.Buchstabieromat)))

	mux.HandleFunc("/admin/user", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.UserAction)))

	mux.HandleFunc("/admin/settings", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.SettingsAction)))
	mux.HandleFunc("/admin/settings/touchicon", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.LogosAction)))
	mux.HandleFunc("/admin/settings/favicon", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.LogosAction)))
	mux.HandleFunc("/admin/settings/logo", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.LogosAction)))

	mux.HandleFunc("/admin/", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.Home)))

	return nil
}
