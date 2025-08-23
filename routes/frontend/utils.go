package frontend

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"
)

//go:embed tmpl
var tmplFs embed.FS

func getTemplateFs() fs.FS {
	if os.Getenv("ENV") == "dev" {
		return os.DirFS("./routes/frontend")
	} else {
		return tmplFs
	}
}

func renderPage(w http.ResponseWriter, r *http.Request, tmpl string, data any) {
	t, err := template.New("layout").Funcs(map[string]any{
		"year": func() string {
			return time.Now().Format("2006")
		},
		"noescape": func(s string) template.HTML {
			return template.HTML(s)
		},
		"last": func(x int, a any) bool {
			return x == reflect.ValueOf(a).Len()-1
		},
		"menuItem": func(label string, href string) template.HTML {
			if strings.HasSuffix(r.URL.Path, href) {
				return template.HTML(fmt.Sprintf(`
<li class="govuk-service-navigation__item govuk-service-navigation__item--active">
	<a class="govuk-service-navigation__link" href="/%s" aria-current="true">
		<strong class="govuk-service-navigation__active-fallback">
			%s
		</strong>
	</a>
</li>`, href, label))
			} else {
				return template.HTML(fmt.Sprintf(`
<li class="govuk-service-navigation__item">
	<a class="govuk-service-navigation__link" href="/%s" aria-current="false">
		%s
	</a>
</li>`, href, label))
			}
		},
	}).ParseFS(getTemplateFs(), "tmpl/layout.gohtml", fmt.Sprintf("tmpl/%s.gohtml", tmpl))
	if err == nil {
		err = t.Execute(w, data)
		if err != nil {
			log.Printf("Error rendering template: %s", err)
		}
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
