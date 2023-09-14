package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Failed to load environment variable")
	}
}
func main() {
	port := os.Getenv("PORT")

	engine := gin.New()
	engine.Use(gin.Logger())
	log.Println(port)
	if err := engine.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
