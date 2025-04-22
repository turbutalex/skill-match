package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetJobs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "List of jobs (placeholder)",
	})
}

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Login route (placeholder)",
	})
}

func Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Register route (placeholder)",
	})
}
