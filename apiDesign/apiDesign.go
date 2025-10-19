package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type dog struct {
	Breed string `json:"breed"`
	Age	  int `json:"age"`
	Name string `json:"name"`
}

var dogs = []dog{
	{Breed: "Labrador", Age: 8, Name: "Bobby"},
	{Breed: "Chihuahua", Age: 16, Name: "Lola"},
	{Breed: "Dachshund", Age: 2, Name: "Slinky"},
}

func getDogs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, dogs)
}

func main(){
	router := gin.Default()
	router.GET("/dogs", getDogs)
	router.Run("localhost:8080")
}