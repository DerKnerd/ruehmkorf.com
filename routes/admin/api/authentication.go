package api

import (
	"encoding/json"
	"net/http"
	"os"
	"ruehmkorf/database"

	"github.com/google/uuid"
	"github.com/pquerna/otp/totp"
	"golang.org/x/crypto/bcrypt"
)

func login(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	var body struct {
		Username      string `json:"username"`
		Password      string `json:"password"`
		TwoFactorCode string `json:"twoFactorCode"`
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := database.SelectOne[database.User](`select u.* from "user" u where u.email = $1`, body.Username)
	if err != nil {
		http.Error(w, "Invalid auth", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		http.Error(w, "Invalid auth", http.StatusUnauthorized)
		return
	}

	if body.TwoFactorCode != "" && user.TotpEnabled {
		totpValid := totp.Validate(body.TwoFactorCode, user.TotpSecret)
		if !totpValid {
			http.Error(w, "Invalid auth", http.StatusUnauthorized)
			return
		}
	} else if user.TotpEnabled {
		encoder.Encode(map[string]string{
			"message": "2fa_required",
		})
		w.WriteHeader(http.StatusOK)
		return
	}

	token := database.Token{
		UserId: user.Id,
		Token:  uuid.NewString(),
	}

	database.GetDbMap().Insert(&token)

	http.SetCookie(w, &http.Cookie{
		Name:     "authentication",
		Value:    token.Token,
		Quoted:   false,
		Path:     "/",
		Secure:   os.Getenv("ENV") != "dev",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})

	encoder.Encode(map[string]any{
		"token":          token.Token,
		"2faSetupNeeded": !user.TotpEnabled,
	})
}

func setup2fa(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Secret string `json:"secret"`
		Code   string `json:"code"`
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := r.Context().Value("user").(database.User)

	totpValid := totp.Validate(body.Code, body.Secret)
	if totpValid {
		user.TotpSecret = body.Secret
		user.TotpEnabled = true
		database.GetDbMap().Update(&user)
		w.WriteHeader(http.StatusNoContent)
	} else {
		http.Error(w, "Invalid code", http.StatusBadRequest)
	}
}

func changePassword(w http.ResponseWriter, r *http.Request) {
	var body struct {
		OldPassword string `json:"oldPassword"`
		NewPassword string `json:"newPassword"`
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := r.Context().Value("user").(database.User)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.OldPassword))
	if err != nil {
		http.Error(w, "Invalid password", http.StatusBadRequest)
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(body.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	user.Password = string(hashed)
	database.GetDbMap().Update(&user)
	w.WriteHeader(http.StatusNoContent)
}
