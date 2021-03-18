package main

import (
	"main/CRUD/models"
	"main/CRUD/routes"
)

func main() {

	db := models.SetupDB()
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Learner{})
	db.AutoMigrate(&models.Tutor{})
	db.AutoMigrate(&models.Course{})

	r := routes.SetupRoutes(db)
	r.Run()
}
