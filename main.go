package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Gender  string `json:"gender"`
	Address string `json:"address"`
}

var users = []User{
	{ID: "1", Name: "ankit", Age: 22, Gender: "Male", Address: "newAddress"},
	{ID: "2", Name: "ankita", Age: 21, Gender: "Female", Address: "otherAddress"},
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func getById(c *gin.Context) {
	id := c.Param("id")
	user, err := getUserById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
	}
	c.IndentedJSON(http.StatusOK, user)
}

func getUserById(id string) (*User, error) {
	for i, k := range users {
		if k.ID == id {
			return &users[i], nil
		}
	}
	return nil, errors.New("User not found!")

}

func createUser(c *gin.Context) {
	var newUser User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:id", getById)
	router.POST("/create", createUser)
	router.Run("localhost:8080")
}
