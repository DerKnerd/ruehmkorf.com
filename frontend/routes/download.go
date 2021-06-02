package routes

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"ruehmkorf.com/database/models"
	"strings"
)

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	slug := strings.TrimPrefix(r.URL.Path, "/download/file/")
	download, err := models.FindDownloadBySlug(slug)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	data, err := ioutil.ReadFile(models.DownloadFilePath + slug)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	if !download.Public {
		w.Header().Add("X-Robots-Tag", "none")
	}

	w.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s%s\"", download.NameEn, download.FileExtension.String))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func DownloadPreview(w http.ResponseWriter, r *http.Request) {
	slug := strings.TrimPrefix(r.URL.Path, "/download/preview/")
	download, err := models.FindDownloadBySlug(slug)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	data, err := ioutil.ReadFile(models.DownloadPreviewImagePath + slug)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	if !download.Public {
		w.Header().Add("X-Robots-Tag", "none")
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
