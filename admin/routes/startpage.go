package routes

import "net/http"

func Startpage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
