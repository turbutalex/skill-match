package handlers

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

const authServiceURL = "http://localhost:8081"

func Login(c *gin.Context) {
	forwardRequest(c, "/api/login")
}

func Register(c *gin.Context) {
	forwardRequest(c, "/api/register")
}

func Profile(c *gin.Context) {
	forwardRequest(c, "/api/profile")
}

func forwardRequest(c *gin.Context, path string) {

	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	resp, err := http.Post(authServiceURL+path, "application/json", bytes.NewBuffer(bodyBytes))
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "auth service unreachable"})
		return
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	c.Data(resp.StatusCode, "application/json", respBody)
}
