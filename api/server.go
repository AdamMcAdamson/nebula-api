package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/UTDNebula/nebula-api/api/controllers"
)

func main() {
	router := mux.NewRouter()

	// subroutes
	courseRouter := router.PathPrefix("/course").Subrouter()
	degreeRouter := router.PathPrefix("/degree").Subrouter()
	examRouter := router.PathPrefix("/exam").Subrouter()
	professorRouter := router.PathPrefix("/professor").Subrouter()
	sectionRouter := router.PathPrefix("/section").Subrouter()

	// route controllers
	courseRouter.HandleFunc("/", controllers.CourseSearch)
	courseRouter.HandleFunc("/{id}", controllers.CourseById)

	examRouter.HandleFunc("/", controllers.ExamSearch)
	examRouter.HandleFunc("/{id}", controllers.ExamById)

	degreeRouter.HandleFunc("/", controllers.DegreeSearch)
	degreeRouter.HandleFunc("/{id}", controllers.DegreeById)

	professorRouter.HandleFunc("/", controllers.ProfessorSearch)
	professorRouter.HandleFunc("/{id}", controllers.ProfessorById)

	sectionRouter.HandleFunc("/", controllers.SectionSearch)
	sectionRouter.HandleFunc("/{id}", controllers.SectionById)

	// // @DEBUG
	// router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello World!")
	// })

	http.ListenAndServe(":80", router)
}
