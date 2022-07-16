package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/UTDNebula/nebula-api/api/controllers"
)

func main() {
	router := gin.Default()

	// subroutes
	courseGroup := router.Group("/course")
	degreeGroup := router.Group("/degree")
	examGroup := router.Group("/exam")
	professorGroup := router.Group("/professor")
	sectionGroup := router.Group("/section")

	// route controllers
	courseGroup.GET("/", controllers.CourseSearch)
	courseGroup.GET("/:id", controllers.CourseById)

	examGroup.GET("/", controllers.ExamSearch)
	examGroup.GET("/:id", controllers.ExamById)

	degreeGroup.GET("/", controllers.DegreeSearch)
	degreeGroup.GET("/:id", controllers.DegreeById)

	professorGroup.GET("/", controllers.ProfessorSearch)
	professorGroup.GET("/:id", controllers.ProfessorById)

	sectionGroup.GET("/", controllers.SectionSearch)
	sectionGroup.GET("/:id", controllers.SectionById)

	// @DEBUG
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	http.ListenAndServe(":8080", router)
}
