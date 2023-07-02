package models

import (
	"time"

	uuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type Note struct {
	ID        string    `gorm:"primaryKey;uuid" json:"id"`
	Content   string    `gorm:"not null" json:"content"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `gorm:"autoCreateTime;->" json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *Note) BeforeCreate(db *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}
