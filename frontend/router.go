package frontend

import "net/http"

func InitRouting(mux *http.ServeMux) error {
	mux.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("Hello frontend world"))
	})

	return nil
}
