package router

import (
	"api-gateway/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Health check
	r.GET("/health", controllers.HealthCheck)

	// Placeholder routes
	api := r.Group("/api")
	{
		api.POST("/login", controllers.Login)
		api.POST("/register", controllers.Register)
		api.GET("/profile", controllers.Profile)
		api.GET("/me", controllers.Me)
	}
}
