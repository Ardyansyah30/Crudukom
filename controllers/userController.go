package controllers

import (
	"crud-ukom/config"
	"crud-ukom/models"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET")) // Mendapatkan JWT secret dari environment variable

// Fungsi untuk memvalidasi apakah nomor telepon hanya berisi angka
func isValidPhoneNumber(phone string) bool {
	regex := regexp.MustCompile(`^[0-9]+$`)
	return regex.MatchString(phone)
}

// Hash password
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Verifikasi password
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Buat token JWT
func generateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // Token berlaku selama 72 jam
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	return tokenString, err
}

// Signup user baru
func Signup(c *gin.Context) {
	var input struct {
		Name        string `json:"Name"`
		Email       string `json:"Email"`
		Password    string `json:"Password"`
		PhoneNumber string `json:"PhoneNumber"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi phone number
	if !isValidPhoneNumber(input.PhoneNumber) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone number must contain only numeric characters"})
		return
	}

	// Hash password
	hashedPassword, err := hashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create user instance
	user := models.User{
		Name:        input.Name,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		Password:    hashedPassword,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Save user to the database
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Generate JWT token
	token, err := generateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Fungsi Login, GetUsers, GetUserByID, UpdateUser, DeleteUser...

// Login user
func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"Email"`
		Password string `json:"Password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Check password
	if !checkPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate JWT token
	token, err := generateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
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

	// Validasi phone number
	if !isValidPhoneNumber(input.PhoneNumber) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone number must contain only numeric characters"})
		return
	}

	// Parse the date from string to time.Time
	dob, err := time.Parse("2006-01-02", input.DateOfBirth)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD."})
		return
	}

	// Update fields
	user.Name = input.Name
	user.Email = input.Email
	user.Password = input.Password
	user.DateOfBirth = dob // Menggunakan dob yang bertipe time.Time
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
