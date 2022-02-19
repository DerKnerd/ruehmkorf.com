package middleware

import (
	"net/http"
	httpUtils "ruehmkorf.com/utils/http"
	"strings"
)

type ErrorData struct {
	Language string
	Url      string
	Host     string
}

func ErrorHandlerMiddleware(action func(http.ResponseWriter, *http.Request) error) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := action(w, r)
		if err != nil {
			errorTemplatePath := ""
			data := ErrorData{
				Host: r.Host,
			}
			de := strings.Index(strings.ToLower(r.URL.Path), "/de") == 0
			if de {
				errorTemplatePath = "frontend/templates/error/de/500.gohtml"
				data.Language = "de"
				data.Url = strings.TrimPrefix(r.URL.Path, "/de/")
			} else {
				errorTemplatePath = "frontend/templates/error/en/500.gohtml"
				data.Language = "en"
				data.Url = strings.TrimPrefix(r.URL.Path, "/en/")
			}

			w.WriteHeader(http.StatusInternalServerError)
			err = httpUtils.RenderFrontend(errorTemplatePath, data, w)
			if err != nil {
				http.Error(w, "The loading failed", http.StatusInternalServerError)
			}
		}
	}
}
