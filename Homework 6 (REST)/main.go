package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type user struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Age      int64  `json:"age"`
}

var users = []user{
	{ID: "1", Username: "almas29@gmail.com", Password: "almas2304", Name: "Almas", Surname: "Ermyrzaev", Age: 19},
	{ID: "2", Username: "admin@gmail.com", Password: "admin", Name: "Admin", Surname: "Adminov", Age: 20},
	{ID: "3", Username: "qwerty@gmail.com", Password: "qwerty", Name: "Qwerty", Surname: "Qwerty", Age: 21},
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func createUser(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func getUserByID(c *gin.Context) {
	id := c.Param("id")

	for _, user := range users {
		if user.ID == id {
			c.IndentedJSON(http.StatusOK, user)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user with ID=" + id + " not found"})
}

func deleteUserByID(c *gin.Context) {
	id := c.Param("id")

	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.POST("/users", createUser)
	router.GET("/users/:id", getUserByID)
	router.DELETE("/users/:id", deleteUserByID)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
