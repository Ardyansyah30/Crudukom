package main

import (
	"crud-ukom/config"
	"crud-ukom/models"
	"crud-ukom/routes"
)

func main() {
	config.ConnectDB()

	config.DB.AutoMigrate(&models.User{}, &models.Question{},
		&models.Packet{}, &models.Exam{}, &models.Order{}, &models.ExamQuestion{})

	router := routes.SetupRoutes()
	router.Run(":8080")
}
