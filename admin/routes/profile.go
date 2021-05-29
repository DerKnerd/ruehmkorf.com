package routes

import (
	"crypto/md5"
	"database/sql"
	"encoding/base64"
	"mime/multipart"
	"net/http"
	"os"
	"ruehmkorf.com/database/models"
	httpUtils "ruehmkorf.com/utils/http"
)

func ProfileList(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		profiles, err := models.FindAllProfiles()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		httpUtils.RenderAdmin("admin/templates/profile/overview.gohtml", profiles, w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

type profileData struct {
	Message string
	Name    string
	Url     string
	Active  bool
}

func saveProfileImage(header *multipart.FileHeader, icon bool) (string, error) {
	path := models.HeaderPath
	if icon {
		path = models.IconPath
	}

	err := os.MkdirAll(path, 0755)
	if err != nil {
		return "", err
	}

	file, err := header.Open()
	if err != nil {
		return "", err
	}

	data := make([]byte, header.Size)
	_, err = file.Read(data)

	if err != nil {
		return "", err
	}

	hash := md5.New()
	hash.Write(data)
	id := base64.RawStdEncoding.EncodeToString(hash.Sum(nil))
	err = os.WriteFile(path+id, data, 0755)

	return path + id, err
}

func ProfileNew(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		httpUtils.RenderAdmin("admin/templates/profile/new.gohtml", profileData{Active: true}, w)
	} else if r.Method == http.MethodPost {
		err := r.ParseMultipartForm(8192 * 1024 * 1024 * 1024)
		if err != nil {
			httpUtils.RenderAdmin("admin/templates/profile/new.gohtml", profileData{Message: err.Error()}, w)
			return
		}

		name := r.FormValue("name")
		url := r.FormValue("url")
		active := r.FormValue("active") == "on"

		_, iconHeader, err := r.FormFile("networkIcon")
		if err != nil {
			httpUtils.RenderAdmin("admin/templates/profile/new.gohtml", profileData{
				Message: "Icon konnte nicht geladen werden",
				Name:    name,
				Url:     url,
				Active:  active,
			}, w)
			return
		}

		iconPath, err := saveProfileImage(iconHeader, true)
		if err != nil {
			httpUtils.RenderAdmin("admin/templates/profile/new.gohtml", profileData{
				Message: "Icon konnte nicht geladen werden",
				Name:    name,
				Url:     url,
				Active:  active,
			}, w)
			return
		}

		_, headerHeader, err := r.FormFile("headerImage")
		headerPath := ""
		if err != http.ErrMissingFile {
			httpUtils.RenderAdmin("admin/templates/profile/new.gohtml", profileData{
				Message: "Header Image konnte nicht geladen werden",
				Name:    name,
				Url:     url,
				Active:  active,
			}, w)
			return
		}

		if headerHeader != nil {
			headerPath, err = saveProfileImage(headerHeader, false)
			if err != nil {
				httpUtils.RenderAdmin("admin/templates/profile/new.gohtml", profileData{
					Message: "Header Image nicht geladen werden",
					Name:    name,
					Url:     url,
					Active:  active,
				}, w)
				return
			}
		}

		profile := models.Profile{
			Name:   name,
			Url:    url,
			Active: active,
			Icon:   iconPath,
			Header: sql.NullString{String: headerPath, Valid: true},
		}

		err = models.CreateProfile(profile)
		if err != nil {
			httpUtils.RenderAdmin("admin/templates/profile/new.gohtml", profileData{
				Message: "Profil konnte nicht gespeichert werden",
				Name:    name,
				Url:     url,
				Active:  active,
			}, w)
			_ = os.Remove(headerPath)
			_ = os.Remove(iconPath)
			return
		}

		http.Redirect(w, r, "/admin/profile", http.StatusFound)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func ProfileEdit(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		profile, err := models.FindProfileById(r.URL.Query().Get("id"))
		data := profileData{}
		if err != nil {
			data.Message = "Profil nicht gefunden"
		} else {
			data.Name = profile.Name
			data.Url = profile.Url
			data.Active = profile.Active
		}
		httpUtils.RenderAdmin("admin/templates/profile/edit.gohtml", data, w)
	} else if r.Method == http.MethodPost {
		profile, err := models.FindProfileById(r.URL.Query().Get("id"))
		if err != nil {
			httpUtils.RenderAdmin("admin/templates/profile/edit.gohtml", profileData{Message: "Profil nicht gefunden"}, w)
			return
		}

		err = r.ParseMultipartForm(8192 * 1024 * 1024 * 1024)
		if err != nil {
			httpUtils.RenderAdmin("admin/templates/profile/edit.gohtml", profileData{Message: err.Error()}, w)
			return
		}

		name := r.FormValue("name")
		url := r.FormValue("url")
		active := r.FormValue("active") == "on"

		_, iconHeader, err := r.FormFile("networkIcon")
		iconPath := profile.Icon
		if err != nil && err != http.ErrMissingFile {
			httpUtils.RenderAdmin("admin/templates/profile/new.gohtml", profileData{
				Message: "Icon konnte nicht geladen werden",
				Name:    name,
				Url:     url,
				Active:  active,
			}, w)
			return
		}

		if iconHeader != nil {
			iconPath, err = saveProfileImage(iconHeader, false)
			if err != nil {
				httpUtils.RenderAdmin("admin/templates/profile/new.gohtml", profileData{
					Message: "Icon nicht geladen werden",
					Name:    name,
					Url:     url,
					Active:  active,
				}, w)
				return
			}
		}

		_, headerHeader, err := r.FormFile("headerImage")
		headerPath := profile.Header.String
		if err != nil && err != http.ErrMissingFile {
			httpUtils.RenderAdmin("admin/templates/profile/new.gohtml", profileData{
				Message: "Header Image konnte nicht geladen werden",
				Name:    name,
				Url:     url,
				Active:  active,
			}, w)
			return
		}

		if iconHeader != nil {
			headerPath, err = saveProfileImage(headerHeader, false)
			if err != nil {
				httpUtils.RenderAdmin("admin/templates/profile/new.gohtml", profileData{
					Message: "Header Image nicht geladen werden",
					Name:    name,
					Url:     url,
					Active:  active,
				}, w)
				return
			}
		}

		profile.Name = name
		profile.Url = url
		profile.Active = active
		profile.Header = sql.NullString{String: headerPath, Valid: true}
		profile.Icon = iconPath

		err = models.UpdateProfile(*profile)
		if err != nil {
			httpUtils.RenderAdmin("admin/templates/news/edit.gohtml", profileData{
				Message: "Profil konnte nicht gespeichert werden",
				Name:    name,
				Url:     url,
				Active:  active,
			}, w)
			_ = os.Remove(headerPath)
			_ = os.Remove(iconPath)
			return
		}

		http.Redirect(w, r, "/admin/profile", http.StatusFound)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func ProfileDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := r.URL.Query().Get("id")
		profile, err := models.FindProfileById(id)
		if err != nil {
			httpUtils.RenderAdmin("admin/templates/profile/delete.gohtml", profileData{
				Message: "Profil nicht gefunden",
			}, w)
			return
		}

		httpUtils.RenderAdmin("admin/templates/profile/delete.gohtml", profileData{
			Name: profile.Name,
		}, w)
	} else if r.Method == http.MethodPost {
		id := r.URL.Query().Get("id")
		profile, err := models.FindProfileById(id)
		if err != nil {
			httpUtils.RenderAdmin("admin/templates/profile/delete.gohtml", newsData{
				Message: "Profil nicht gefunden",
			}, w)
			return
		}

		err = models.DeleteProfile(id)
		if err != nil {
			httpUtils.RenderAdmin("admin/templates/profile/delete.gohtml", newsData{
				Message: "Profil nicht gefunden",
			}, w)
			return
		}

		_ = os.Remove(profile.Icon)
		_ = os.Remove(profile.Header.String)

		http.Redirect(w, r, "/admin/profile", http.StatusFound)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
