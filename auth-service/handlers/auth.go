package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var req AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	// Dummy login logic
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful (dummy)",
		"user":    req.Email,
	})
}

func Register(c *gin.Context) {
	var req AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	// Dummy register logic
	c.JSON(http.StatusOK, gin.H{
		"message": "User registered (dummy)",
		"user":    req.Email,
	})
}
