package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"content-service/internal/database"
)

type Server struct {
	port int

	db     database.Service
	gormDB database.GormService
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	gormDB := database.NewGorm()
	gormDB.AutoMigrate() // Run migrations

	NewServer := &Server{
		port: port,

		db:     database.New(),
		gormDB: gormDB,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
