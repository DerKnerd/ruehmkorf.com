package frontend

import (
	"net/http"
	"ruehmkorf.com/frontend/routes"
)

func InitRouting(mux *http.ServeMux) error {
	mux.HandleFunc("/news/hero/", routes.NewsHeroImage)

	mux.HandleFunc("/profile/icon/", routes.ProfileIcon)
	mux.HandleFunc("/profile/header/", routes.ProfileHeader)

	mux.HandleFunc("/download/file/", routes.DownloadFile)
	mux.HandleFunc("/download/preview/", routes.DownloadPreview)

	return nil
}
