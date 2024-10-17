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
		Name                   string `json:"Name"`
		Email                  string `json:"Email"`
		Password               string `json:"Password"`
		DateOfBirth            string `json:"DateOfBirth"` // Accept the date as a string
		Gender                 string `json:"Gender"`
		PhoneNumber            string `json:"PhoneNumber"`
		EducationalInstitution string `json:"EducationalInstitution"`
		Profession             string `json:"Profession"`
		Address                string `json:"Address"`
		Province               string `json:"Province"`
		City                   string `json:"City"`
	}

	// Bind the JSON to input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse the date from string to time.Time
	dob, err := time.Parse("2006-01-02", input.DateOfBirth)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD."})
		return
	}

	// Create user instance
	user := models.User{
		Name:                   input.Name,
		Email:                  input.Email,
		Password:               input.Password,
		DateOfBirth:            models.DateOfBirth(dob),
		Gender:                 input.Gender,
		PhoneNumber:            input.PhoneNumber,
		EducationalInstitution: input.EducationalInstitution,
		Profession:             input.Profession,
		Address:                input.Address,
		Province:               input.Province,
		City:                   input.City,
		CreatedAt:              time.Now(),
		UpdatedAt:              time.Now(),
	}

	// Save user to the database
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Get all users
func GetUsers(c *gin.Context) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
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

	var input struct {
		Name                   string `json:"Name"`
		Email                  string `json:"Email"`
		Password               string `json:"Password"`
		DateOfBirth            string `json:"DateOfBirth"`
		Gender                 string `json:"Gender"`
		PhoneNumber            string `json:"PhoneNumber"`
		EducationalInstitution string `json:"EducationalInstitution"`
		Profession             string `json:"Profession"`
		Address                string `json:"Address"`
		Province               string `json:"Province"`
		City                   string `json:"City"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dob, err := time.Parse("2006-01-02", input.DateOfBirth)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD."})
		return
	}

	// Update fields
	user.Name = input.Name
	user.Email = input.Email
	user.Password = input.Password
	user.DateOfBirth = models.DateOfBirth(dob)
	user.Gender = input.Gender
	user.PhoneNumber = input.PhoneNumber
	user.EducationalInstitution = input.EducationalInstitution
	user.Profession = input.Profession
	user.Address = input.Address
	user.Province = input.Province
	user.City = input.City
	user.UpdatedAt = time.Now()

	// Save the updated user to the database
	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// Delete a user by ID
func DeleteUser(c *gin.Context) {
	var user models.User
	if err := config.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
