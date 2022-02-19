package models

import (
	"database/sql"
	"github.com/thanhpk/randstr"
	"golang.org/x/crypto/bcrypt"
	"ruehmkorf.com/database"
)

// language=sql
const CreateUserTable = `
CREATE TABLE "user" (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name text NOT NULL,
    email text UNIQUE NOT NULL,
    password text NOT NULL,
    activated boolean NOT NULL DEFAULT false,
    two_factor_code text NULL DEFAULT NULL 
)`

type User struct {
	Id            string         `json:"id"`
	Name          string         `json:"name"`
	Email         string         `json:"email"`
	Password      string         `json:"-"`
	Activated     bool           `json:"activated"`
	TwoFactorCode sql.NullString `db:"two_factor_code" json:"-"`
}

var hashCost = 13

func hashPassword(input string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(input), hashCost)

	return string(hashed), err
}

func FindUserById(id string) (*User, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	defer db.Close()
	user := new(User)
	if err = db.Get(user, "SELECT * FROM \"user\" WHERE id = $1", id); err != nil {
		return nil, err
	}

	return user, nil
}

func FindUserByEmail(email string) (*User, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	defer db.Close()
	user := new(User)
	if err = db.Get(user, "SELECT * FROM \"user\" WHERE email = $1", email); err != nil {
		return nil, err
	}

	return user, nil
}

func FindUserByEmailAndPassword(email string, password string) (*User, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	defer db.Close()
	user := new(User)
	if err = db.Get(user, "SELECT * FROM \"user\" WHERE email = $1 AND activated = true", email); err != nil {
		return nil, err
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, err
	}

	return user, nil
}

func FindAllUsers() ([]User, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	defer db.Close()
	users := make([]User, 0)
	if err = db.Select(&users, "SELECT * FROM \"user\""); err != nil {
		return nil, err
	}

	return users, err
}

func CreateUser(user User) (string, error) {
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return "", err
	}

	db, err := database.Connect()
	if err != nil {
		return "", err
	}

	defer db.Close()
	_, err = db.Exec("INSERT INTO \"user\" (name, email, password, activated) VALUES ($1, $2, $3, $4)", user.Name, user.Email, hashedPassword, user.Activated)
	if err != nil {
		return "", err
	}

	createdUser, err := FindUserByEmail(user.Email)

	return createdUser.Id, err
}

func UpdateUser(user User, newPassword bool) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}

	defer db.Close()
	_, err = db.Exec("UPDATE \"user\" SET name = $1, email = $2, activated = $3 WHERE id = $4", user.Name, user.Email, user.Activated, user.Id)
	if err != nil {
		return err
	}

	if newPassword {
		hashedPassword, err := hashPassword(user.Password)
		if err != nil {
			return err
		}

		_, err = db.Exec("UPDATE \"user\" SET password = $1 WHERE id = $2", hashedPassword, user.Id)
		if err != nil {
			return err
		}
	}

	return err
}

func SetTwoFactorCode(user User) (string, error) {
	db, err := database.Connect()
	if err != nil {
		return "", err
	}

	defer db.Close()

	twoFactorCode := randstr.String(6)
	_, err = db.Exec("UPDATE \"user\" SET two_factor_code = $1 WHERE id = $2", twoFactorCode, user.Id)

	return twoFactorCode, err
}

func ResetTwoFactorCode(user User) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec("UPDATE \"user\" SET two_factor_code = null WHERE id = $1", user.Id)

	return err
}

func DeleteUser(id string) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec("DELETE FROM auth_token WHERE user_id = $1", id)
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM \"user\" WHERE id = $1", id)

	return err
}
