package main

import (
	"api-gateway/router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	router.SetupRoutes(r)

	log.Println("Starting API Gateway on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
