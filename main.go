package main

import (
	"myapp/controllers"
	"myapp/middlewares"
	"myapp/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()
	r.Use(middlewares.M1)
	// Serve static files from the "static" directory
	r.Static("/static", "./static")

	// Route for serving the index.html file
	r.GET("/", controllers.GET)
	r.Use(middlewares.M2)
	r.POST("/", controllers.POST)
	r.Use(middlewares.M3)
	r.GET("/auth/google", controllers.Login)
	r.Use(middlewares.M1)
	r.GET("/auth/google/callback", controllers.Callback)

	r.GET("/login", controllers.GETLogin)
	r.POST("/login", controllers.POSTLogin)

	// Start the server
	r.Run(":8080")
}
