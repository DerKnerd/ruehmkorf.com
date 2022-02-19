package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/lib/pq"
	"io/ioutil"
	"net/http"
	"os"
	"ruehmkorf.com/database/models"
	"strings"
	"time"
)

type downloadItem struct {
	Id               string    `json:"id"`
	NameDe           string    `json:"nameDe"`
	NameEn           string    `json:"nameEn"`
	Slug             string    `json:"slug"`
	Date             time.Time `json:"date"`
	SelfDestruct     bool      `json:"selfDestruct"`
	SelfDestructDays int       `json:"selfDestructDays"`
	Public           bool      `json:"public"`
	DescriptionDe    string    `json:"descriptionDe"`
	DescriptionEn    string    `json:"descriptionEn"`
	Type             string    `json:"type"`
	FileExtension    string    `json:"fileExtension"`
}

func DownloadAction(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		slug := r.URL.Query().Get("slug")
		if slug != "" {
			downloadDetails(w, slug)
		} else {
			downloadList(w)
		}
	} else if r.Method == http.MethodPost {
		downloadNew(w, r)
	} else if r.Method == http.MethodPut {
		downloadEdit(w, r)
	} else if r.Method == http.MethodDelete {
		downloadDelete(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func downloadItemFromDownload(download models.Download) downloadItem {
	return downloadItem{
		Id:               download.Id,
		NameDe:           download.NameDe,
		NameEn:           download.NameEn,
		Slug:             download.Slug,
		Date:             download.Date,
		SelfDestruct:     download.SelfDestruct,
		SelfDestructDays: int(download.SelfDestructDays.Int32),
		Public:           download.Public,
		DescriptionDe:    download.DescriptionDe.String,
		DescriptionEn:    download.DescriptionEn.String,
		Type:             download.Type,
		FileExtension:    download.FileExtension.String,
	}
}

func downloadList(w http.ResponseWriter) {
	downloads, err := models.FindAllDownloads()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	entries := make([]downloadItem, 0)
	for _, item := range downloads {
		entries = append(entries, downloadItemFromDownload(item))
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(entries)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func downloadDetails(w http.ResponseWriter, slug string) {
	download, err := models.FindDownloadBySlug(slug)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(downloadItemFromDownload(*download))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type downloadData struct {
	NameDe           string `json:"nameDe"`
	NameEn           string `json:"nameEn"`
	Slug             string `json:"slug"`
	Public           bool   `json:"public"`
	SelfDestructDays int    `json:"selfDestructDays"`
	DescriptionDe    string `json:"descriptionDe"`
	DescriptionEn    string `json:"descriptionEn"`
	Date             string `json:"date"`
	Extension        string `json:"extension"`
}

func UploadPreviewAction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		slug := r.URL.Query().Get("slug")
		download, err := models.FindDownloadBySlug(slug)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		path := models.DownloadPreviewImagePath

		err = os.MkdirAll(path, 0755)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = os.WriteFile(path+slug, data, 0755)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		mimeType := http.DetectContentType(data)
		split := strings.Split(mimeType, "/")

		download.Type = split[0]
		w.WriteHeader(http.StatusNoContent)
	}
}

func downloadNew(w http.ResponseWriter, r *http.Request) {
	var data downloadData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	parsedDate, err := time.Parse("2006-01-02T15:04:05Z", data.Date)
	if err != nil {
		parsedDate = time.Now()
	}

	var selfDestruct bool
	if data.SelfDestructDays == 0 {
		selfDestruct = false
	} else if data.SelfDestructDays > 0 {
		selfDestruct = true
	}

	download := models.Download{
		NameDe:       data.NameDe,
		NameEn:       data.NameEn,
		Slug:         data.Slug,
		Date:         parsedDate,
		SelfDestruct: selfDestruct,
		SelfDestructDays: sql.NullInt32{
			Int32: int32(data.SelfDestructDays),
			Valid: true,
		},
		Public: data.Public,
		DescriptionDe: sql.NullString{
			String: data.DescriptionDe,
			Valid:  true,
		},
		DescriptionEn: sql.NullString{
			String: data.DescriptionEn,
			Valid:  true,
		},
		Type: "",
		FileExtension: sql.NullString{
			String: data.Extension,
			Valid:  true,
		},
	}

	err = models.CreateDownload(download)

	if conv, ok := err.(*pq.Error); ok == true && conv.Code == "23505" {
		w.WriteHeader(http.StatusConflict)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func downloadEdit(w http.ResponseWriter, r *http.Request) {
	download, err := models.FindDownloadBySlug(r.URL.Query().Get("slug"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var data downloadData
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	parsedDate, err := time.Parse("2006-01-02T15:04:05Z", data.Date)
	if err != nil {
		parsedDate = time.Now()
	}

	var selfDestruct bool
	if data.SelfDestructDays == 0 {
		selfDestruct = false
	} else if data.SelfDestructDays > 0 {
		selfDestruct = true
	}

	download.NameDe = data.NameDe
	download.NameEn = data.NameEn
	download.Slug = data.Slug
	download.Date = parsedDate
	download.SelfDestruct = selfDestruct
	download.SelfDestructDays = sql.NullInt32{
		Int32: int32(data.SelfDestructDays),
		Valid: true,
	}
	download.Public = data.Public
	download.DescriptionDe = sql.NullString{
		String: data.DescriptionDe,
		Valid:  true,
	}
	download.DescriptionEn = sql.NullString{
		String: data.DescriptionEn,
		Valid:  true,
	}
	download.FileExtension = sql.NullString{
		String: data.Extension,
		Valid:  true,
	}

	err = models.UpdateDownload(*download)

	if conv, ok := err.(*pq.Error); ok == true && conv.Code == "23505" {
		w.WriteHeader(http.StatusConflict)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func downloadDelete(w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Query().Get("slug")
	err := models.DeleteDownloadBySlug(slug)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = os.Remove(models.DownloadFilePath + slug)
	err = os.Remove(models.DownloadPreviewImagePath + slug)

	w.WriteHeader(http.StatusNoContent)
}

var chunksTempDir = "./data/chunks"

func UploadFileAction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		slug := r.URL.Query().Get("slug")
		download, err := models.FindDownloadBySlug(slug)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		chunksDir := fmt.Sprintf("%s/%s", chunksTempDir, download.Id)

		err = os.MkdirAll(chunksDir, 0755)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if strings.HasSuffix(r.URL.Path, "chunk") {
			err = uploadSingleChunk(r, chunksDir)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		} else if strings.HasSuffix(r.URL.Path, "finish") {
			err = concatChunksToFile(chunksDir, download)
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func concatChunksToFile(chunksDir string, download *models.Download) error {
	chunks, err := os.ReadDir(chunksDir)
	if err != nil {
		return err
	}

	downloadFile, err := os.OpenFile(models.DownloadFilePath+"/"+download.Slug, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}

	defer downloadFile.Close()

	for _, chunk := range chunks {
		if chunk.IsDir() {
			continue
		}

		data, err := ioutil.ReadFile(chunksDir + "/" + chunk.Name())
		if err != nil {
			return err
		}

		_, err = downloadFile.Write(data)
		if err != nil {
			return err
		}
	}

	_ = os.RemoveAll(chunksDir)

	return nil
}

func uploadSingleChunk(r *http.Request, tmpPath string) error {
	idx := r.URL.Query().Get("index")
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(fmt.Sprintf("%s/%s", tmpPath, idx), data, 0755)
}
