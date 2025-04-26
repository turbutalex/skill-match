package controllers

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

const authServiceURL = "http://localhost:8081"

func Login(c *gin.Context) {
	forwardPOSTRequest(c, "/api/login")
}

func Register(c *gin.Context) {
	forwardPOSTRequest(c, "/api/register")
}

func Profile(c *gin.Context) {
	forwardPOSTRequest(c, "/api/profile")
}

func Me(c *gin.Context) {
	forwardGETRequest(c, "/api/me")
}

func forwardPOSTRequest(c *gin.Context, path string) {

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

func forwardGETRequest(c *gin.Context, path string) {
	req, err := http.NewRequest("GET", authServiceURL+path, nil)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Encountered error while building request " + err.Error()})
		return
	}

	req.Header.Add("Authorization", c.GetHeader("Authorization"))

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "auth service unreachable " + err.Error()})
		return
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	c.Data(resp.StatusCode, "application/json", respBody)
}
