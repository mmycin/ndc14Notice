package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mmycin/backendv2/router"
	"github.com/mmycin/mongorm/utils"
)

func main() {
	err := godotenv.Load()
	utils.HandleError(err)
	var port string
	port = os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	var URL string = "localhost:" + port
	app := router.Router()
	app.Run(URL)
	log.Println("Listening on ", URL)
}
