package main

import (
	"crud-ukom/config"
	"crud-ukom/models"
	"crud-ukom/routes"
)

func main() {
	config.ConnectDB()

	// Auto migrate the User model
	config.DB.AutoMigrate(&models.User{})

	router := routes.SetupRoutes()
	router.Run(":8080") // Run the server on port 8080
}
