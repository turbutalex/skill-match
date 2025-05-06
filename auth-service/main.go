package main

import (
	"auth-service/database"
	"auth-service/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	database.InitDB()
	r := gin.Default()
	r.Use(cors.Default())

	router.SetupRoutes(r)

	log.Println("Starting Auth Service on :8081")
	if err := r.Run(":8081"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
