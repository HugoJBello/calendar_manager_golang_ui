package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"-"`
	Username  string    `gorm:"uniqueIndex" json:"-"`
	Email     string    `gorm:"unique" json:"-"`
	Password  string    `json:"-"`
	Bio       string    `json:"-"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"-"`
}
