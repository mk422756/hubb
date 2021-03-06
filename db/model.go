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
	Pages       []Page
}

type Page struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	UserID    uint
	User      User
	Name      string
	Text      string `sql:"type:text"`
	Image     string
	Tags      []Tag `gorm:"many2many:page_tags;"`
}

type Tag struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Name      string
	Pages     []Page `gorm:"many2many:page_tags;"`
}
