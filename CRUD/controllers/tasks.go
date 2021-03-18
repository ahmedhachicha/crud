package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"main/CRUD/models"
	"net/http"
	"time"
)

type CreateTaskInput struct {
	FirstName string `json:firstName`
	LastName string `json:lastname`
	Email string 	`json:email`
	BirthDate time.Time 	`json:birthDate`
	Country string 	`json:country`
}

type UpdateTaskInput struct {
	FirstName string `json:firstName`
	LastName string `json:lastname`
	Email string 	`json:email`
	BirthDate time.Time 	`json:birthDate`
	Country string 	`json:country`
}


// Get all tasks
func FindTasks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var users []models.User
	db.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users,"message":"ok"})
}

// Create new task
func CreateTask(c *gin.Context) {
	// Validate input
	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{FirstName: input.FirstName, LastName:input.LastName, Email:input.Email, BirthDate:input.BirthDate, Country:input.Country }
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// Find a task
func FindTask(c *gin.Context) { // Get model if exist
	var task models.User
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": task})
}
// Update a task
func UpdateTask(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var user models.User
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// Validate input
	var input UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&user).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteTask(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var user models.User
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	db.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
