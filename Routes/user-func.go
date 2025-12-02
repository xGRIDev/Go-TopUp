package routes

import (
	"net/http"

	utils "example.com/topup-restapi/Utils"
	"example.com/topup-restapi/models"
	"github.com/gin-gonic/gin"
)

func UserSignUp(context *gin.Context) {
	var user models.Users
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"mesage": "Cannot Fetch The Data"})
		return
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create User"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User Table Created."})
}

func UserSignIn(context *gin.Context) {
	var users models.Users

	err := context.ShouldBindJSON(&users)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could Not Parse Request The Data"})
	}

	err = users.LoginCredentialValidate()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could Not Authenticated User."})
		return
	}

	token, err := utils.JWTGenerate(users.Email, users.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could Not authenticated user."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Login Successfully.", "token": token})
}
