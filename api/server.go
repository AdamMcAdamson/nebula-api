package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/UTDNebula/nebula-api/api/controllers"
)

func main() {
	router := mux.NewRouter()
	courseRouter := router.PathPrefix("/course").Subrouter()

	courseRouter.HandleFunc("/", controllers.CourseSearch)
	courseRouter.HandleFunc("/{id}", controllers.CourseById)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	http.ListenAndServe(":80", router)
}
