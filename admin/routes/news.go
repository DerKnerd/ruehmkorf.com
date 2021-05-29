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

func NewsList(w http.ResponseWriter, r *http.Request) {
	count, err := strconv.Atoi(r.URL.Query().Get("count"))
	if err != nil {
		count = 20
	}

	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		offset = 0
	}

	news, totalCount, err := models.FindAllNews(offset, count)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	totalPages := totalCount / count
	if totalPages == 0 {
		totalPages = 1
	}

	httpUtils.RenderAdmin("admin/templates/news/overview.gohtml", OverviewModel{
		Items:      news,
		Count:      count,
		Offset:     offset,
		NextOffset: offset + count,
		PrevOffset: offset - count,
		Page:       offset/count + 1,
		TotalPages: totalPages,
		TotalCount: totalCount,
	}, w)
}

type newsData struct {
	Message   string
	Id        string
	TitleDe   string
	TitleEn   string
	Slug      string
	Date      string
	Tags      string
	HeroImage string
	Public    bool
	GistDe    string
	GistEn    string
	ContentDe string
	ContentEn string
}

func saveHero(slug string, header *multipart.FileHeader) error {
	err := os.MkdirAll(models.HeroPath, 0755)
	if err != nil {
		return err
	}

	file, err := header.Open()
	if err != nil {
		return err
	}

	data := make([]byte, header.Size)
	_, err = file.Read(data)

	if err != nil {
		return err
	}

	err = os.WriteFile(models.HeroPath+slug, data, 0755)

	return err
}

func NewsNew(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		httpUtils.RenderAdmin("admin/templates/news/new.gohtml", newsData{Public: true}, w)
	} else if r.Method == http.MethodPost {
		err := r.ParseMultipartForm(8192 * 1024 * 1024 * 1024)
		if err != nil {
			httpUtils.RenderAdmin("admin/templates/news/new.gohtml", newsData{Message: err.Error()}, w)
			return
		}

		titleDe := r.FormValue("titleDe")
		titleEn := r.FormValue("titleEn")
		slug := r.FormValue("slug")
		date := r.FormValue("date")
		tags := r.FormValue("tags")
		public := r.FormValue("public") == "on"
		gistDe := r.FormValue("gistDe")
		gistEn := r.FormValue("gistEn")
		contentDe := r.FormValue("contentDe")
		contentEn := r.FormValue("contentEn")

		returnError := func(err error) {
			httpUtils.RenderAdmin("admin/templates/news/new.gohtml", newsData{
				Message:   err.Error(),
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

		_, header, err := r.FormFile("heroImage")
		if err != nil && err != http.ErrMissingFile {
			httpUtils.RenderAdmin("admin/templates/news/new.gohtml", newsData{
				Message:   "Hero image konnte nicht geladen werden",
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
			return
		}

		parsedDate, err := time.Parse("2006-01-02", date)
		if err != nil {
			parsedDate = time.Now()
		}

		newsEntry := models.News{
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
		news, err := models.CreateNews(newsEntry)
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

		if header != nil {
			err = saveHero(slug, header)
			if err != nil {
				returnError(err)
				return
			}
		}

		http.Redirect(w, r, "/admin/news", http.StatusFound)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func NewsEdit(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		slug := r.URL.Query().Get("slug")
		news, err := models.FindNewsBySlug(slug)
		if err != nil {
			httpUtils.RenderAdmin("admin/templates/news/edit.gohtml", newsData{
				Message: "Nachricht nicht gefunden",
			}, w)
			return
		}

		var tags []string
		for _, tag := range news.Tags {
			tags = append(tags, tag.Tag)
		}

		httpUtils.RenderAdmin("admin/templates/news/edit.gohtml", newsData{
			Message:   "",
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
				Message:   err.Error(),
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
				Message:   err.Error(),
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

		_, header, err := r.FormFile("heroImage")
		if err == nil {
			err = saveHero(slug, header)
			if err != nil {
				returnError(err)
				return
			}
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
			httpUtils.RenderAdmin("admin/templates/news/delete.gohtml", newsData{
				Message: "Nachricht nicht gefunden",
			}, w)
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
			httpUtils.RenderAdmin("admin/templates/news/delete.gohtml", newsData{
				Message: "Nachricht nicht gefunden",
			}, w)
			return
		}

		err = os.Remove(models.HeroPath + slug)
		if err != nil {
			httpUtils.RenderAdmin("admin/templates/news/delete.gohtml", newsData{
				Message: "Nachricht gelöscht, Hero image nicht",
			}, w)
			return
		}

		http.Redirect(w, r, "/admin/news", http.StatusFound)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
