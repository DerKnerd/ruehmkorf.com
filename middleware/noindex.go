package middleware

import (
	"net/http"
)

func NoIndexMiddleware(action func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Robots-Tag", "none")
		action(w, r)
	}
}
