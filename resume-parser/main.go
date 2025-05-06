package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/unidoc/unipdf/v3/common/license"
	"log"
	"resume-parser/database"
	"resume-parser/router"
	"resume-parser/utils"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	licenceKey := utils.GetEnv(utils.PDF_APIKEY, "defaultKey")
	err = license.SetMeteredKey(licenceKey)
	if err != nil {
		panic(err)
	}

	database.Connect()
	r := gin.Default()
	r.Use(cors.Default())
	router.SetupRoutes(r)

	log.Println("Starting Parser Service on :8082")
	if err := r.Run(":8082"); err != nil {
		log.Fatal("Server failed to start:", err)
	}

}
