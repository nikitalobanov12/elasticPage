package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"content-service/internal/handlers"
	"content-service/internal/services"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:5173"}, // Add your frontend URLs
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	// Initialize services and handlers
	s3Service := services.NewS3Service()
	textbookHandler := handlers.NewTextbookHandler(s.gormDB)
	uploadHandler := handlers.NewUploadHandler(s.gormDB, s3Service)

	// Basic routes
	r.GET("/", s.HelloWorldHandler)
	r.GET("/health", s.healthHandler)

	// API routes
	api := r.Group("/api/v1")
	{
		// Textbook routes
		api.GET("/textbooks", textbookHandler.GetTextbooks)
		api.POST("/textbooks", textbookHandler.CreateTextbook)
		api.GET("/textbooks/:id", textbookHandler.GetTextbook)
		api.PUT("/textbooks/:id", textbookHandler.UpdateTextbook)
		api.DELETE("/textbooks/:id", textbookHandler.DeleteTextbook)

		// Upload routes
		api.POST("/upload", uploadHandler.UploadTextbook)
	}

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}
