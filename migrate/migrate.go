package main

import (
	"address-book-backend/initializers"
	"address-book-backend/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.AddressBook{})
}
