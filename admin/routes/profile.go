package routes

import (
	"net/http"
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
