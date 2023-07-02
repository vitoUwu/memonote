package models

import (
	"time"

	uuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string    `gorm:"primaryKey;uuid" json:"id"`
	Username  string    `gorm:"unique" json:"username"`
	Password  string    `gorm:"not null" json:"-"`
	Notes     []Note    `gorm:"foreignKey:UserID" json:"notes"`
	CreatedAt time.Time `gorm:"autoCreateTime;->" json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) BeforeCreate(db *gorm.DB) (err error) {
	db.Statement.SetColumn("id", uuid.New().String())
	return
}
