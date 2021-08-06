package routes

import (
	"encoding/json"
	"net/http"
	"ruehmkorf.com/database/models"
)

type buchstabieromatData struct {
	DescriptionDe string `json:"descriptionDe"`
	InfoTextDeDe  string `json:"infoTextDeDe"`
	InfoTextEnDe  string `json:"infoTextEnDe"`
	DescriptionEn string `json:"descriptionEn"`
	InfoTextDeEn  string `json:"infoTextDeEn"`
	InfoTextEnEn  string `json:"infoTextEnEn"`
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
		encoder := json.NewEncoder(w)
		encoder.SetEscapeHTML(true)
		err := encoder.Encode(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else if r.Method == http.MethodPut {
		decoder := json.NewDecoder(r.Body)
		var data buchstabieromatData
		err := decoder.Decode(&data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		models.UpdateSetting("DescriptionDe", data.DescriptionDe)
		models.UpdateSetting("InfoTextDeDe", data.InfoTextDeDe)
		models.UpdateSetting("InfoTextEnDe", data.InfoTextEnDe)
		models.UpdateSetting("DescriptionEn", data.DescriptionEn)
		models.UpdateSetting("InfoTextDeEn", data.InfoTextDeEn)
		models.UpdateSetting("InfoTextEnEn", data.InfoTextEnEn)
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
