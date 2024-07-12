package routes

import (
	"net/http"

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

	err = user.Save()
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(201, gin.H{"status": "User created successfully", "data": user})

}

func LoginUser(context *gin.Context) {
	var user models.User
	err := context.BindJSON(&user)

	if err != nil {
		context.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	// check email and password
	dbUser, err := models.GetUser(user.Email)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	isPasswordValid := utils.ComparePassword(dbUser.Password, user.Password)

	if !isPasswordValid {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// generate token

	token := utils.GenerateToken(user.Email, int64(dbUser.ID))
	if token == "" {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
	}

	context.JSON(http.StatusOK, gin.H{"token": token})
	
}
