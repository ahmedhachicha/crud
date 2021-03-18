package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"main/CRUD/models"
	"net/http"
)

type CreateCourseInput struct {
	Title 		string `json:title`
	Level 		string `json:level`
	Description string `json:title`
	UserRefer   uint
}

type UpdateCourseInput struct {
	Title 		string `json:title`
	Level 		string `json:level`
	Description string `json:title`
	UserRefer   uint
}


// Get all Courses
func FindCourses(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var Courses []models.Course
	db.Find(&Courses)
	c.JSON(http.StatusOK, gin.H{"data": Courses,"message":"ok"})
}

// Create new Course
func CreateCourse(c *gin.Context) {
	// Validate input
	var input CreateCourseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Course := models.Course{Title: input.Title, Level:input.Level, Description:input.Description,UserRefer:input.UserRefer}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&Course)
	c.JSON(http.StatusOK, gin.H{"data": Course})
}

// Find a Course
func FindCourse(c *gin.Context) { // Get model if exist
	var Course models.Course
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&Course).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Course})
}
// Update a Course
func UpdateCourse(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var Course models.Course
	if err := db.Where("id = ?", c.Param("id")).First(&Course).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// Validate input
	var input UpdateCourseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&Course).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": Course})
}

func DeleteCourse(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var Course models.Course
	if err := db.Where("id = ?", c.Param("id")).First(&Course).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	db.Delete(&Course)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
