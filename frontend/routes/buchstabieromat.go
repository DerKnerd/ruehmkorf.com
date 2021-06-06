package routes

import (
	"bytes"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"net/http"
	"ruehmkorf.com/database/models"
	httpUtils "ruehmkorf.com/utils/http"
)

type BuchstabieroMatData struct {
	BaseData
	Description string
	InfoTextDe  string
	InfoTextEn  string
}

func BuchstabieroMatPage(w http.ResponseWriter, r *http.Request, language string) error {
	description := ""
	infoTextDe := ""
	infoTextEn := ""
	if language == "de" {
		description = models.FindSettingByKey("DescriptionDe")
		infoTextDe = models.FindSettingByKey("InfoTextDeDe")
		infoTextEn = models.FindSettingByKey("InfoTextEnDe")
	} else {
		description = models.FindSettingByKey("DescriptionEn")
		infoTextDe = models.FindSettingByKey("InfoTextDeEn")
		infoTextEn = models.FindSettingByKey("InfoTextEnEn")
	}
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
	)
	var descriptionBuffer bytes.Buffer
	var infoTextDeBuffer bytes.Buffer
	var infoTextEnBuffer bytes.Buffer
	_ = md.Convert([]byte(description), &descriptionBuffer)
	_ = md.Convert([]byte(infoTextDe), &infoTextDeBuffer)
	_ = md.Convert([]byte(infoTextEn), &infoTextEnBuffer)

	data := BuchstabieroMatData{
		BaseData: BaseData{
			Language: language,
			Url:      "buchstabier-o-mat",
			Host:     r.Host,
		},
		Description: descriptionBuffer.String(),
		InfoTextDe:  infoTextDeBuffer.String(),
		InfoTextEn:  infoTextEnBuffer.String(),
	}

	return httpUtils.RenderFrontend("frontend/templates/buchstabieromat/index.gohtml", data, w)
}
