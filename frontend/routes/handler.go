package routes

import (
	"net/http"
	"ruehmkorf.com/frontend/middleware"
	httpUtils "ruehmkorf.com/utils/http"
	"strings"
)

func Error404Handler(w http.ResponseWriter, r *http.Request) error {
	errorTemplatePath := ""
	data := middleware.ErrorData{}
	de := strings.Index(strings.ToLower(r.URL.Path), "/de") == 0
	if de {
		errorTemplatePath = "frontend/templates/error/de/404.gohtml"
		data.Language = "de"
		data.Url = strings.TrimPrefix(r.URL.Path, "/de/")
	} else {
		errorTemplatePath = "frontend/templates/error/en/404.gohtml"
		data.Language = "en"
		data.Url = strings.TrimPrefix(r.URL.Path, "/en/")
	}

	w.WriteHeader(http.StatusNotFound)

	return httpUtils.RenderFrontend(errorTemplatePath, data, w)
}
