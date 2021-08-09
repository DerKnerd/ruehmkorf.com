package routes

import (
	"io/ioutil"
	"net/http"
	"os"
	"ruehmkorf.com/database/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("Auth")
	if err != nil {
		http.Redirect(w, r, "/admin/login", http.StatusFound)
		return
	}

	if err = models.GetAuthTokenByToken(cookie.Value); err != nil {
		http.Redirect(w, r, "/admin/login", http.StatusFound)
		return
	}

	file, err := os.Open("admin/templates/home/home.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data)
}
