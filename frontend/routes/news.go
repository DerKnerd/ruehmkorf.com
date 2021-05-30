package routes

import (
	"bytes"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"io/ioutil"
	"net/http"
	"ruehmkorf.com/database/models"
	httpUtils "ruehmkorf.com/utils/http"
	"strings"
)

func NewsHeroImage(w http.ResponseWriter, r *http.Request) {
	slug := strings.TrimPrefix(r.URL.Path, "/news/hero/")
	news, err := models.FindNewsBySlug(slug)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	data, err := ioutil.ReadFile(models.HeroPath + slug)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	if !news.Public {
		w.Header().Add("X-Robots-Tag", "none")
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

type newsListItem struct {
	Title string
	Date  string
	Slug  string
}

type newsListOverview struct {
	Items    []newsListItem
	Language string
	Tags     []string
}

type newsPageItem struct {
	Title    string
	Date     string
	Gist     string
	Content  string
	Language string
	Slug     string
	Tags     []string
}

func NewsList(w http.ResponseWriter, r *http.Request, language string) error {
	topic := r.URL.Query().Get("topic")
	news, err := models.FindAllNewsForFrontend(language, topic)

	if err != nil {
		return err
	}

	var items []newsListItem
	for _, item := range news {
		if language == "de" && item.ContentDe.String != "" {
			items = append(items, newsListItem{
				Title: item.TitleDe,
				Date:  item.Date.Format("2006-01-02"),
				Slug:  item.Slug,
			})
		} else if item.ContentEn.String != "" {
			items = append(items, newsListItem{
				Title: item.TitleEn,
				Date:  item.Date.Format("2006-01-02"),
				Slug:  item.Slug,
			})
		}
	}

	tags, _ := models.FindAllTags()

	return httpUtils.RenderFrontend("frontend/templates/news/index.gohtml", newsListOverview{
		Items:    items,
		Language: language,
		Tags:     tags,
	}, w)
}

func NewsPage(w http.ResponseWriter, r *http.Request, language string) error {
	slug := ""
	if language == "de" {
		slug = strings.TrimPrefix(r.URL.Path, "/de/news/")
	} else {
		slug = strings.TrimPrefix(r.URL.Path, "/en/news/")
	}

	err := models.CheckIfNewsExistsBySlugAndLanguage(slug, language)
	if err != nil {
		return Error404Handler(w, r)
	}

	news, err := models.FindNewsBySlug(slug)
	if err != nil {
		return err
	}

	result := newsPageItem{
		Language: language,
		Date:     news.Date.Format("2006-01-02"),
		Slug:     slug,
	}
	content := ""
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
	)
	if language == "de" {
		result.Gist = news.GistDe.String
		result.Title = news.TitleDe
		var buffer bytes.Buffer
		_ = md.Convert([]byte(news.ContentDe.String), &buffer)
		content = buffer.String()
	} else {
		result.Gist = news.GistEn.String
		result.Title = news.TitleEn
		var buffer bytes.Buffer
		_ = md.Convert([]byte(news.ContentEn.String), &buffer)
		content = buffer.String()
	}

	result.Content = content
	tags, _ := models.FindAllTags()
	result.Tags = tags

	return httpUtils.RenderFrontend("frontend/templates/news/page.gohtml", result, w)
}
