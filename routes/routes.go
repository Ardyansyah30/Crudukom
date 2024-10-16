package routes

import (
	"crud-ukom/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	// User Routes
	router.POST("/users", controllers.CreateUser)
	router.GET("/users", controllers.GetUsers)
	router.GET("/users/:id", controllers.GetUserByID)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)

	// Question routes
	router.GET("/questions", controllers.GetQuestions)
	router.GET("/questions/:id", controllers.GetQuestionByID)
	router.POST("/questions", controllers.CreateQuestion)
	router.PUT("/questions/:id", controllers.UpdateQuestion)
	router.DELETE("/questions/:id", controllers.DeleteQuestion)

	// Packet routes
	router.GET("/packets", controllers.GetPackets)
	router.GET("/packets/:id", controllers.GetPacketByID)
	router.POST("/packets", controllers.CreatePacket)
	router.PUT("/packets/:id", controllers.UpdatePacket)
	router.DELETE("/packets/:id", controllers.DeletePacket)

	return router
}
