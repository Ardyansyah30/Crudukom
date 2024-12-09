package routes

import (
	"crud-ukom/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// User routes			//ryan
	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id", controllers.GetUserByID)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Question routes		//ryan
	r.POST("/questions", controllers.CreateQuestion)
	r.GET("/questions", controllers.GetQuestions)
	r.GET("/questions/:id", controllers.GetQuestionByID)
	r.PUT("/questions/:id", controllers.UpdateQuestion)
	r.DELETE("/questions/:id", controllers.DeleteQuestion)

	// Packet routes		//ryan
	// Route untuk paket
	r.GET("/packets", controllers.GetPackets)
	r.GET("/packets/:packet_id", controllers.GetPacketByID)
	r.POST("/packets", controllers.CreatePacket)
	r.PUT("/packets/:packet_id", controllers.UpdatePacket)
	r.DELETE("/packets/:packet_id", controllers.DeletePacket)

	// Exam Routes
	r.POST("/exams", controllers.CreateExam)
	r.GET("/exams", controllers.GetExams)
	r.GET("/exams/:id", controllers.GetExamByID)
	r.PUT("/exams/:id", controllers.UpdateExam)
	r.DELETE("/exams/:id", controllers.DeleteExam)

	// Order Routes
	r.POST("/orders", controllers.CreateOrder)
	r.GET("/orders", controllers.GetOrders)
	r.GET("/orders/:id", controllers.GetOrderByID)
	r.PUT("/orders/:id", controllers.UpdateOrder)
	r.DELETE("/orders/:id", controllers.DeleteOrder)

	// Exam_questions Routes
	r.POST("/exam_questions", controllers.CreateExamQuestion)
	r.GET("/exam_questions", controllers.GetExamQuestions)
	r.GET("/exam_questions/:id", controllers.GetExamQuestionByID)
	r.PUT("/exam_questions/:id", controllers.UpdateExamQuestion)
	r.DELETE("/exam_questions/:id", controllers.DeleteExamQuestion)

	return r
}
