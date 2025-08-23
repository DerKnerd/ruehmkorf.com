package frontend

import (
	"net/http"
	"ruehmkorf/database"

	"github.com/gorilla/mux"
)

func profiles(w http.ResponseWriter, r *http.Request) {
	profiles, err := database.Select[database.Profile]("select * from profile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	renderPage(w, r, "profiles", struct {
		Profiles []database.Profile
	}{profiles})
}

func SetupRouter(router *mux.Router) {
	router.Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/profiles", http.StatusFound)
	})
	router.Path("/profiles").HandlerFunc(profiles)
}
