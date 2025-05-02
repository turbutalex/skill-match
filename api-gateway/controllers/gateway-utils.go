package controllers

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

const authServiceURL = "http://localhost:8081"
const resumeServiceURL = "http://localhost:8082"

func forwardPOSTRequest(c *gin.Context, path string, baseURL string, contentType string) {

	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	resp, err := http.Post(baseURL+path, contentType, bytes.NewBuffer(bodyBytes))
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "auth service unreachable"})
		return
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	c.Data(resp.StatusCode, "application/json", respBody)
}

func forwardGETRequest(c *gin.Context, path string, baseURL string) {
	req, err := http.NewRequest("GET", baseURL+path, nil)
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
