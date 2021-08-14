package main

import (
	"github.com/thanhpk/randstr"
	"log"
	"net/http"
	"os"
	"ruehmkorf.com/bom"
	"ruehmkorf.com/database/migrations"
	"ruehmkorf.com/database/models"
	"ruehmkorf.com/utils"
)

func find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return -1
}

func main() {
	if utils.ContainsString(os.Args, "import-bom") {
		idx := find(os.Args, "import-bom")
		if idx == -1 {
			log.Panicln("Path must be provided")
		}

		path := os.Args[idx+1]
		if err := bom.ImportBomRunes(path); err != nil {
			log.Panicln(err.Error())
		}

		return
	}

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
			if _, err = models.CreateUser(models.User{
				Name:      name,
				Email:     email,
				Password:  password,
				Activated: true,
			}); err != nil {
				panic(err)
			} else {
				log.Printf("Password: %s", password)
			}
		}
	}

	if utils.ContainsString(os.Args, "selfdestruct") {
		downloads, err := models.FindAllDownloadsToSelfDestruct()
		if err != nil {
			log.Panicln(err.Error())
			return
		}

		for _, download := range downloads {
			err = os.Remove(models.DownloadPreviewImagePath + download.Slug)
			if err != nil {
				log.Println(err.Error())
			}

			err = os.Remove(models.DownloadFilePath + download.Slug)
			if err != nil {
				log.Println(err.Error())
			}

			err = models.DeleteDownloadBySlug(download.Slug)
			if err != nil {
				log.Println(err.Error())
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
