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

	mux.HandleFunc("/admin/profile", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.ProfileAction)))
	mux.HandleFunc("/admin/profile/header", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.UploadProfileImage)))
	mux.HandleFunc("/admin/profile/icon", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.UploadProfileImage)))

	mux.HandleFunc("/admin/download", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.DownloadAction)))
	mux.HandleFunc("/admin/download/preview", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.UploadPreviewAction)))
	mux.HandleFunc("/admin/download/file/chunk", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.UploadFileAction)))
	mux.HandleFunc("/admin/download/file/finish", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.UploadFileAction)))

	mux.HandleFunc("/admin/buchstabieromat", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.Buchstabieromat)))

	mux.HandleFunc("/admin/user", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.UserAction)))

	mux.HandleFunc("/admin/settings", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.SettingsAction)))
	mux.HandleFunc("/admin/settings/touchicon", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.LogosAction)))
	mux.HandleFunc("/admin/settings/favicon", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.LogosAction)))
	mux.HandleFunc("/admin/settings/logo", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.LogosAction)))

	mux.HandleFunc("/admin/", middleware2.NoIndexMiddleware(middleware.CheckLoginMiddleware(routes.Home)))

	return nil
}
