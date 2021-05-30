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
	BaseData
	Items []newsListItem
	Tags  []string
}

type newsPageItem struct {
	BaseData
	Title   string
	Date    string
	Gist    string
	Content string
	Slug    string
	Tags    []string
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
	cleanUrl := ""
	if language == "de" {
		cleanUrl = strings.TrimPrefix(r.URL.Path, "/de/")
	} else {
		cleanUrl = strings.TrimPrefix(r.URL.Path, "/en/")
	}

	tags, _ := models.FindAllTags()

	return httpUtils.RenderFrontend("frontend/templates/news/index.gohtml", newsListOverview{
		BaseData: BaseData{
			Language: language,
			Url:      cleanUrl,
		},
		Items: items,
		Tags:  tags,
	}, w)
}

func NewsPage(w http.ResponseWriter, r *http.Request, language string) error {
	slug := ""
	cleanUrl := ""
	if language == "de" {
		slug = strings.TrimPrefix(r.URL.Path, "/de/news/")
		cleanUrl = strings.TrimPrefix(r.URL.Path, "/de/")
	} else {
		slug = strings.TrimPrefix(r.URL.Path, "/en/news/")
		cleanUrl = strings.TrimPrefix(r.URL.Path, "/en/")
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
		BaseData: BaseData{
			Language: language,
			Url:      cleanUrl,
		},
		Date: news.Date.Format("2006-01-02"),
		Slug: slug,
	}
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
	)

	var contentBuffer bytes.Buffer
	var gistBuffer bytes.Buffer
	if language == "de" {
		result.Gist = news.GistDe.String
		result.Title = news.TitleDe
		_ = md.Convert([]byte(news.ContentDe.String), &contentBuffer)
		_ = md.Convert([]byte(news.GistDe.String), &gistBuffer)
	} else {
		result.Gist = news.GistEn.String
		result.Title = news.TitleEn
		_ = md.Convert([]byte(news.ContentEn.String), &contentBuffer)
		_ = md.Convert([]byte(news.GistEn.String), &gistBuffer)
	}

	result.Content = contentBuffer.String()
	result.Gist = gistBuffer.String()
	tags, _ := models.FindAllTags()
	result.Tags = tags

	return httpUtils.RenderFrontend("frontend/templates/news/page.gohtml", result, w)
}
