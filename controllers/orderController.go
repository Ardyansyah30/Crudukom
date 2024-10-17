package controllers

import (
	"crud-ukom/config"
	"crud-ukom/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Create a new Order
func CreateOrder(c *gin.Context) {
	var input struct {
		IDUser        int     `json:"id_user" binding:"required"`
		IDPacket      int     `json:"id_packet" binding:"required"`
		PaymentStatus string  `json:"payment_status" binding:"required"`
		OrderDate     string  `json:"order_date" binding:"required"` // Expecting date as string
		TotalPrice    float64 `json:"total_price" binding:"required"`
	}

	// Bind the JSON input
	if err := c.ShouldBindJSON(&input); err != nil {
		// Respond with error if validation fails
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse the date string into time.Time format
	orderDate, err := time.Parse("2006-01-02", input.OrderDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
		return
	}

	// Create an order object
	order := models.Order{
		IDUser:        input.IDUser,
		IDPacket:      input.IDPacket,
		PaymentStatus: input.PaymentStatus,
		OrderDate:     orderDate,
		TotalPrice:    input.TotalPrice,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	// Save to database
	config.DB.Create(&order)

	// Respond with the created order
	c.JSON(http.StatusOK, order)
}

// Get all Orders
func GetOrders(c *gin.Context) {
	var orders []models.Order
	config.DB.Find(&orders)
	c.JSON(http.StatusOK, orders)
}

// Get Order by ID
func GetOrderByID(c *gin.Context) {
	var order models.Order
	if err := config.DB.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

// Update an Order by ID
func UpdateOrder(c *gin.Context) {
	var order models.Order
	if err := config.DB.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	var input struct {
		IDUser        int     `json:"id_user" binding:"required"`
		IDPacket      int     `json:"id_packet" binding:"required"`
		PaymentStatus string  `json:"payment_status" binding:"required"`
		OrderDate     string  `json:"order_date" binding:"required"`
		TotalPrice    float64 `json:"total_price" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse date
	orderDate, err := time.Parse("2006-01-02", input.OrderDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	// Update fields
	order.IDUser = input.IDUser
	order.IDPacket = input.IDPacket
	order.PaymentStatus = input.PaymentStatus
	order.OrderDate = orderDate
	order.TotalPrice = input.TotalPrice
	order.UpdatedAt = time.Now()

	config.DB.Save(&order)
	c.JSON(http.StatusOK, order)
}

// Delete an Order by ID
func DeleteOrder(c *gin.Context) {
	var order models.Order
	if err := config.DB.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	config.DB.Delete(&order)
	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
