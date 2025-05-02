package controllers

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func Upload(c *gin.Context) {
	forwardFileUpload(c, "/api/resume", resumeServiceURL, "multipart/form-data")
}

func forwardFileUpload(c *gin.Context, path string, baseURL string, contentType string) {

	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	req, err := http.NewRequest("POST", baseURL+path, bytes.NewBuffer(bodyBytes))
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Encountered error while building request " + err.Error()})
		return
	}

	req.Header.Add("Authorization", c.GetHeader("Authorization"))
	req.Header.Add("Content-Type", c.GetHeader("Content-Type"))

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
