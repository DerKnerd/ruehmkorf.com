package public

import (
	"io"
	"net/http"
	"os"
	"strings"
)

func InitRouting(mux *http.ServeMux) error {
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
		if strings.LastIndex(r.URL.Path, ".css") == len(r.URL.Path)-4 {
			contentType = "text/css"
		}
		if strings.LastIndex(r.URL.Path, ".js") == len(r.URL.Path)-3 {
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
