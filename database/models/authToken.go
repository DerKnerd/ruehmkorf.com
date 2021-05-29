package models

import (
	"github.com/satori/go.uuid"
	"log"
	"ruehmkorf.com/database"
)

// language=sql
var CreateAuthTokenTable = `
CREATE TABLE "auth_token" (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    token text NOT NULL UNIQUE,
    two_factor_approved boolean NOT NULL DEFAULT false,
    last_used_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    user_id uuid REFERENCES "user"(id) NOT NULL
)`

type AuthToken struct {
	Id                string
	Token             string
	UserId            string `db:"user_id"`
	TwoFactorApproved bool   `db:"two_factor_approved"`
}

func GetAuthTokenByToken(token string) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}

	defer db.Close()
	authToken := new(AuthToken)

	err = db.Get(authToken, "SELECT id, token, user_id FROM auth_token WHERE token = $1 AND last_used_at < CURRENT_TIMESTAMP + interval '24 hour' AND two_factor_approved = true", token)
	if err == nil {
		_, err = db.Exec("UPDATE auth_token SET last_used_at = CURRENT_TIMESTAMP WHERE id = $1", authToken.Id)
	} else {
		_ = DeleteAuthToken(token)
	}

	return err
}

func DeleteAuthToken(token string) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec("DELETE FROM \"auth_token\" WHERE token = $1", token)

	return err
}

func CreateAuthToken(userId string) (string, error) {
	token := uuid.NewV4()
	db, err := database.Connect()
	if err != nil {
		return "", err
	}

	defer db.Close()

	_, err = db.Exec("INSERT INTO \"auth_token\" (token, user_id) VALUES ($1, $2)", token.String(), userId)

	return token.String(), nil
}

func TwoFactorApprove(token string, twoFactorToken string) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}

	defer db.Close()
	user := new(User)

	err = db.Get(user, "SELECT * FROM \"user\" WHERE two_factor_code = $1", twoFactorToken)
	if err != nil {
		return err
	}

	result, err := db.Exec("UPDATE auth_token SET two_factor_approved = true, last_used_at = CURRENT_TIMESTAMP WHERE token = $1 AND user_id = $2", token, user.Id)

	log.Printf("%v", result)

	return err
}
