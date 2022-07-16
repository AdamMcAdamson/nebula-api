package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SectionSearch(c *gin.Context) {
	name := c.Query("name")              // value of specific query parameter: string
	queryParams := c.Request.URL.Query() // map of all query params: map[string][]string

	fmt.Println(queryParams)
	message := "You're searching for the section by query parameters. For instance the name is: " + name
	c.String(http.StatusOK, message)
}

func SectionById(c *gin.Context) {
	id := c.Param("id") // value of specfic URL parameter: string

	fmt.Println(id)
	message := "You're searching for the section by id: " + id
	c.String(http.StatusOK, message)
}
