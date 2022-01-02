package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	Fname string `json:"first name"`
	Lname string `json:"last name"`
}

var users = []user{
	{Fname: "aydin", Lname: "nemati"},
	{Fname: "khashy", Lname: "mosavi"},
	{Fname: "ali", Lname: "tavafi"},
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)

	router.Run("localhost:3000")
}
