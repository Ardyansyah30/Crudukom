package controllers

import (
	"crud-ukom/config"
	"crud-ukom/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all questions
func GetQuestions(c *gin.Context) {
	var questions []models.Question
	config.DB.Find(&questions)
	c.JSON(http.StatusOK, questions)
}

// Get question by ID
func GetQuestionByID(c *gin.Context) {
	var question models.Question
	id := c.Param("id")
	if err := config.DB.First(&question, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found!"})
		return
	}
	c.JSON(http.StatusOK, question)
}

// Create new question
func CreateQuestion(c *gin.Context) {
	var input models.Question
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&input)
	c.JSON(http.StatusOK, input)
}

// Update question by ID
func UpdateQuestion(c *gin.Context) {
	var question models.Question
	id := c.Param("id")
	if err := config.DB.First(&question, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found!"})
		return
	}

	var input models.Question
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&question).Updates(input)
	c.JSON(http.StatusOK, question)
}

// Delete question by ID
func DeleteQuestion(c *gin.Context) {
	var question models.Question
	id := c.Param("id")
	if err := config.DB.First(&question, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found!"})
		return
	}
	config.DB.Delete(&question)
	c.JSON(http.StatusOK, gin.H{"message": "Question deleted successfully!"})
}
