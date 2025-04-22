package router

import (
	"api-gateway/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Health check
	r.GET("/health", handlers.HealthCheck)

	// Placeholder routes
	api := r.Group("/api")
	{
		api.GET("/jobs", handlers.GetJobs)
		api.POST("/login", handlers.Login)
		api.POST("/register", handlers.Register)
	}
}
