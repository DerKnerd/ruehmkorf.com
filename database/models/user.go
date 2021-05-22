package models

import (
	"golang.org/x/crypto/bcrypt"
	"ruehmkorf.com/database"
)

// language=sql
var CreateUserTable = `
CREATE TABLE "user" (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name text NOT NULL,
    email text UNIQUE NOT NULL,
    password text NOT NULL,
    twoFactorCode text NULL DEFAULT NULL 
)`

type User struct {
	Id            string
	Name          string
	Email         string
	Password      string
	TwoFactorCode string
}

var hashCost = 13

func hashPassword(input string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(input), hashCost)

	return string(hashed), err
}

func FindById(id string) (*User, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	user := new(User)
	if err = db.Get(&user, "SELECT id, name, email, password, twoFactorCode FROM \"user\" WHERE id = $1", id); err != nil {
		return nil, err
	}

	return user, nil
}

func FindByEmailAndPassword(email string, password string) (*User, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	user := new(User)
	if err = db.Get(&user, "SELECT id, name, email, password, twoFactorCode FROM \"user\" WHERE email = $1", email); err != nil {
		return nil, err
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, err
	}

	return user, nil
}

func CreateUser(user User) error {
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}

	db, err := database.Connect()
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO \"user\" (name, email, password) VALUES ($1, $2, $3)", user.Name, user.Email, hashedPassword)

	return err
}
