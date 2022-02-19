package routes

import (
	"crypto/md5"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"ruehmkorf.com/database/models"
	"strings"
)

func getProfileIcon(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	profile, err := models.FindProfileById(id)

	data, err := ioutil.ReadFile(profile.Icon)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func getProfileHeader(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	profile, err := models.FindProfileById(id)

	data, err := ioutil.ReadFile(profile.Header.String)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func ProfileAction(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := r.URL.Query().Get("id")
		if id != "" {
			profileDetails(w, id)
		} else {
			profileList(w)
		}
	} else if r.Method == http.MethodPost {
		profileNew(w, r)
	} else if r.Method == http.MethodPut {
		profileEdit(w, r)
	} else if r.Method == http.MethodDelete {
		profileDelete(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func profileList(w http.ResponseWriter) {
	profiles, err := models.FindAllProfiles()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(profiles)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func profileDetails(w http.ResponseWriter, id string) {
	profile, err := models.FindProfileById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(profile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type profileData struct {
	Name   string `json:"name"`
	Url    string `json:"url"`
	Active bool   `json:"active"`
}

func saveProfileImage(reader io.ReadCloser, icon bool) (string, error) {
	path := models.HeaderPath
	if icon {
		path = models.IconPath
	}

	err := os.MkdirAll(path, 0755)
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}

	hash := md5.New()
	hash.Write(data)
	id := strings.ReplaceAll(base64.RawStdEncoding.EncodeToString(hash.Sum(nil)), "/", "")
	err = os.WriteFile(path+id, data, 0755)

	return path + id, err
}

func profileNew(w http.ResponseWriter, r *http.Request) {
	var data profileData

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	profile := models.Profile{
		Name:   data.Name,
		Url:    data.Url,
		Active: data.Active,
		Icon:   "",
		Header: sql.NullString{String: "", Valid: true},
	}

	id, err := models.CreateProfile(profile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(id))
}

func UploadProfileImage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if strings.HasSuffix(r.URL.Path, "icon") {
			getProfileIcon(w, r)
		} else if strings.HasSuffix(r.URL.Path, "header") {
			getProfileHeader(w, r)
		}
	} else if r.Method == http.MethodPost {
		id := r.URL.Query().Get("id")
		profile, err := models.FindProfileById(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		path, err := saveProfileImage(r.Body, strings.HasSuffix(r.URL.Path, "icon"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if strings.HasSuffix(r.URL.Path, "icon") {
			profile.Icon = path
		} else {
			profile.Header = sql.NullString{String: path, Valid: true}
		}
		err = models.UpdateProfile(*profile)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func profileEdit(w http.ResponseWriter, r *http.Request) {
	profile, err := models.FindProfileById(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var data profileData
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	profile.Name = data.Name
	profile.Url = data.Url
	profile.Active = data.Active

	err = models.UpdateProfile(*profile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func profileDelete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	profile, err := models.FindProfileById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = models.DeleteProfile(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_ = os.Remove(profile.Icon)
	_ = os.Remove(profile.Header.String)

	w.WriteHeader(http.StatusNoContent)
}
