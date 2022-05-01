package frontend

import (
	"net/http"
	"ruehmkorf.com/frontend/middleware"
	"ruehmkorf.com/frontend/routes"
)

func InitRouting(mux *http.ServeMux) error {
	mux.HandleFunc("/news/hero/", routes.NewsHeroImage)

	mux.HandleFunc("/profile/icon/", routes.ProfileIcon)
	mux.HandleFunc("/profile/header/", routes.ProfileHeader)

	mux.HandleFunc("/download/file/", routes.DownloadFile)
	mux.HandleFunc("/download/preview/", routes.DownloadPreview)

	mux.HandleFunc("/de/download/", middleware.ErrorHandlerMiddleware(middleware.LanguageDetectorMiddleware(routes.DownloadPage)))
	mux.HandleFunc("/en/download/", middleware.ErrorHandlerMiddleware(middleware.LanguageDetectorMiddleware(routes.DownloadPage)))
	mux.HandleFunc("/download/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/en"+r.URL.Path, http.StatusFound)
	})

	mux.HandleFunc("/de/spellotron/", middleware.ErrorHandlerMiddleware(middleware.LanguageDetectorMiddleware(routes.BuchstabieroMatPage)))
	mux.HandleFunc("/en/spellotron/", middleware.ErrorHandlerMiddleware(middleware.LanguageDetectorMiddleware(routes.BuchstabieroMatPage)))
	mux.HandleFunc("/spellotron/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/en"+r.URL.Path, http.StatusFound)
	})

	mux.HandleFunc("/de/profile/", middleware.ErrorHandlerMiddleware(middleware.LanguageDetectorMiddleware(routes.ProfileList)))
	mux.HandleFunc("/en/profile/", middleware.ErrorHandlerMiddleware(middleware.LanguageDetectorMiddleware(routes.ProfileList)))
	mux.HandleFunc("/profile/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/en"+r.URL.Path, http.StatusFound)
	})

	mux.HandleFunc("/de/data-protection/", middleware.ErrorHandlerMiddleware(middleware.LanguageDetectorMiddleware(routes.DataProtectionPage)))
	mux.HandleFunc("/en/data-protection/", middleware.ErrorHandlerMiddleware(middleware.LanguageDetectorMiddleware(routes.DataProtectionPage)))
	mux.HandleFunc("/data-protection/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/en"+r.URL.Path, http.StatusFound)
	})

	mux.HandleFunc("/de/cookies/", middleware.ErrorHandlerMiddleware(middleware.LanguageDetectorMiddleware(routes.CookiesPage)))
	mux.HandleFunc("/en/cookies/", middleware.ErrorHandlerMiddleware(middleware.LanguageDetectorMiddleware(routes.CookiesPage)))
	mux.HandleFunc("/cookies/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/en"+r.URL.Path, http.StatusFound)
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/en/profile", http.StatusFound)
	})
	mux.HandleFunc("/en", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/en/profile", http.StatusFound)
	})
	mux.HandleFunc("/de", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/de/profile", http.StatusFound)
	})

	mux.HandleFunc("/en/", middleware.ErrorHandlerMiddleware(routes.Error404Handler))
	mux.HandleFunc("/de/", middleware.ErrorHandlerMiddleware(routes.Error404Handler))

	return nil
}
