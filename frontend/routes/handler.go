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
	if strings.Index(strings.ToLower(r.URL.Path), "/de") == 0 {
		errorTemplatePath = "frontend/templates/error/de/404.gohtml"
		data.Language = "de"
	} else {
		errorTemplatePath = "frontend/templates/error/en/404.gohtml"
		data.Language = "en"
	}

	w.WriteHeader(http.StatusNotFound)

	return httpUtils.RenderFrontend(errorTemplatePath, data, w)
}
