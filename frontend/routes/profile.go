package routes

import (
	"io/ioutil"
	"net/http"
	"ruehmkorf.com/database/models"
	"strings"
)

func ProfileIcon(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/profile/icon/")
	profile, err := models.FindProfileById(id)
	if err != nil || !profile.Active {
		http.NotFound(w, r)
		return
	}

	data, err := ioutil.ReadFile(profile.Icon)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func ProfileHeader(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/profile/header/")
	profile, err := models.FindProfileById(id)
	if err != nil || !profile.Active || !profile.Header.Valid {
		http.NotFound(w, r)
		return
	}

	data, err := ioutil.ReadFile(profile.Header.String)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
