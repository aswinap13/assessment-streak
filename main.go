package main

import (
	"log"
	"os"

	"github.com/aswinap13/cmd/routes"
	"github.com/joho/godotenv"
)

func main() {
	r := routes.SetupRouter()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env variables")
	}
	port := os.Getenv("APP_PORT")
	

	r.Run(":" + port)
	log.Fatal(r.Run(":" + port))
}
