package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
}

var persons = []Person{
	{Id: "0", Name: "rasul", Surname: "ramixanov", Age: 20},
	{Id: "3", Name: "roshka", Surname: "roshoo", Age: 20},
}

func GetUser(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, persons)
}

func personById(c *gin.Context) {
	id := c.Param("id")
	person, err := getUserById(id)

	if err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, person)
}

func getUserById(id string) (*Person, error) {
	for k, v := range persons {
		if v.Id == id {
			return &persons[k], nil

		}
	}
	return nil, errors.New("Person not found")

}

func CreateUser(c *gin.Context) {
	var newPerson Person
	if err := c.BindJSON(&newPerson); err != nil {
		return
	}
	persons = append(persons, newPerson)
	c.IndentedJSON(http.StatusCreated, persons)

}

func main() {
	router := gin.Default()
	router.GET("/persons", GetUser)
	router.GET("/persons/:id", personById)
	router.POST("/persons", CreateUser)
	router.Run("localhost:8080")

}
