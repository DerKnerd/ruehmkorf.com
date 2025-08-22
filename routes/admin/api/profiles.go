package api

import (
	"encoding/json"
	"net/http"
	"ruehmkorf/database"

	"github.com/gorilla/mux"
)

func getAllProfiles(w http.ResponseWriter, r *http.Request) {
	profiles, err := database.Select[database.Profile]("select * from profile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(profiles)
}

func getProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	profile, err := database.Get[database.Profile](id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(profile)
}

func createProfile(w http.ResponseWriter, r *http.Request) {
	var body struct {
		LinkTarget  string `json:"linkTarget"`
		LinkLabel   string `json:"linkLabel"`
		Description string `json:"description"`
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	profile := &database.Profile{
		LinkTarget:  body.LinkTarget,
		LinkLabel:   body.LinkLabel,
		Description: body.Description,
	}
	err = database.GetDbMap().Insert(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	w.WriteHeader(http.StatusCreated)
	encoder.Encode(profile)
}

func updateProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	profile, err := database.Get[database.Profile](id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var body struct {
		LinkTarget  string `json:"linkTarget"`
		LinkLabel   string `json:"linkLabel"`
		Description string `json:"description"`
	}

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	profile.LinkTarget = body.LinkTarget
	profile.LinkLabel = body.LinkLabel
	profile.Description = body.Description

	_, err = database.GetDbMap().Update(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func deleteProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	profile, err := database.Get[database.Profile](id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	_, err = database.GetDbMap().Delete(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
