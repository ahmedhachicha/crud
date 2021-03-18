package routes

import (



	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"main/CRUD/controllers"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET("/tasks", controllers.FindTasks)
	r.POST("/tasks", controllers.CreateTask)
	r.GET("/tasks/:id", controllers.FindTask)
	r.PATCH("/tasks/:id", controllers.UpdateTask)
	r.DELETE("tasks/:id", controllers.DeleteTask)

	r.GET("/tutor", controllers.FindTutors)
	r.POST("/tutor", controllers.CreateTutor)
	r.GET("/tutor/:id", controllers.FindTutor)
	r.PATCH("/tutor/:id", controllers.UpdateTutor)
	r.DELETE("tutor/:id", controllers.DeleteTutor)

	r.GET("/learner", controllers.FindLearners)
	r.POST("/learner", controllers.CreateLearner)
	r.GET("/learner/:id", controllers.FindLearner)
	r.PATCH("/learner/:id", controllers.UpdateLearner)
	r.DELETE("learner/:id", controllers.DeleteLearner)

	r.GET("/course", controllers.FindCourses)
	r.POST("/course", controllers.CreateCourse)
	r.GET("/course/:id", controllers.FindCourse)
	r.PATCH("/course/:id", controllers.UpdateCourse)
	r.DELETE("course/:id", controllers.DeleteCourse)
	return r
}
