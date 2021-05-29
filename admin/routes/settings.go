package routes

import (
	"mime/multipart"
	"net/http"
	"os"
	httpUtils "ruehmkorf.com/utils/http"
)

type settingsData struct {
	Message string
}

const SettingsPath = "./data/public/settings/"

func saveData(header *multipart.FileHeader, fileName string) error {
	path := SettingsPath
	err := os.MkdirAll(path, 0755)
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

	err = os.WriteFile(path+fileName, data, 0755)

	return err
}

func SettingsView(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		httpUtils.RenderAdmin("admin/templates/settings/index.gohtml", nil, w)
	} else if r.Method == http.MethodPost {
		err := r.ParseMultipartForm(8192 * 1024 * 1024 * 1024)
		if err != nil {
			httpUtils.RenderAdmin("admin/templates/settings/index.gohtml", settingsData{Message: err.Error()}, w)
			return
		}

		_, faviconHeader, err := r.FormFile("favicon")
		if err != nil && err != http.ErrMissingFile {
			httpUtils.RenderAdmin("admin/templates/settings/index.gohtml", settingsData{
				Message: "Favicon konnte nicht geladen werden",
			}, w)
			return
		}

		if faviconHeader != nil {
			err = saveData(faviconHeader, "favicon.ico")
			if err != nil {
				httpUtils.RenderAdmin("admin/templates/settings/index.gohtml", settingsData{
					Message: "Favicon nicht geladen werden",
				}, w)
				return
			}
		}

		_, logoHeader, err := r.FormFile("logo")
		if err != nil && err != http.ErrMissingFile {
			httpUtils.RenderAdmin("admin/templates/settings/index.gohtml", settingsData{
				Message: "Logo konnte nicht geladen werden",
			}, w)
			return
		}

		if logoHeader != nil {
			err = saveData(logoHeader, "logo.png")
			if err != nil {
				httpUtils.RenderAdmin("admin/templates/settings/index.gohtml", settingsData{
					Message: "Logo nicht geladen werden",
				}, w)
				return
			}
		}

		_, touchIconHeader, err := r.FormFile("touchIcon")
		if err != nil && err != http.ErrMissingFile {
			httpUtils.RenderAdmin("admin/templates/settings/index.gohtml", settingsData{
				Message: "Touch icon konnte nicht geladen werden",
			}, w)
			return
		}

		if touchIconHeader != nil {
			err = saveData(touchIconHeader, "touchicon.png")
			if err != nil {
				httpUtils.RenderAdmin("admin/templates/settings/index.gohtml", settingsData{
					Message: "Touch icon nicht geladen werden",
				}, w)
				return
			}
		}

		http.Redirect(w, r, "/admin/settings", http.StatusFound)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
