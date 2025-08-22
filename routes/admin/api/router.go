package api

import (
	"context"
	"net/http"
	"ruehmkorf/database"
	"strings"

	"github.com/gorilla/mux"
)

func contentTypeJson(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, req)
	})
}

func checkAuth() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var auth string
			authCookie, err := r.Cookie("authentication")
			if err != nil {
				auth = strings.TrimLeft(r.Header.Get("Authorization"), "Bearer ")
			} else {
				auth = authCookie.Value
			}

			if auth == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			user, err := database.SelectOne[database.User](`select u.* from "user" u inner join token t on t.user_id = u.id where t.token = $1`, auth)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			token, err := database.SelectOne[database.Token](`select t.* from "token" t where t.token = $1`, auth)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r.WithContext(context.WithValue(context.WithValue(r.Context(), "user", user), "token", token)))
		})
	}
}

func SetupRouter(router *mux.Router) {
	apiRouter := router.PathPrefix("/api").Subrouter()
	authSubRouter := apiRouter.PathPrefix("/authentication").Subrouter()

	authSubRouter.
		Methods(http.MethodPost).
		Path("/login").
		HandlerFunc(login)
	authSubRouter.
		Methods(http.MethodDelete).
		Path("/login").
		Handler(checkAuth()(http.HandlerFunc(logout)))
	authSubRouter.
		Methods(http.MethodPost).
		Path("/2fa").
		Handler(http.HandlerFunc(setup2fa))
	authSubRouter.
		Methods(http.MethodPut).
		Path("/password").
		Handler(checkAuth()(http.HandlerFunc(changePassword)))
	authSubRouter.
		Methods(http.MethodHead).
		Path("/login").
		Handler(checkAuth()(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
		})))

	apiRouter.Use(contentTypeJson)
}
