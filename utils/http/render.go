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

func RenderFrontend(tmpl string, tmplData interface{}, w http.ResponseWriter) error {
	layout, err := template.New("layout").Funcs(template.FuncMap{
		"unsafe": func(data string) template.HTML {
			return template.HTML(data)
		},
	}).ParseFiles(tmpl, "frontend/templates/layout.gohtml")
	if err != nil {
		return err
	}

	return layout.ExecuteTemplate(w, "layout", tmplData)
}
