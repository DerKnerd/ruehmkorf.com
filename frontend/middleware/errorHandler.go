package middleware

import (
	"net/http"
	httpUtils "ruehmkorf.com/utils/http"
	"strings"
)

type ErrorData struct {
	Language string
}

func ErrorHandlerMiddleware(action func(http.ResponseWriter, *http.Request) error) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := action(w, r)
		if err != nil {
			errorTemplatePath := ""
			data := ErrorData{}
			if strings.Index(strings.ToLower(r.URL.Path), "/de") == 0 {
				errorTemplatePath = "frontend/templates/error/de/500.gohtml"
				data.Language = "de"
			} else {
				errorTemplatePath = "frontend/templates/error/en/500.gohtml"
				data.Language = "en"
			}

			w.WriteHeader(http.StatusInternalServerError)
			err = httpUtils.RenderFrontend(errorTemplatePath, data, w)
			if err != nil {
				http.Error(w, "The loading failed", http.StatusInternalServerError)
			}
		}
	}
}
