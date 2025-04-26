package router

import (
	"auth-service/controllers"
	"auth-service/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	api := r.Group("/api")
	{
		api.POST("/login", controllers.Login)
		api.POST("/register", controllers.Register)

		api.Use(middleware.AuthMiddleware())
		api.GET("/profile", controllers.Profile)
		api.GET("/me", controllers.GetCurrentUser)
	}
}
