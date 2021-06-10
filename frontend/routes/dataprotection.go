package routes

import (
	"bytes"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"net/http"
	"ruehmkorf.com/database/models"
	httpUtils "ruehmkorf.com/utils/http"
)

type DataProtectionData struct {
	BaseData
	Content string
}

func DataProtectionPage(w http.ResponseWriter, r *http.Request, language string) error {
	content := ""
	if language == "de" {
		content = models.FindSettingByKey("DataProtectionDe")
	} else {
		content = models.FindSettingByKey("DataProtectionEn")
	}
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
	)
	var contentBuffer bytes.Buffer
	_ = md.Convert([]byte(content), &contentBuffer)

	data := DataProtectionData{
		BaseData: BaseData{
			Language: language,
			Url:      "data-protection",
			Host:     r.Host,
		},
		Content: contentBuffer.String(),
	}

	return httpUtils.RenderFrontend("frontend/templates/dataprotection/index.gohtml", data, w)
}
