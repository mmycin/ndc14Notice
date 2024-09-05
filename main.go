package main

import (
	"log"
	"os"

	"github.com/mmycin/ndc14notice/router"
)

func main() {
	var port string
	port = os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	var URL string = "0.0.0.0:" + port
	app := router.Router()
	app.Run(URL)
	log.Println("Listening on ", URL)
}
