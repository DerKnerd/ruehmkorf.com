package routes

import (
	"database/sql"
	"mime/multipart"
	"net/http"
	"os"
	"ruehmkorf.com/database/models"
	httpUtils "ruehmkorf.com/utils/http"
	"strconv"
	"strings"
	"time"
)

func DownloadList(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		count, err := strconv.Atoi(r.URL.Query().Get("count"))
		if err != nil {
			count = 20
		}

		offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
		if err != nil {
			offset = 0
		}

		downloads, totalCount, err := models.FindAllDownloads(offset, count)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		totalPages := totalCount / count
		if totalPages == 0 {
			totalPages = 1
		}

		httpUtils.RenderAdmin("admin/templates/download/overview.gohtml", OverviewModel{
			Items:      downloads,
			Count:      count,
			Offset:     offset,
			NextOffset: offset + count,
			PrevOffset: offset - count,
			Page:       offset/count + 1,
			TotalPages: totalPages,
			TotalCount: totalCount,
		}, w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

type downloadData struct {
	Message          string
	NameDe           string
	NameEn           string
	Slug             string
	Public           bool
	SelfDestructDays string
	DescriptionDe    string
	DescriptionEn    string
	Date             string
}

func saveDownloadFile(slug string, header *multipart.FileHeader, preview bool) (string, string, error) {
	path := models.DownloadFilePath
	if preview {
		path = models.DownloadPreviewImagePath
	}

	err := os.MkdirAll(path, 0755)
	if err != nil {
		return "", "", err
	}

	file, err := header.Open()
	if err != nil {
		return "", "", err
	}

	data := make([]byte, header.Size)
	_, err = file.Read(data)
	if err != nil {
		return "", "", err
	}

	err = os.WriteFile(path+slug, data, 0755)

	mimeType := http.DetectContentType(data)
	split := strings.Split(mimeType, "/")

	return path + slug, split[0], err
}

func DownloadNew(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		httpUtils.RenderAdmin("admin/templates/download/new.gohtml", downloadData{
			Public: true,
		}, w)
	} else if r.Method == http.MethodPost {
		err := r.ParseMultipartForm(8192 * 1024 * 1024 * 1024)
		if err != nil {
			httpUtils.RenderAdmin("admin/templates/download/new.gohtml", downloadData{Message: err.Error()}, w)
			return
		}

		slug := r.FormValue("slug")
		nameDe := r.FormValue("nameDe")
		nameEn := r.FormValue("nameEn")
		date := r.FormValue("date")
		selfDestructDays := r.FormValue("selfDestructDays")
		public := r.FormValue("public") == "on"
		descriptionDe := r.FormValue("descriptionDe")
		descriptionEn := r.FormValue("descriptionEn")

		_, downloadFileHeader, err := r.FormFile("file")
		if err != nil {
			httpUtils.RenderAdmin("admin/templates/download/new.gohtml", downloadData{
				Message:          "Datei konnte nicht geladen werden",
				Slug:             slug,
				NameDe:           nameDe,
				NameEn:           nameEn,
				Date:             date,
				SelfDestructDays: selfDestructDays,
				Public:           public,
				DescriptionDe:    descriptionDe,
				DescriptionEn:    descriptionEn,
			}, w)
			return
		}

		downloadFilePath, fileType, err := saveDownloadFile(slug, downloadFileHeader, false)
		if err != nil {
			httpUtils.RenderAdmin("admin/templates/download/new.gohtml", downloadData{
				Message:          "Datei konnte nicht geladen werden",
				Slug:             slug,
				NameDe:           nameDe,
				NameEn:           nameEn,
				Date:             date,
				SelfDestructDays: selfDestructDays,
				Public:           public,
				DescriptionDe:    descriptionDe,
				DescriptionEn:    descriptionEn,
			}, w)
			return
		}

		_, previewImageHeader, err := r.FormFile("previewImage")
		previewPath := ""
		if err != nil && err != http.ErrMissingFile {
			httpUtils.RenderAdmin("admin/templates/download/new.gohtml", downloadData{
				Message:          "Vorschau konnte nicht geladen werden",
				Slug:             slug,
				NameDe:           nameDe,
				NameEn:           nameEn,
				Date:             date,
				SelfDestructDays: selfDestructDays,
				Public:           public,
				DescriptionDe:    descriptionDe,
				DescriptionEn:    descriptionEn,
			}, w)
			return
		}

		if previewImageHeader != nil {
			previewPath, _, err = saveDownloadFile(slug, previewImageHeader, true)
			if err != nil {
				httpUtils.RenderAdmin("admin/templates/download/new.gohtml", downloadData{
					Message:          "Vorschau konnte nicht geladen werden",
					Slug:             slug,
					NameDe:           nameDe,
					NameEn:           nameEn,
					Date:             date,
					SelfDestructDays: selfDestructDays,
					Public:           public,
					DescriptionDe:    descriptionDe,
					DescriptionEn:    descriptionEn,
				}, w)
				return
			}
		}

		parsedDate, err := time.Parse("2006-01-02", date)
		if err != nil {
			parsedDate = time.Now()
		}

		var selfDestruct bool
		parsedSelfDestructDays, err := strconv.Atoi(selfDestructDays)
		if err != nil {
			selfDestruct = false
		} else if parsedSelfDestructDays > 0 {
			selfDestruct = true
		}

		lastDotIndex := strings.LastIndex(downloadFileHeader.Filename, ".")
		download := models.Download{
			NameDe: nameDe,
			NameEn: nameEn,
			Slug:   slug,
			Date:   parsedDate,
			FileExtension: sql.NullString{
				String: downloadFileHeader.Filename[0:lastDotIndex],
				Valid:  true,
			},
			SelfDestruct: selfDestruct,
			SelfDestructDays: sql.NullInt32{
				Int32: int32(parsedSelfDestructDays),
				Valid: true,
			},
			Public: public,
			DescriptionDe: sql.NullString{
				String: descriptionDe,
				Valid:  true,
			},
			DescriptionEn: sql.NullString{
				String: descriptionEn,
				Valid:  true,
			},
			Type: fileType,
		}

		err = models.CreateDownload(download)
		if err != nil {
			httpUtils.RenderAdmin("admin/templates/download/new.gohtml", downloadData{
				Message:          err.Error(),
				Slug:             slug,
				NameDe:           nameDe,
				NameEn:           nameEn,
				Date:             date,
				SelfDestructDays: selfDestructDays,
				Public:           public,
				DescriptionDe:    descriptionDe,
				DescriptionEn:    descriptionEn,
			}, w)
			_ = os.Remove(previewPath)
			_ = os.Remove(downloadFilePath)
			return
		}

		http.Redirect(w, r, "/admin/download", http.StatusFound)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func DownloadEdit(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		download, err := models.FindDownloadBySlug(r.URL.Query().Get("slug"))
		if err != nil {
			httpUtils.RenderAdmin("admin/templates/download/edit.gohtml", downloadData{
				Message: "Nachricht nicht gefunden",
			}, w)
			return
		}

		httpUtils.RenderAdmin("admin/templates/download/edit.gohtml", downloadData{
			NameDe:           download.NameDe,
			NameEn:           download.NameEn,
			Slug:             download.Slug,
			Public:           download.Public,
			SelfDestructDays: strconv.Itoa(int(download.SelfDestructDays.Int32)),
			DescriptionDe:    download.DescriptionDe.String,
			DescriptionEn:    download.DescriptionEn.String,
			Date:             download.Date.Format("2006-01-02"),
		}, w)
	} else if r.Method == http.MethodPost {
		err := r.ParseMultipartForm(8192 * 1024 * 1024 * 1024)
		if err != nil {
			httpUtils.RenderAdmin("admin/templates/download/edit.gohtml", downloadData{Message: err.Error()}, w)
			return
		}

		download, err := models.FindDownloadBySlug(r.URL.Query().Get("slug"))
		if err != nil {
			httpUtils.RenderAdmin("admin/templates/download/edit.gohtml", downloadData{
				Message: "Nachricht nicht gefunden",
			}, w)
			return
		}

		slug := r.FormValue("slug")
		nameDe := r.FormValue("nameDe")
		nameEn := r.FormValue("nameEn")
		date := r.FormValue("date")
		selfDestructDays := r.FormValue("selfDestructDays")
		public := r.FormValue("public") == "on"
		descriptionDe := r.FormValue("descriptionDe")
		descriptionEn := r.FormValue("descriptionEn")

		_, downloadFileHeader, err := r.FormFile("file")
		downloadFilePath := ""
		if err != nil && err != http.ErrMissingFile {
			httpUtils.RenderAdmin("admin/templates/download/edit.gohtml", downloadData{
				Message:          "Vorschau konnte nicht geladen werden",
				Slug:             slug,
				NameDe:           nameDe,
				NameEn:           nameEn,
				Date:             date,
				SelfDestructDays: selfDestructDays,
				Public:           public,
				DescriptionDe:    descriptionDe,
				DescriptionEn:    descriptionEn,
			}, w)
			return
		}

		fileType := download.Type
		if downloadFileHeader != nil {
			_, fileType, err = saveDownloadFile(slug, downloadFileHeader, false)
			if err != nil {
				httpUtils.RenderAdmin("admin/templates/download/edit.gohtml", downloadData{
					Message:          "Vorschau konnte nicht geladen werden",
					Slug:             slug,
					NameDe:           nameDe,
					NameEn:           nameEn,
					Date:             date,
					SelfDestructDays: selfDestructDays,
					Public:           public,
					DescriptionDe:    descriptionDe,
					DescriptionEn:    descriptionEn,
				}, w)
				return
			}

			lastDotIndex := strings.LastIndex(downloadFileHeader.Filename, ".")
			download.FileExtension = sql.NullString{
				String: downloadFileHeader.Filename[lastDotIndex:len(downloadFileHeader.Filename)],
				Valid:  true,
			}
		}

		_, previewImageHeader, err := r.FormFile("previewImage")
		previewPath := ""
		if err != nil && err != http.ErrMissingFile {
			httpUtils.RenderAdmin("admin/templates/download/edit.gohtml", downloadData{
				Message:          "Vorschau konnte nicht geladen werden",
				Slug:             slug,
				NameDe:           nameDe,
				NameEn:           nameEn,
				Date:             date,
				SelfDestructDays: selfDestructDays,
				Public:           public,
				DescriptionDe:    descriptionDe,
				DescriptionEn:    descriptionEn,
			}, w)
			return
		}

		if previewImageHeader != nil {
			previewPath, _, err = saveDownloadFile(slug, previewImageHeader, true)
			if err != nil {
				httpUtils.RenderAdmin("admin/templates/download/edit.gohtml", downloadData{
					Message:          "Vorschau konnte nicht geladen werden",
					Slug:             slug,
					NameDe:           nameDe,
					NameEn:           nameEn,
					Date:             date,
					SelfDestructDays: selfDestructDays,
					Public:           public,
					DescriptionDe:    descriptionDe,
					DescriptionEn:    descriptionEn,
				}, w)
				return
			}
		}

		parsedDate, err := time.Parse("2006-01-02", date)
		if err != nil {
			parsedDate = time.Now()
		}

		var selfDestruct bool
		parsedSelfDestructDays, err := strconv.Atoi(selfDestructDays)
		if err != nil {
			selfDestruct = false
		} else if parsedSelfDestructDays > 0 {
			selfDestruct = true
		}

		download.NameDe = nameDe
		download.NameEn = nameEn
		download.Slug = slug
		download.Date = parsedDate
		download.SelfDestruct = selfDestruct
		download.SelfDestructDays = sql.NullInt32{
			Int32: int32(parsedSelfDestructDays),
			Valid: true,
		}
		download.Public = public
		download.DescriptionDe = sql.NullString{
			String: descriptionDe,
			Valid:  true,
		}
		download.DescriptionEn = sql.NullString{
			String: descriptionEn,
			Valid:  true,
		}
		download.Type = fileType

		err = models.UpdateDownload(*download)
		if err != nil {
			httpUtils.RenderAdmin("admin/templates/download/edit.gohtml", downloadData{
				Message:          "Download konnte nicht gespeichert werden",
				Slug:             slug,
				NameDe:           nameDe,
				NameEn:           nameEn,
				Date:             date,
				SelfDestructDays: selfDestructDays,
				Public:           public,
				DescriptionDe:    descriptionDe,
				DescriptionEn:    descriptionEn,
			}, w)
			_ = os.Remove(previewPath)
			_ = os.Remove(downloadFilePath)
			return
		}

		http.Redirect(w, r, "/admin/download", http.StatusFound)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func DownloadDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		slug := r.URL.Query().Get("slug")
		news, err := models.FindDownloadBySlug(slug)
		if err != nil {
			httpUtils.RenderAdmin("admin/templates/download/delete.gohtml", downloadData{
				Message: "Download nicht gefunden",
			}, w)
			return
		}

		httpUtils.RenderAdmin("admin/templates/download/delete.gohtml", downloadData{
			NameEn: news.NameEn,
			NameDe: news.NameDe,
		}, w)
	} else if r.Method == http.MethodPost {
		slug := r.URL.Query().Get("slug")
		err := models.DeleteDownloadBySlug(slug)
		if err != nil {
			httpUtils.RenderAdmin("admin/templates/download/delete.gohtml", downloadData{
				Message: "Download nicht gefunden",
			}, w)
			return
		}

		err = os.Remove(models.DownloadFilePath + slug)
		err = os.Remove(models.DownloadPreviewImagePath + slug)

		http.Redirect(w, r, "/admin/download", http.StatusFound)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
