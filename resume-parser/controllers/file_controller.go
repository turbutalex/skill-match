package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"resume-parser/database"
	"resume-parser/models"
	"resume-parser/parser"
	"time"
)

func UploadResume(c *gin.Context) {
	content := c.GetHeader("Content-Type")
	boundary := c.GetHeader("Boundary")
	print(content, boundary)

	file, err := c.FormFile("resume")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	filePath := filepath.Join("uploads", file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	text, parseResult, err := parser.Extract(filePath)

	resume := models.Resume{
		FileName:   file.Filename,
		RawText:    text,
		ParsedName: parseResult.Name,
		Email:      parseResult.Email,
		Phone:      parseResult.Phone,
		Skills:     parseResult.Skills,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, insertErr := database.ResumeCollection.InsertOne(ctx, resume)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save resume"})
		return
	}

	c.JSON(http.StatusOK, resume)
}
