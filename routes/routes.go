package routes

import (
	"crud-ukom/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// User routes
	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id", controllers.GetUserByID)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)

	// Question routes
	r.POST("/", controllers.CreateQuestion)
	r.GET("/", controllers.GetQuestions)
	r.GET("/:id", controllers.GetQuestionByID)
	r.PUT("/:id", controllers.UpdateQuestion)
	r.DELETE("/:id", controllers.DeleteQuestion)
	r.GET("/package/:id_package", controllers.GetQuestionsByPackageID)

	// Scoring route
	r.POST("/score", controllers.CalculateScore) // Calculate user score

	// Packet routes
	r.GET("/packets", controllers.GetPackets)
	r.GET("/packets/:id", controllers.GetPacketByID)
	r.POST("/packets", controllers.CreatePacket)
	r.PUT("/packets/:id", controllers.UpdatePacket)
	r.DELETE("/packets/:id", controllers.DeletePacket)

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
