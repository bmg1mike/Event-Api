package routes

import (
	"events.com/models"
	"events.com/utils"
	"github.com/gin-gonic/gin"
)

func CreateUser(context *gin.Context) {
	var user models.User

	err := context.BindJSON(&user)

	if err != nil {
		context.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	user.Password = utils.HashPassword(user.Password)
	err = user.Save()
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(201, gin.H{"status": "User created successfully", "data": user})

}
