package db

import (
	"time"
)

type User struct {
	ID          uint `gorm:"primary_key"`
	UID         string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
	Name        string
	AccountID   string
	Image       string
	Description string
	Pages       []Page `gorm:"foreignkey:UserID"`
}

type Page struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	UserID    uint
	Name      string
	Text      string
	Image     string
}
