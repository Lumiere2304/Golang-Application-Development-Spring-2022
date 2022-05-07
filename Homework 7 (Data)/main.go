package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"practice11/handlers"
	"practice11/models"
)

func main() {
	route := gin.Default()
	models.ConnectDB()

	route.GET("/user", handlers.GetAllUsers)
	route.GET("/user/:id", handlers.GetUser)
	route.POST("/user", handlers.CreateUser)
	route.PUT("/user/:id", handlers.UpdateUser)
	route.DELETE("/user/:id", handlers.DeleteUser)

	err := route.Run()
	if err != nil {
		return
	}
}
