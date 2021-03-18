package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"main/CRUD/models"
	"net/http"
	"time"
)

type CreateLearnerInput struct {
	FirstName 	string 		`json:firstName`
	LastName 	string 		`json:lastname`
	Email 		string 		`json:email`
	BirthDate 	time.Time 	`json:birthDate`
	Country 	string 		`json:country`
	Grade       string 		`json:grade`
	Institution string 		`json:Institution`
}

type UpdateLearnerInput struct {
	FirstName 	string 		`json:firstName`
	LastName 	string 		`json:lastname`
	Email 		string 		`json:email`
	BirthDate 	time.Time 	`json:birthDate`
	Country 	string 		`json:country`
	Grade       string 		`json:grade`
	Institution string 		`json:Institution`
}


// Get all learners
func FindLearners(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var learners []models.Learner
	db.Find(&learners)
	c.JSON(http.StatusOK, gin.H{"data": learners,"message":"ok"})
}

// Create new learner
func CreateLearner(c *gin.Context) {
	// Validate input
	var input CreateLearnerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	learner := models.Learner{FirstName: input.FirstName, LastName:input.LastName, Email:input.Email, BirthDate:input.BirthDate, Country:input.Country,Grade:input.Grade,Institution: input.Institution}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&learner)
	c.JSON(http.StatusOK, gin.H{"data": learner})
}

// Find a learner
func FindLearner(c *gin.Context) { // Get model if exist
	var learner models.Learner
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&learner).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": learner})
}
// Update a learner
func UpdateLearner(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var learner models.Learner
	if err := db.Where("id = ?", c.Param("id")).First(&learner).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// Validate input
	var input UpdateLearnerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&learner).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": learner})
}

func DeleteLearner(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var learner models.Learner
	if err := db.Where("id = ?", c.Param("id")).First(&learner).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	db.Delete(&learner)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
