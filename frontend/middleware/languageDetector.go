package middleware

import (
	"net/http"
	"strings"
)

func LanguageDetectorMiddleware(action func(http.ResponseWriter, *http.Request, string) error) func(http.ResponseWriter, *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		language := "en"
		if strings.Index(strings.ToLower(r.URL.Path), "/de") == 0 {
			language = "de"
		}
		return action(w, r, language)
	}
}
