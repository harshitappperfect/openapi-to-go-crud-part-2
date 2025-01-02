package main

import (
	"log"
	"sample/db"
	"sample/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()

	router := gin.Default()

	router.GET("/items", handlers.GetItems)
	router.POST("/items", handlers.CreateItem)

	// Add routes for other handlers here

	log.Fatal(router.Run(":8080"))
}
