package middleware

import (
	"net/http"
	"ruehmkorf.com/database/models"
)

func CheckLoginMiddleware(action func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("Auth")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if err = models.GetAuthTokenByToken(cookie.Value); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		action(w, r)
	}
}
