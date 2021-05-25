package main

import (
	"github.com/thanhpk/randstr"
	"log"
	"net/http"
	"os"
	"ruehmkorf.com/database/migrations"
	"ruehmkorf.com/database/models"
	"ruehmkorf.com/utils"
)

func main() {
	if utils.ContainsString(os.Args, "migrate") {
		err := migrations.Migrate()
		if err != nil {
			panic(err)
		}
	}

	if utils.ContainsString(os.Args, "install") {
		name := os.Getenv("FIRST_ADMIN_NAME")
		email := os.Getenv("FIRST_ADMIN_EMAIL")
		password := os.Getenv("FIRST_ADMIN_PASSWORD")
		if password == "" {
			password = randstr.String(20)
		}

		_, err := models.FindUserByEmail(email)
		if err != nil {
			if models.CreateUser(models.User{
				Name:      name,
				Email:     email,
				Password:  password,
				Activated: true,
			}) != nil {
				panic(err)
			} else {
				log.Printf("Password: %s", password)
			}
		}
	}

	if utils.ContainsString(os.Args, "start") {
		mux := http.NewServeMux()
		err := InitRouting(mux)
		if err != nil {
			panic(err)
		}

		log.Println("Serving at localhost:8090...")
		log.Fatal(http.ListenAndServe(":8090", mux))
	}
}
