package routes

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"ruehmkorf.com/database/models"
	"strings"
)

type settingsData struct {
	CookiesDe        string `json:"cookiesDe"`
	CookiesEn        string `json:"cookiesEn"`
	DataProtectionDe string `json:"dataProtectionDe"`
	DataProtectionEn string `json:"dataProtectionEn"`
}

var SettingsPath = os.Getenv("DATA_DIR") + "/public/settings/"

func saveData(reader io.ReadCloser, fileName string) error {
	path := SettingsPath
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	err = os.WriteFile(path+fileName, data, 0755)

	return err
}

func SettingsAction(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		getSettings(w)
	} else if r.Method == http.MethodPut {
		putSettings(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func putSettings(w http.ResponseWriter, r *http.Request) {
	var data settingsData

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	models.UpdateSetting("CookiesDe", data.CookiesDe)
	models.UpdateSetting("CookiesEn", data.CookiesEn)
	models.UpdateSetting("DataProtectionDe", data.DataProtectionDe)
	models.UpdateSetting("DataProtectionEn", data.DataProtectionEn)

	w.WriteHeader(http.StatusNoContent)
}

func getSettings(w http.ResponseWriter) {
	data := settingsData{
		CookiesDe:        models.FindSettingByKey("CookiesDe"),
		CookiesEn:        models.FindSettingByKey("CookiesEn"),
		DataProtectionDe: models.FindSettingByKey("DataProtectionDe"),
		DataProtectionEn: models.FindSettingByKey("DataProtectionEn"),
	}

	encoder := json.NewEncoder(w)
	err := encoder.Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func LogosAction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		fileName := ""
		if strings.HasSuffix(r.URL.Path, "favicon") {
			fileName = "favicon.ico"
		} else if strings.HasSuffix(r.URL.Path, "logo") {
			fileName = "logo.png"
		} else if strings.HasSuffix(r.URL.Path, "touchicon") {
			fileName = "touchicon.png"
		}

		if fileName != "" {
			err := saveData(r.Body, fileName)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
