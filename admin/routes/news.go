package routes

import (
	"database/sql"
	"encoding/json"
	"github.com/lib/pq"
	"io/ioutil"
	"net/http"
	"os"
	"ruehmkorf.com/database/models"
	httpUtils "ruehmkorf.com/utils/http"
	"strings"
	"time"
)

func NewsAction(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
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
				Tags:    item.Tags,
				Public:  item.Public,
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
	} else if r.Method == http.MethodPost {
		NewsNew(w, r)
	}
}

type newsData struct {
	Id        string `json:"id"`
	TitleDe   string `json:"titleDe"`
	TitleEn   string `json:"titleEn"`
	Slug      string `json:"slug"`
	Date      string `json:"date"`
	Tags      string `json:"tags"`
	Public    bool   `json:"public"`
	GistDe    string `json:"gistDe"`
	GistEn    string `json:"gistEn"`
	ContentDe string `json:"contentDe"`
	ContentEn string `json:"contentEn"`
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

func NewsNew(w http.ResponseWriter, r *http.Request) {
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

		splitTags := strings.Split(newsEntry.Tags, ",")
		resultingTags, err := models.CreateTags(splitTags)
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

func NewsEdit(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		slug := r.URL.Query().Get("slug")
		news, err := models.FindNewsBySlug(slug)
		if err != nil {
			httpUtils.RenderAdmin("admin/templates/news/edit.gohtml", newsData{}, w)
			return
		}

		var tags []string
		for _, tag := range news.Tags {
			tags = append(tags, tag.Tag)
		}

		httpUtils.RenderAdmin("admin/templates/news/edit.gohtml", newsData{
			Id:        news.Id,
			TitleDe:   news.TitleDe,
			TitleEn:   news.TitleEn,
			Slug:      news.Slug,
			Date:      news.Date.Format("2006-01-02"),
			Tags:      strings.Join(tags, ","),
			Public:    news.Public,
			GistDe:    news.GistDe.String,
			GistEn:    news.GistEn.String,
			ContentDe: news.ContentDe.String,
			ContentEn: news.ContentEn.String,
		}, w)
	} else if r.Method == http.MethodPost {
		slug := r.URL.Query().Get("slug")
		news, err := models.FindNewsBySlug(slug)
		if err != nil {
			http.Redirect(w, r, "/admin/news/edit?slug="+slug, http.StatusFound)
			return
		}

		err = r.ParseMultipartForm(8192 * 1024 * 1024 * 1024)
		if err != nil {
			var tags []string
			for _, tag := range news.Tags {
				tags = append(tags, tag.Tag)
			}

			httpUtils.RenderAdmin("admin/templates/news/edit.gohtml", newsData{
				Id:        news.Id,
				TitleDe:   news.TitleDe,
				TitleEn:   news.TitleEn,
				Slug:      news.Slug,
				Date:      news.Date.Format("2006-01-02"),
				Tags:      strings.Join(tags, ","),
				Public:    news.Public,
				GistDe:    news.GistDe.String,
				GistEn:    news.GistEn.String,
				ContentDe: news.ContentDe.String,
				ContentEn: news.ContentEn.String,
			}, w)
			return
		}

		titleDe := r.FormValue("titleDe")
		titleEn := r.FormValue("titleEn")
		date := r.FormValue("date")
		tags := r.FormValue("tags")
		public := r.FormValue("public") == "on"
		gistDe := r.FormValue("gistDe")
		gistEn := r.FormValue("gistEn")
		contentDe := r.FormValue("contentDe")
		contentEn := r.FormValue("contentEn")

		returnError := func(err error) {
			httpUtils.RenderAdmin("admin/templates/news/edit.gohtml", newsData{
				TitleDe:   titleDe,
				TitleEn:   titleEn,
				Slug:      slug,
				Date:      date,
				Tags:      tags,
				Public:    public,
				GistDe:    gistDe,
				GistEn:    gistEn,
				ContentDe: contentDe,
				ContentEn: contentEn,
			}, w)
		}

		parsedDate, err := time.Parse("2006-01-02", date)
		if err != nil {
			parsedDate = time.Now()
		}

		updatedNewsEntry := models.News{
			TitleDe:   titleDe,
			TitleEn:   titleEn,
			Slug:      slug,
			Date:      parsedDate,
			Public:    public,
			GistDe:    sql.NullString{String: gistDe, Valid: true},
			GistEn:    sql.NullString{String: gistEn, Valid: true},
			ContentDe: sql.NullString{String: contentDe, Valid: true},
			ContentEn: sql.NullString{String: contentEn, Valid: true},
		}
		updatedNewsEntry.Slug = news.Slug
		updatedNewsEntry.Id = news.Id
		_, err = models.UpdateNews(updatedNewsEntry)
		if err != nil {
			returnError(err)
			return
		}

		splitTags := strings.Split(tags, ",")
		resultingTags, err := models.CreateTags(splitTags)
		if err != nil {
			returnError(err)
			return
		}

		err = models.SetNewsTags(news.Id, resultingTags)
		if err != nil {
			returnError(err)
			return
		}

		http.Redirect(w, r, "/admin/news", http.StatusFound)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func NewsDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		slug := r.URL.Query().Get("slug")
		news, err := models.FindNewsBySlug(slug)
		if err != nil {
			httpUtils.RenderAdmin("admin/templates/news/delete.gohtml", newsData{}, w)
			return
		}

		httpUtils.RenderAdmin("admin/templates/news/delete.gohtml", newsData{
			TitleEn: news.TitleEn,
			TitleDe: news.TitleDe,
		}, w)
	} else if r.Method == http.MethodPost {
		slug := r.URL.Query().Get("slug")
		err := models.DeleteNewsBySlug(slug)
		if err != nil {
			httpUtils.RenderAdmin("admin/templates/news/delete.gohtml", newsData{}, w)
			return
		}

		err = os.Remove(models.HeroPath + slug)
		if err != nil {
			httpUtils.RenderAdmin("admin/templates/news/delete.gohtml", newsData{}, w)
			return
		}

		http.Redirect(w, r, "/admin/news", http.StatusFound)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
