package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"resume-parser/database"
	"resume-parser/router"
)

func main() {
	database.Connect()
	r := gin.Default()
	r.Use(cors.Default())
	router.SetupRoutes(r)

	log.Println("Starting Parser Service on :8082")
	if err := r.Run(":8082"); err != nil {
		log.Fatal("Server failed to start:", err)
	}

}
