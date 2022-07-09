package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func CourseSearch(w http.ResponseWriter, r *http.Request) {
	urlParams := r.URL.Query()

	fmt.Println(urlParams)
	fmt.Fprintf(w, "You're searching for the course by query\n")
}

func CourseById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Println(vars)
	fmt.Fprintf(w, "You're searching for the course by id %s\n", vars["id"])
}
