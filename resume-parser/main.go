package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"resume-parser/router"
)

func main() {
	r := gin.Default()

	router.SetupRoutes(r)

	log.Println("Starting Parser Service on :8082")
	if err := r.Run(":8082"); err != nil {
		log.Fatal("Server failed to start:", err)
	}

}
