package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"practice11/models"
)

type CreateUserInput struct {
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type UpdateUserInput struct {
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
	Username string `json:"username" binding:"required"`
}

func GetAllUsers(context *gin.Context) {
	var users []models.User
	models.DB.Find(&users)
	context.JSON(http.StatusOK, gin.H{"users": users})
}

func GetUser(context *gin.Context) {
	var user models.User

	if err := models.DB.Where("id=?", context.Param("id")).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"user": user})
}

func CreateUser(context *gin.Context) {
	var input CreateUserInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Name: input.Name, Surname: input.Surname, Username: input.Username}
	models.DB.Create(&user)
	context.JSON(http.StatusOK, gin.H{"users": user})
}

func UpdateUser(context *gin.Context) {
	var user models.User

	if err := models.DB.Where("id=?", context.Param("id")).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	var input UpdateUserInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&user).Update(input)
	context.JSON(http.StatusOK, gin.H{"users": user})
}

func DeleteUser(context *gin.Context) {
	var user models.User

	if err := models.DB.Where("id=?", context.Param("id")).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	models.DB.Delete(&user)
	context.JSON(http.StatusOK, gin.H{"deleted": true})
}
