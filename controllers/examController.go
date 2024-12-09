package controllers

import (
	"crud-ukom/config"
	"crud-ukom/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Create a new exam
func CreateExam(c *gin.Context) {
	var exam models.Exam
	if err := c.ShouldBindJSON(&exam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simpan data ke database
	if err := config.DB.Create(&exam).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Hitung durasi ujian
	duration := GetExamDuration(exam.StartTime, exam.EndTime)

	// Kirim respons bersama dengan durasi
	c.JSON(http.StatusOK, gin.H{
		"exam":     exam,
		"duration": duration.Minutes(), // durasi dalam menit
	})
}

// Get all exams
func GetExams(c *gin.Context) {
	var exams []models.Exam
	if err := config.DB.Find(&exams).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, exams)
}

// Get an exam by ID
func GetExamByID(c *gin.Context) {
	var exam models.Exam
	if err := config.DB.First(&exam, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exam not found"})
		return
	}
	c.JSON(http.StatusOK, exam)
}

// Update an exam by ID
func UpdateExam(c *gin.Context) {
	var exam models.Exam
	if err := config.DB.First(&exam, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exam not found"})
		return
	}

	var input models.Exam
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update field exam dengan input baru
	exam.IDPacket = input.IDPacket
	exam.IDUser = input.IDUser
	exam.NameExam = input.NameExam
	exam.StartTime = input.StartTime
	exam.EndTime = input.EndTime
	exam.PaymentStatus = input.PaymentStatus
	exam.Score = input.Score
	exam.UpdatedAt = time.Now()

	config.DB.Save(&exam)
	c.JSON(http.StatusOK, exam)
}

// Delete an exam by ID
func DeleteExam(c *gin.Context) {
	if err := config.DB.Delete(&models.Exam{}, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exam not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Exam deleted successfully"})
}

// Fungsi untuk menghitung durasi ujian
func GetExamDuration(startTime, endTime time.Time) time.Duration {
	return endTime.Sub(startTime)
}
