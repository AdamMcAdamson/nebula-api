package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CourseSearch(c *gin.Context) {
	name := c.Query("name")              // value of specific query parameter: string
	queryParams := c.Request.URL.Query() // map of all query params: map[string][]string

	fmt.Println(queryParams)
	message := "You're searching for the course by query parameters. For instance the name is: " + name
	c.String(http.StatusOK, message)
}

func CourseById(c *gin.Context) {
	id := c.Param("id") // value of specfic URL parameter: string

	fmt.Println(id)
	message := "You're searching for the course by id: " + id
	c.String(http.StatusOK, message)
}
