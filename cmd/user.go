package cmd

import (
	"log"
	"syscall"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"

	"ruehmkorf/database"
)

func CreateUser(username string) {
	log.Println("Creating user", username)
	log.Print("Password: ")
	password, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		log.Fatalln("Error reading password:", err)
	}

	hashed, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln("Error hashing password:", err)
	}

	user := database.User{
		Email:       username,
		Password:    string(hashed),
		TotpEnabled: false,
		TotpSecret:  "",
	}

	err = database.GetDbMap().Insert(&user)
	if err != nil {
		log.Fatalln("Error inserting user:", err)
	}

	log.Println("User created")
}
