package routes

import (
	"bytes"
	"fmt"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"io/ioutil"
	"net/http"
	"ruehmkorf.com/database/models"
	httpUtils "ruehmkorf.com/utils/http"
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

type downloadListItem struct {
	Name string
	Date string
	Slug string
}

type downloadListOverview struct {
	BaseData
	Items []downloadListItem
}

type downloadPageItem struct {
	BaseData
	Name        string
	Date        string
	Description string
	Slug        string
}

func DownloadList(w http.ResponseWriter, r *http.Request, language string) error {
	fileType := r.URL.Query().Get("fileType")
	download, err := models.FindAllDownloadsForFrontend(fileType)

	if err != nil {
		return err
	}

	var items []downloadListItem
	for _, item := range download {
		if language == "de" {
			items = append(items, downloadListItem{
				Name: item.NameDe,
				Date: item.Date.Format("2006-01-02"),
				Slug: item.Slug,
			})
		} else {
			items = append(items, downloadListItem{
				Name: item.NameEn,
				Date: item.Date.Format("2006-01-02"),
				Slug: item.Slug,
			})
		}
	}
	cleanUrl := ""
	if language == "de" {
		cleanUrl = strings.TrimPrefix(r.URL.Path, "/de/")
	} else {
		cleanUrl = strings.TrimPrefix(r.URL.Path, "/en/")
	}

	return httpUtils.RenderFrontend("frontend/templates/download/index.gohtml", downloadListOverview{
		BaseData: BaseData{
			Language: language,
			Url:      cleanUrl,
			Host:     r.Host,
		},
		Items: items,
	}, w)
}

func DownloadPage(w http.ResponseWriter, r *http.Request, language string) error {
	slug := ""
	cleanUrl := ""
	if language == "de" {
		slug = strings.TrimPrefix(r.URL.Path, "/de/download/")
		cleanUrl = strings.TrimPrefix(r.URL.Path, "/de/")
	} else {
		slug = strings.TrimPrefix(r.URL.Path, "/en/download/")
		cleanUrl = strings.TrimPrefix(r.URL.Path, "/en/")
	}

	err := models.CheckIfDownloadExistsBySlug(slug)
	if err != nil {
		return Error404Handler(w, r)
	}

	download, err := models.FindDownloadBySlug(slug)
	if err != nil {
		return err
	}

	result := downloadPageItem{
		BaseData: BaseData{
			Language: language,
			Url:      cleanUrl,
			Host:     r.Host,
		},
		Date: download.Date.Format("2006-01-02"),
		Slug: slug,
	}
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
	)

	var descriptionBuffer bytes.Buffer
	if language == "de" {
		result.Name = download.NameDe
		_ = md.Convert([]byte(download.DescriptionDe.String), &descriptionBuffer)
	} else {
		result.Name = download.NameEn
		_ = md.Convert([]byte(download.DescriptionEn.String), &descriptionBuffer)
	}

	result.Description = descriptionBuffer.String()

	return httpUtils.RenderFrontend("frontend/templates/download/page.gohtml", result, w)
}
