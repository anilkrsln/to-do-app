package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ToDo struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	UserID      uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `json:"description"`
	Completed   bool      `gorm:"default:false" json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
}

func (todo *ToDo) BeforeCreate(tx *gorm.DB) (err error) {
	todo.ID = uuid.New()
	return
}
