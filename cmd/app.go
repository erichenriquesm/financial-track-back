package main

import (
	"financial-track/database"
	"financial-track/middleware"
	"financial-track/route"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ File .env not found, using environment variables")
	}

	database.Connect()
	database.Migrate()

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "81"
	}

	server := gin.Default()

	route.RegisterHealthRoutes(server)
	route.RegisterUserRoutes(server)

	auth := server.Group("/")
	auth.Use(middleware.AuthMiddleware())

	// Authenticated routes
	route.RegisterExpenseRoutes(auth)

	server.Run(":" + port)
}
