package routes

import (
	"net/http"
	"ruehmkorf.com/database/models"
	"ruehmkorf.com/mailer"
	"time"

	httpUtils "ruehmkorf.com/utils/http"
)

type loginData struct {
	Message string
	Email   string
}

func TwoFactor(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		cookie, err := r.Cookie("Auth")
		if err != nil {
			data := loginData{Message: "Die Zugangsdaten sind ungültig"}
			w.WriteHeader(http.StatusUnauthorized)
			httpUtils.RenderSingle("admin/templates/login/twoFactor.gohtml", data, w)
			return
		}

		twoFactorToken := r.PostFormValue("token")
		authToken := cookie.Value

		err = models.TwoFactorApprove(authToken, twoFactorToken)
		if err != nil {
			data := loginData{Message: "Die Zugangsdaten sind ungültig"}
			w.WriteHeader(http.StatusUnauthorized)
			httpUtils.RenderSingle("admin/templates/login/twoFactor.gohtml", data, w)
			return
		}

		http.Redirect(w, r, "/admin", http.StatusFound)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		httpUtils.RenderSingle("admin/templates/login/login.gohtml", nil, w)
	} else if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		email := r.PostFormValue("email")
		password := r.PostFormValue("password")
		user, err := models.FindUserByEmailAndPassword(email, password)
		if err != nil || user == nil {
			data := loginData{Message: "Die Zugangsdaten sind ungültig", Email: email}
			w.WriteHeader(http.StatusUnauthorized)
			httpUtils.RenderSingle("admin/templates/login/login.gohtml", data, w)
			return
		}

		authToken, err := models.CreateAuthToken(user.Id)
		if err != nil {
			data := loginData{Message: "Die Zugangsdaten sind ungültig", Email: email}
			w.WriteHeader(http.StatusUnauthorized)
			httpUtils.RenderSingle("admin/templates/login/login.gohtml", data, w)
			return
		}

		code, err := models.SetTwoFactorCode(*user)
		if err != nil {
			_ = models.DeleteAuthToken(authToken)
			data := loginData{Message: "Die Zugangsdaten sind ungültig", Email: email}
			w.WriteHeader(http.StatusUnauthorized)
			httpUtils.RenderSingle("admin/templates/login/login.gohtml", data, w)
			return
		}

		if mailer.SendTwoFactorMail(code, user.Email) != nil {
			_ = models.DeleteAuthToken(authToken)
			data := loginData{Message: "Die 2FA Email konnte nicht versendet werden", Email: email}
			w.WriteHeader(http.StatusUnauthorized)
			httpUtils.RenderSingle("admin/templates/login/login.gohtml", data, w)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "Auth",
			Value:    authToken,
			Expires:  time.Unix(time.Now().Add(time.Hour*24).Unix(), 0),
			HttpOnly: true,
			Path:     "/",
		})

		httpUtils.RenderSingle("admin/templates/login/twoFactor.gohtml", nil, w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("Auth")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if err = models.GetAuthTokenByToken(cookie.Value); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = models.DeleteAuthToken(cookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin", http.StatusFound)
}
