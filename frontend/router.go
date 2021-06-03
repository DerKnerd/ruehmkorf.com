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

	mux.HandleFunc("/de/news", middleware.ErrorHandlerMiddleware(middleware.LanguageDetectorMiddleware(routes.NewsList)))
	mux.HandleFunc("/en/news", middleware.ErrorHandlerMiddleware(middleware.LanguageDetectorMiddleware(routes.NewsList)))
	mux.HandleFunc("/news", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/en/news", http.StatusFound)
	})

	mux.HandleFunc("/de/news/", middleware.ErrorHandlerMiddleware(middleware.LanguageDetectorMiddleware(routes.NewsPage)))
	mux.HandleFunc("/en/news/", middleware.ErrorHandlerMiddleware(middleware.LanguageDetectorMiddleware(routes.NewsPage)))
	mux.HandleFunc("/news/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/en"+r.URL.Path, http.StatusFound)
	})

	mux.HandleFunc("/de/buchstabier-o-mat/", middleware.ErrorHandlerMiddleware(middleware.LanguageDetectorMiddleware(routes.BuchstabieroMatPage)))
	mux.HandleFunc("/en/buchstabier-o-mat/", middleware.ErrorHandlerMiddleware(middleware.LanguageDetectorMiddleware(routes.BuchstabieroMatPage)))
	mux.HandleFunc("/buchstabier-o-mat/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/en"+r.URL.Path, http.StatusFound)
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/en/news", http.StatusFound)
	})
	mux.HandleFunc("/en", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/en/news", http.StatusFound)
	})
	mux.HandleFunc("/de", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/de/news", http.StatusFound)
	})

	mux.HandleFunc("/en/", middleware.ErrorHandlerMiddleware(routes.Error404Handler))
	mux.HandleFunc("/de/", middleware.ErrorHandlerMiddleware(routes.Error404Handler))

	return nil
}
