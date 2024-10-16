package controllers

import (
	"crud-ukom/config"
	"crud-ukom/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Create a new user
func CreateUser(c *gin.Context) {
	var input struct {
		Name                           string `json:"Name"`
		Email                          string `json:"Email"`
		Password                       string `json:"Password"`
		DateOfBirth                    string `json:"DateOfBirth"` // We receive the date as a string
		Gender                         string `json:"Gender"`
		PhoneNumber                    string `json:"PhoneNumber"`
		CurrentEducationalInstitutions string `json:"CurrentEducationalInstitutions"`
		Profession                     string `json:"Profession"`
		Address                        string `json:"Address"`
		Province                       string `json:"Province"`
		City                           string `json:"City"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse DateOfBirth from string to time.Time
	layout := "2006-01-02"
	dob, err := time.Parse(layout, input.DateOfBirth)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD."})
		return
	}

	// Create a new User struct with the parsed date
	user := models.User{
		Name:                           input.Name,
		Email:                          input.Email,
		Password:                       input.Password,
		DateOfBirth:                    dob, // Store parsed time.Time
		Gender:                         input.Gender,
		PhoneNumber:                    input.PhoneNumber,
		CurrentEducationalInstitutions: input.CurrentEducationalInstitutions,
		Profession:                     input.Profession,
		Address:                        input.Address,
		Province:                       input.Province,
		City:                           input.City,
		CreatedAt:                      time.Now(),
		UpdatedAt:                      time.Now(),
	}

	// Save the user to the database
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Get all users
func GetUsers(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

// Get a user by ID
func GetUserByID(c *gin.Context) {
	var user models.User
	if err := config.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// Update a user by ID
func UpdateUser(c *gin.Context) {
	var user models.User
	if err := config.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.UpdatedAt = time.Now()
	config.DB.Save(&user)
	c.JSON(http.StatusOK, user)
}

// Delete a user by ID
func DeleteUser(c *gin.Context) {
	var user models.User
	if err := config.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	config.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
