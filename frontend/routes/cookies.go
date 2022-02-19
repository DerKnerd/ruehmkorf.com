package routes

import (
	"bytes"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"net/http"
	"ruehmkorf.com/database/models"
	httpUtils "ruehmkorf.com/utils/http"
)

type CookiesData struct {
	BaseData
	Content string
}

func CookiesPage(w http.ResponseWriter, r *http.Request, language string) error {
	content := ""
	if language == "de" {
		content = models.FindSettingByKey("CookiesDe")
	} else {
		content = models.FindSettingByKey("CookiesEn")
	}
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
	)
	var contentBuffer bytes.Buffer
	_ = md.Convert([]byte(content), &contentBuffer)

	data := DataProtectionData{
		BaseData: BaseData{
			Language: language,
			Url:      "cookies",
			Host:     r.Host,
		},
		Content: contentBuffer.String(),
	}

	return httpUtils.RenderFrontend("frontend/templates/cookies/index.gohtml", data, w)
}
