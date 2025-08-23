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

func spot(w http.ResponseWriter, r *http.Request) {
	spotMappings, err := database.Select[database.SpotMapping]("select * from spot_mapping")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	renderPage(w, r, "spot", struct {
		SpotMappings []database.SpotMapping
	}{
		SpotMappings: spotMappings,
	})
}

func SetupRouter(router *mux.Router) {
	router.Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/profiles", http.StatusFound)
	})
	router.Path("/profiles").HandlerFunc(profiles)
	router.Path("/spell-o-tron").HandlerFunc(spot)
}
