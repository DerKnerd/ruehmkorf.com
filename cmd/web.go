package cmd

import (
	"io/fs"
	"log"
	"net/http"
	"os"
	"ruehmkorf/routes"

	"github.com/gorilla/mux"
)

func WebUi(static fs.FS) {
	router := mux.NewRouter()

	if os.Getenv("ENV") == "dev" {
		router.PathPrefix("/static/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Service-Worker-Allowed", "/")
			http.FileServerFS(os.DirFS(".")).ServeHTTP(w, r)
		})
	} else {
		router.PathPrefix("/static/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Service-Worker-Allowed", "/")
			http.FileServerFS(static).ServeHTTP(w, r)
		})
	}

	routes.SetupRouter(router)

	listenAddress := os.Getenv("LISTEN_ADDRESS")
	if listenAddress == "" {
		listenAddress = ":8090"
	}

	log.Printf("Serving at %s...", listenAddress)
	err := http.ListenAndServe(listenAddress, router)
	if err != nil {
		panic(err)
	}
}
