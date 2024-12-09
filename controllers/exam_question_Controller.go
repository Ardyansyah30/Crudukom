package controllers

import (
	"crud-ukom/config"
	"crud-ukom/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Create a new ExamQuestion
func CreateExamQuestion(c *gin.Context) {
	var input struct {
		IDExam     int `json:"id_exam" binding:"required"`
		IDQuestion int `json:"id_question" binding:"required"`
		UserAnswer int `json:"user_answer" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	examQuestion := models.ExamQuestion{
		IDExam:     input.IDExam,
		IDQuestion: input.IDQuestion,
		UserAnswer: input.UserAnswer,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	if err := config.DB.Create(&examQuestion).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, examQuestion)
}

// Get all ExamQuestions
func GetExamQuestions(c *gin.Context) {
	var examQuestions []models.ExamQuestion
	config.DB.Find(&examQuestions)
	c.JSON(http.StatusOK, examQuestions)
}

// Get ExamQuestion by ID
func GetExamQuestionByID(c *gin.Context) {
	var examQuestion models.ExamQuestion
	if err := config.DB.Where("id = ?", c.Param("id")).First(&examQuestion).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ExamQuestion not found"})
		return
	}
	c.JSON(http.StatusOK, examQuestion)
}

// Update an ExamQuestion by ID
func UpdateExamQuestion(c *gin.Context) {
	var examQuestion models.ExamQuestion
	if err := config.DB.Where("id = ?", c.Param("id")).First(&examQuestion).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ExamQuestion not found"})
		return
	}

	var input struct {
		IDExam     int `json:"id_exam" binding:"required"`
		IDQuestion int `json:"id_question" binding:"required"`
		UserAnswer int `json:"user_answer" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	examQuestion.IDExam = input.IDExam
	examQuestion.IDQuestion = input.IDQuestion
	examQuestion.UserAnswer = input.UserAnswer
	examQuestion.UpdatedAt = time.Now()

	config.DB.Save(&examQuestion)
	c.JSON(http.StatusOK, examQuestion)
}

// Delete an ExamQuestion by ID
func DeleteExamQuestion(c *gin.Context) {
	var examQuestion models.ExamQuestion
	if err := config.DB.Where("id = ?", c.Param("id")).First(&examQuestion).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ExamQuestion not found"})
		return
	}

	config.DB.Delete(&examQuestion)
	c.JSON(http.StatusOK, gin.H{"message": "ExamQuestion deleted successfully"})
}
