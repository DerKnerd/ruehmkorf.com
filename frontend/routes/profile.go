package routes

import (
	"io/ioutil"
	"net/http"
	"os"
	"ruehmkorf.com/database/models"
	httpUtils "ruehmkorf.com/utils/http"
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

type profileListOverview struct {
	BaseData
	Items []profileListItem
}

type profileListItem struct {
	BaseData
	Name    string
	Url     string
	Id      string
	HasIcon bool
}

func ProfileList(w http.ResponseWriter, r *http.Request, language string) error {
	profiles, err := models.FindAllActiveProfiles()

	if err != nil {
		return err
	}

	var items []profileListItem
	for _, item := range profiles {
		_, err = os.Stat(item.Header.String)
		items = append(items, profileListItem{
			Name:    item.Name,
			Url:     item.Url,
			Id:      item.Id,
			HasIcon: !os.IsNotExist(err),
		})
	}
	cleanUrl := ""
	if language == "de" {
		cleanUrl = strings.TrimPrefix(r.URL.Path, "/de/")
	} else {
		cleanUrl = strings.TrimPrefix(r.URL.Path, "/en/")
	}

	return httpUtils.RenderFrontend("frontend/templates/profile/index.gohtml", profileListOverview{
		BaseData: BaseData{
			Language: language,
			Url:      cleanUrl,
			Host:     r.Host,
		},
		Items: items,
	}, w)
}
