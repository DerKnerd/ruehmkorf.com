package cmd

import (
	"log"
	"ruehmkorf/database"
	"syscall"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
)

func CreateUser(username string) {
	log.Println("Creating user", username)
	var password []byte
	var err error
	for len(password) == 0 {
		log.Print("Password: ")
		password, err = term.ReadPassword(syscall.Stdin)
		if err != nil {
			log.Fatalln("Error reading password:", err)
		}
	}

	hashed, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln("Error hashing password:", err)
	}

	user := database.User{
		Username:    username,
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
