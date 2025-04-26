package router

import (
	"github.com/gin-gonic/gin"
	"resume-parser/controllers"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/resume", controllers.UploadResume)
	}
}
