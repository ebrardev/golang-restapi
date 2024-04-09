package main

import (
	"net/http"

	"ebrarcode.dev/restapi-go/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	// server.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)

}

func createEvent(context *gin.Context) {
	var event models.Event
	context.ShouldBindJSON(&event)
}
