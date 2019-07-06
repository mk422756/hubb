package db

import (
	"time"
)

type User struct {
	ID          uint   `gorm:"primary_key"`
	UID         string `sql:"index"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
	Name        string     `sql:"index"`
	AccountID   string     `sql:"index"`
	Image       string
	Description string
	Twitter     string
	Instagram   string
	Facebook    string
	Homepage    string
	Birthday    time.Time `sql:"type:date" gorm:"default:'1970-01-01'"`
	Pages       []Page    `gorm:"foreignkey:UserID"`
}

type Page struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	UserID    uint
	Name      string
	Text      string `sql:"type:text"`
	Image     string
}
