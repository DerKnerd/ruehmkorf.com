package public

import (
	"io"
	"net/http"
	"os"
	"ruehmkorf.com/database/models"
	"strings"
)

func InitRouting(mux *http.ServeMux) error {
	mux.HandleFunc("/public/admin/", func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open("./" + r.URL.Path)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		data, err := io.ReadAll(file)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		contentType := http.DetectContentType(data)
		if strings.HasSuffix(r.URL.Path, ".css") {
			contentType = "text/css"
		}
		if strings.HasSuffix(r.URL.Path, ".js") {
			contentType = "application/javascript"
		}

		if contentType == "application/javascript" {
			cookie, err := r.Cookie("Auth")
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			if err = models.GetAuthTokenByToken(cookie.Value); err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}
		}

		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", contentType)
			w.Header().Set("X-Robots-Tag", "none")
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/public/", func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open("./" + r.URL.Path)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		data, err := io.ReadAll(file)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		contentType := http.DetectContentType(data)
		if strings.HasSuffix(r.URL.Path, ".css") {
			contentType = "text/css"
		}
		if strings.HasSuffix(r.URL.Path, ".js") {
			contentType = "application/javascript"
		}

		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", contentType)
			w.Header().Set("X-Robots-Tag", "none")
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	return nil
}
