package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Textbook struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `json:"description"`
	FilePath    string    `gorm:"not null" json:"file_path"`
	UserID      uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	User        User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (t *Textbook) BeforeCreate(tx *gorm.DB) error {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return nil
}
