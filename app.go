package main

import (
	"log"
	"net/http"
	"os"
	"ruehmkorf.com/utils"
)

func main() {
	if utils.ContainsString(os.Args, "migrate") {
	} else if utils.ContainsString(os.Args, "start") {
		mux := http.NewServeMux()
		err := InitRouting(mux)
		if err != nil {
			panic(err)
		}

		log.Println("Serving at localhost:8090...")
		log.Fatal(http.ListenAndServe(":8090", mux))
	}
}
