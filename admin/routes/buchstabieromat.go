package routes

import (
	"net/http"
	"ruehmkorf.com/database/models"
	httpUtils "ruehmkorf.com/utils/http"
)

type buchstabieromatData struct {
	Message       string
	DescriptionDe string
	InfoTextDeDe  string
	InfoTextEnDe  string
	DescriptionEn string
	InfoTextDeEn  string
	InfoTextEnEn  string
}

func Buchstabieromat(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data := buchstabieromatData{
			DescriptionDe: models.FindSettingByKey("DescriptionDe"),
			InfoTextDeDe:  models.FindSettingByKey("InfoTextDeDe"),
			InfoTextEnDe:  models.FindSettingByKey("InfoTextEnDe"),
			DescriptionEn: models.FindSettingByKey("DescriptionEn"),
			InfoTextDeEn:  models.FindSettingByKey("InfoTextDeEn"),
			InfoTextEnEn:  models.FindSettingByKey("InfoTextEnEn"),
		}
		httpUtils.RenderAdmin("admin/templates/buchstabieromat/index.gohtml", data, w)
	} else if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			httpUtils.RenderAdmin("admin/templates/buchstabieromat/index.gohtml", buchstabieromatData{
				Message: err.Error(),
			}, w)
			return
		}

		descriptionDe := r.FormValue("descriptionDe")
		infoTextDeDe := r.FormValue("infoTextDeDe")
		infoTextEnDe := r.FormValue("infoTextEnDe")
		descriptionEn := r.FormValue("descriptionEn")
		infoTextDeEn := r.FormValue("infoTextDeEn")
		infoTextEnEn := r.FormValue("infoTextEnEn")

		models.UpdateSetting("DescriptionDe", descriptionDe)
		models.UpdateSetting("InfoTextDeDe", infoTextDeDe)
		models.UpdateSetting("InfoTextEnDe", infoTextEnDe)
		models.UpdateSetting("DescriptionEn", descriptionEn)
		models.UpdateSetting("InfoTextDeEn", infoTextDeEn)
		models.UpdateSetting("InfoTextEnEn", infoTextEnEn)
		http.Redirect(w, r, "/admin/buchstabieromat", http.StatusFound)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
