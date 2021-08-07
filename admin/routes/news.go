package routes

import (
	"database/sql"
	"encoding/json"
	"github.com/lib/pq"
	"io/ioutil"
	"net/http"
	"os"
	"ruehmkorf.com/database/models"
	"time"
)

func NewsAction(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		newsList(w)
	} else if r.Method == http.MethodPost {
		newsNew(w, r)
	} else if r.Method == http.MethodPut {
		newsEdit(w, r)
	} else if r.Method == http.MethodDelete {
		newsDelete(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func newsList(w http.ResponseWriter) {
	type newsItem struct {
		Id        string       `json:"id"`
		TitleDe   string       `json:"titleDe"`
		TitleEn   string       `json:"titleEn"`
		Slug      string       `json:"slug"`
		Date      time.Time    `json:"date"`
		Tags      []models.Tag `json:"tags"`
		Public    bool         `json:"public"`
		GistDe    string       `json:"gistDe"`
		GistEn    string       `json:"gistEn"`
		ContentDe string       `json:"contentDe"`
		ContentEn string       `json:"contentEn"`
	}

	news, err := models.FindAllNews()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if news == nil {
		w.Write([]byte("[]"))
		return
	}

	var entries []newsItem
	for _, item := range news {
		elem := newsItem{
			Id:      item.Id,
			TitleDe: item.TitleDe,
			TitleEn: item.TitleEn,
			Slug:    item.Slug,
			Date:    item.Date,
			Public:  item.Public,
		}

		tags, err := models.FindTagsByNews(item)
		if err == nil {
			elem.Tags = tags
		}
		if item.GistEn.Valid {
			elem.GistEn = item.GistEn.String
		}
		if item.GistDe.Valid {
			elem.GistDe = item.GistDe.String
		}
		if item.ContentDe.Valid {
			elem.ContentDe = item.ContentDe.String
		}
		if item.ContentEn.Valid {
			elem.ContentEn = item.ContentEn.String
		}

		entries = append(entries, elem)
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(&entries)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func newsNew(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		newsEntry := newsData{}

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&newsEntry)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		parsedDate, err := time.Parse("2006-01-02", newsEntry.Date)
		if err != nil {
			parsedDate = time.Now()
		}

		news, err := models.CreateNews(models.News{
			TitleDe:   newsEntry.TitleDe,
			TitleEn:   newsEntry.TitleEn,
			Slug:      newsEntry.Slug,
			Public:    newsEntry.Public,
			GistDe:    sql.NullString{Valid: true, String: newsEntry.GistDe},
			GistEn:    sql.NullString{Valid: true, String: newsEntry.GistEn},
			ContentDe: sql.NullString{Valid: true, String: newsEntry.ContentDe},
			ContentEn: sql.NullString{Valid: true, String: newsEntry.ContentEn},
			Date:      parsedDate,
		})

		if conv, ok := err.(*pq.Error); ok == true && conv.Code == "23505" {
			w.WriteHeader(http.StatusConflict)
			return
		}

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		resultingTags, err := models.CreateTags(newsEntry.Tags)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = models.SetNewsTags(news.Id, resultingTags)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func newsEdit(w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Query().Get("slug")
	news, err := models.FindNewsBySlug(slug)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var bodyData newsData
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&bodyData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	parsedDate, err := time.Parse("2006-01-02", bodyData.Date)
	if err != nil {
		parsedDate = time.Now()
	}

	updatedNewsEntry := models.News{
		TitleDe:   bodyData.TitleDe,
		TitleEn:   bodyData.TitleEn,
		Slug:      slug,
		Date:      parsedDate,
		Public:    bodyData.Public,
		GistDe:    sql.NullString{String: bodyData.GistDe, Valid: true},
		GistEn:    sql.NullString{String: bodyData.GistEn, Valid: true},
		ContentDe: sql.NullString{String: bodyData.ContentDe, Valid: true},
		ContentEn: sql.NullString{String: bodyData.ContentEn, Valid: true},
	}
	updatedNewsEntry.Slug = news.Slug
	updatedNewsEntry.Id = news.Id
	_, err = models.UpdateNews(updatedNewsEntry)
	if err != nil {
		if err, ok := err.(pq.Error); ok && err.Code == "23505" {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	resultingTags, err := models.CreateTags(bodyData.Tags)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = models.SetNewsTags(news.Id, resultingTags)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func newsDelete(w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Query().Get("slug")
	err := models.DeleteNewsBySlug(slug)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	_ = os.Remove(models.HeroPath + slug)
	w.WriteHeader(http.StatusNoContent)
}

type newsData struct {
	Id        string   `json:"id"`
	TitleDe   string   `json:"titleDe"`
	TitleEn   string   `json:"titleEn"`
	Slug      string   `json:"slug"`
	Date      string   `json:"date"`
	Tags      []string `json:"tags"`
	Public    bool     `json:"public"`
	GistDe    string   `json:"gistDe"`
	GistEn    string   `json:"gistEn"`
	ContentDe string   `json:"contentDe"`
	ContentEn string   `json:"contentEn"`
}

func saveHero(slug string, data []byte) error {
	err := os.MkdirAll(models.HeroPath, 0755)
	if err != nil {
		return err
	}

	err = os.WriteFile(models.HeroPath+slug, data, 0755)

	return err
}

func UploadHero(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		slug := r.URL.Query().Get("slug")
		if slug == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = saveHero(r.URL.Query().Get("slug"), data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
