package repository

import (
	"github.com/hubbdevelopers/db"
	"github.com/jinzhu/gorm"
)

type Page interface {
	Find() ([]*db.Page, error)
	FindById(id uint) (*db.Page, error)
	FindByUser(obj *db.User) ([]*db.Page, error)
	Create(page *db.Page) error
	Update(page *db.Page) error
	Delete(page *db.Page) error
}

type PageRepository struct {
	*gorm.DB
}

func NewPageRepository(db *gorm.DB) Page {
	newRepo := &PageRepository{
		db,
	}
	return newRepo
}

func (r *PageRepository) Find() ([]*db.Page, error) {
	pages := []*db.Page{}
	if result := r.DB.Find(&pages); result.Error != nil {
		return nil, result.Error
	}
	return pages, nil
}

func (r *PageRepository) FindById(id uint) (*db.Page, error) {
	page := db.Page{}
	if result := r.DB.First(&page, id); result.Error != nil {
		return nil, result.Error
	}
	return &page, nil
}

func (r *PageRepository) FindByUser(obj *db.User) ([]*db.Page, error) {
	pages := []*db.Page{}
	if result := r.DB.Model(obj).Related(&pages); result.Error != nil {
		return nil, result.Error
	}
	return pages, nil
}

func (r *PageRepository) Create(page *db.Page) error {
	if result := r.DB.Create(page); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *PageRepository) Update(page *db.Page) error {
	if result := r.DB.Save(page); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *PageRepository) Delete(page *db.Page) error {
	if result := r.DB.Delete(page); result.Error != nil {
		return result.Error
	}
	return nil
}
