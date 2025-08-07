package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"shared/models"
)

type GormService interface {
	GetDB() *gorm.DB
	AutoMigrate() error
	Close() error
	Health() map[string]string
}

type gormService struct {
	db *gorm.DB
}

var gormInstance *gormService

func NewGorm() GormService {
	if gormInstance != nil {
		return gormInstance
	}

	database := os.Getenv("BLUEPRINT_DB_DATABASE")
	password := os.Getenv("BLUEPRINT_DB_PASSWORD")
	username := os.Getenv("BLUEPRINT_DB_USERNAME")
	port := os.Getenv("BLUEPRINT_DB_PORT")
	host := os.Getenv("BLUEPRINT_DB_HOST")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		host, username, password, database, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	gormInstance = &gormService{db: db}
	return gormInstance
}

func (g *gormService) GetDB() *gorm.DB {
	return g.db
}

func (g *gormService) AutoMigrate() error {
	return g.db.AutoMigrate(
		&models.User{},
		&models.Textbook{},
	)
}

func (g *gormService) Close() error {
	sqlDB, err := g.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (g *gormService) Health() map[string]string {
	stats := make(map[string]string)

	sqlDB, err := g.db.DB()
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("failed to get underlying sql.DB: %v", err)
		return stats
	}

	err = sqlDB.Ping()
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db ping failed: %v", err)
		return stats
	}

	stats["status"] = "up"
	stats["message"] = "Database is healthy"
	return stats
}
