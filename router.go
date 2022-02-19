package main

import (
	"io"
	"net/http"
	"os"
	"ruehmkorf.com/admin"
	"ruehmkorf.com/admin/routes"
	"ruehmkorf.com/frontend"
	"ruehmkorf.com/public"
)

func loadIcon(name string, w http.ResponseWriter, r *http.Request) {
	file, err := os.Open(routes.SettingsPath + name)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	_, err = io.Copy(w, file)
	if err != nil {
		http.NotFound(w, r)
		return
	}
}

func InitRouting(mux *http.ServeMux) error {
	if err := admin.InitRouting(mux); err != nil {
		return err
	}

	if err := frontend.InitRouting(mux); err != nil {
		return err
	}

	if err := public.InitRouting(mux); err != nil {
		return err
	}

	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		loadIcon("favicon.ico", w, r)
	})
	mux.HandleFunc("/touchicon.png", func(w http.ResponseWriter, r *http.Request) {
		loadIcon("touchicon.png", w, r)
	})
	mux.HandleFunc("/logo.png", func(w http.ResponseWriter, r *http.Request) {
		loadIcon("logo.png", w, r)
	})

	return nil
}
