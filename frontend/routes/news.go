package routes

import (
	"io/ioutil"
	"net/http"
	"ruehmkorf.com/database/models"
	"strings"
)

func HeroImage(w http.ResponseWriter, r *http.Request) {
	slug := strings.TrimPrefix(r.URL.Path, "/news/hero/")
	_, err := models.FindNewsBySlug(slug)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	data, err := ioutil.ReadFile(models.HeroPath + slug)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
