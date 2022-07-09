package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func ExamSearch(w http.ResponseWriter, r *http.Request) {
	urlParams := r.URL.Query()

	fmt.Println(urlParams)
	fmt.Fprintf(w, "You're searching for the exam by query\n")
}

func ExamById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Println(vars)
	fmt.Fprintf(w, "You're searching for the exam by id %s\n", vars["id"])
}
