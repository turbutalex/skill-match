package controllers

import (
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	forwardPOSTRequest(c, "/api/login", authServiceURL, "application/json")
}

func Register(c *gin.Context) {
	forwardPOSTRequest(c, "/api/register", authServiceURL, "application/json")
}

func Profile(c *gin.Context) {
	forwardPOSTRequest(c, "/api/profile", authServiceURL, "application/json")
}

func Me(c *gin.Context) {
	forwardGETRequest(c, "/api/me", authServiceURL)
}
