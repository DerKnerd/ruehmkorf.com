package main

import (
	"net/http"
	"ruehmkorf.com/admin"
	"ruehmkorf.com/frontend"
	"ruehmkorf.com/public"
)

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

	return nil
}
