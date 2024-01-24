package main

import (
	"address-book-backend/initializers"
	"address-book-backend/routers"
	"fmt"
	"log"
	"os"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	router := routers.SetupRouter()

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router.Run(fmt.Sprintf(":%s", port))
}
