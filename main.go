package main

import (
	"crud-ukom/config"
	"crud-ukom/models"
	"crud-ukom/routes"
	"log"
)

func main() {
	// Initialize database connection
	config.ConnectDB()

	// Auto migrate database models
	err := config.DB.AutoMigrate(&models.User{}, &models.Question{},
		&models.Packet{}, &models.Exam{}, &models.Order{},
		&models.ExamQuestion{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Setup and run router
	router := routes.SetupRoutes()
	err = router.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
