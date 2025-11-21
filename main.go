package main

import (
	db "example.com/topup-restapi/DB"
	routes "example.com/topup-restapi/Routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.DBinit()
	server := gin.Default()

	routes.TopupRegistRoutes(server)

	server.Run(":8020")
}
