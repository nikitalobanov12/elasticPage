package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"content-service/internal/database"
	"shared/models"
)

type TextbookHandler struct {
	db database.GormService
}

func NewTextbookHandler(db database.GormService) *TextbookHandler {
	return &TextbookHandler{db: db}
}

func (h *TextbookHandler) CreateTextbook(c *gin.Context) {
	var textbook models.Textbook
	if err := c.ShouldBindJSON(&textbook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.db.GetDB().Create(&textbook).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create textbook"})
		return
	}

	c.JSON(http.StatusCreated, textbook)
}

func (h *TextbookHandler) GetTextbooks(c *gin.Context) {
	var textbooks []models.Textbook

	userID := c.Query("user_id")
	query := h.db.GetDB()

	if userID != "" {
		if _, err := uuid.Parse(userID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id format"})
			return
		}
		query = query.Where("user_id = ?", userID)
	}

	if err := query.Preload("User").Find(&textbooks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch textbooks"})
		return
	}

	c.JSON(http.StatusOK, textbooks)
}

func (h *TextbookHandler) GetTextbook(c *gin.Context) {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid textbook ID format"})
		return
	}

	var textbook models.Textbook
	if err := h.db.GetDB().Preload("User").First(&textbook, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Textbook not found"})
		return
	}

	c.JSON(http.StatusOK, textbook)
}

func (h *TextbookHandler) UpdateTextbook(c *gin.Context) {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid textbook ID format"})
		return
	}

	var textbook models.Textbook
	if err := h.db.GetDB().First(&textbook, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Textbook not found"})
		return
	}

	var updateData models.Textbook
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.db.GetDB().Model(&textbook).Updates(updateData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update textbook"})
		return
	}

	c.JSON(http.StatusOK, textbook)
}

func (h *TextbookHandler) DeleteTextbook(c *gin.Context) {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid textbook ID format"})
		return
	}

	if err := h.db.GetDB().Delete(&models.Textbook{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete textbook"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Textbook deleted successfully"})
}
