package routes

import (
	"net/http"
	"strconv"

	utils "example.com/topup-restapi/Utils"
	M_topups "example.com/topup-restapi/models"
	"github.com/gin-gonic/gin"
)

func getTopUp(context *gin.Context) {
	//topups := M_topup.GetAllItem()
	topups, err := M_topups.GetAllItem()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fethcing the Data. Please Try Again"})
	}
	//	context.JSON(http.StatusOK, topups)
	context.JSON(http.StatusOK, topups)
}

func createTopUp(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		return
	}

	err := utils.VerifToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		return
	}

	var topup M_topups.TopUp
	err = context.ShouldBindJSON(&topup)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't Parse Request The Data"})
		return
	}

	topup.ID = 1
	topup.UserID = 1

	err = topup.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not Create Topup. Please Try Again"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Create Topup", "topup": topup})

}

func getTopUpID(context *gin.Context) {
	TopupID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not Parse Topup-ID"})
		return
	}

	topup, err := M_topups.GetTopUpByID(TopupID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could Fetch Topup"})
		return
	}

	context.JSON(http.StatusOK, topup)
}

func updateItemTopUp(context *gin.Context) {
	topupID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not Parse Topup-ID"})
		return
	}
	_, err = M_topups.GetTopUpByID(topupID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the data"})
		return
	}

	var TopupUpdated M_topups.TopUp
	err = context.ShouldBindJSON(&TopupUpdated)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event Updated Successfully"})

}

func deleteItemTopUp(context *gin.Context) {

	TopupID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not Parse Topup-ID"})
		return
	}

	Topup, err := M_topups.GetTopUpByID(TopupID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the data"})
		return
	}

	err = Topup.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the data"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Deleted Item Successfully"})

}
