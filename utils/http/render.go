package http

import (
	"html/template"
	"net/http"
)

func RenderSingle(tmpl string, tmplData interface{}, w http.ResponseWriter) {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, tmplData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func RenderAdmin(tmpl string, tmplData interface{}, w http.ResponseWriter) {
	layout, err := template.New("layout").ParseFiles(tmpl, "admin/templates/layout.gohtml")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = layout.ExecuteTemplate(w, "layout", tmplData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}
