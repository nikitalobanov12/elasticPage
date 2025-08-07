package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"content-service/internal/database"
	"content-service/internal/services"
	"shared/models"
)

type UploadHandler struct {
	db        database.GormService
	s3Service services.S3Service
}

func NewUploadHandler(db database.GormService, s3Service services.S3Service) *UploadHandler {
	return &UploadHandler{
		db:        db,
		s3Service: s3Service,
	}
}

func (h *UploadHandler) UploadTextbook(c *gin.Context) {
	// Get form data
	title := c.PostForm("title")
	description := c.PostForm("description")
	userIDStr := c.PostForm("user_id")

	if title == "" || userIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title and user_id are required"})
		return
	}

	// Validate user_id format
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id format"})
		return
	}

	// Get uploaded file
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	// Upload file to S3
	filePath, err := h.s3Service.UploadFile(file, header, userIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create textbook record
	textbook := models.Textbook{
		Title:       title,
		Description: description,
		FilePath:    filePath,
		UserID:      userID,
	}

	if err := h.db.GetDB().Create(&textbook).Error; err != nil {
		// If database creation fails, try to clean up the uploaded file
		h.s3Service.DeleteFile(filePath)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create textbook record"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Textbook uploaded successfully",
		"textbook": textbook,
	})
}
