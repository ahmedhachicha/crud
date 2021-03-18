package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"main/CRUD/models"
	"net/http"
	"time"
)

type CreateTutorInput struct {
	FirstName 	string 		`json:firstName`
	LastName 	string 		`json:lastname`
	Email 		string 		`json:email`
	BirthDate 	time.Time 	`json:birthDate`
	Country 	string 		`json:country`
	Biography   string      `json : biography`
	Courses []	models.Course  `gorm:"foreignKey:UserRefer"`

}

type UpdateTutorInput struct {
	FirstName 	string 		`json:firstName`
	LastName 	string 		`json:lastname`
	Email 		string 		`json:email`
	BirthDate 	time.Time 	`json:birthDate`
	Country 	string 		`json:country`
	Biography   string      `json : biography`

}


// Get all tutors
func FindTutors(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var tutors []models.Tutor
	db.Find(&tutors)
	c.JSON(http.StatusOK, gin.H{"data": tutors,"message":"ok"})
}

// Create new tutor
func CreateTutor(c *gin.Context) {
	// Validate input
	var input CreateTutorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tutor := models.Tutor{FirstName: input.FirstName, LastName:input.LastName, Email:input.Email, BirthDate:input.BirthDate, Country:input.Country,Biography:input.Biography,Courses: input.Courses}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&tutor)
	c.JSON(http.StatusOK, gin.H{"data": tutor})
}

// Find a tutor
func FindTutor(c *gin.Context) { // Get model if exist
	var tutor models.Tutor
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&tutor).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": tutor})
}
// Update a tutor
func UpdateTutor(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var tutor models.Tutor
	if err := db.Where("id = ?", c.Param("id")).First(&tutor).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// Validate input<
	var input UpdateTutorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&tutor).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": tutor})
}

func DeleteTutor(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var tutor models.Tutor
	if err := db.Where("id = ?", c.Param("id")).First(&tutor).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	db.Delete(&tutor)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
