package main

import (
	"auth-service/router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	router.SetupRoutes(r)

	log.Println("Starting Auth Service on :8081")
	if err := r.Run(":8081"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
