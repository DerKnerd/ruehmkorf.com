package public

import (
	"io"
	"net/http"
	"os"
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

		if r.Method == http.MethodGet {
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		}
	})

	return nil
}
