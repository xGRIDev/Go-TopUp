package routes

import "github.com/gin-gonic/gin"

func TopupRegistRoutes(server *gin.Engine) {
	server.GET("/top-up", getTopUp)
	server.POST("/top-up", createTopUp)
	server.GET("/top-up/:id", getTopUpID)
	server.PUT("/top-up/:id", updateItemTopUp)
	server.DELETE("/top-up/:id", deleteItemTopUp)
}
