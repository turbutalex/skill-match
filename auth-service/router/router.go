package router

import (
	"auth-service/handlers"
	"auth-service/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	api := r.Group("/api")
	{
		api.POST("/login", handlers.Login)
		api.POST("/register", handlers.Register)

		api.Use(middleware.AuthMiddleware())
		api.GET("/profile", handlers.Profile)
	}
}
