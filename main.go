package main

import (
	"ebrarcode.dev/restapi-go/db"

	"ebrarcode.dev/restapi-go/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
