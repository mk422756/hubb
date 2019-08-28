package repository

import (
	"github.com/hubbdevelopers/db"
	"github.com/jinzhu/gorm"
)

type Page interface {
	Find() ([]*db.Page, error)
	FindById(id uint) (*db.Page, error)
	FindByUser(obj *db.User) ([]*db.Page, error)
	FindByTagId(id uint) ([]*db.Page, error)
	Create(page *db.Page) error
	Update(page *db.Page) error
	Delete(page *db.Page) error
	Associate(page *db.Page, tags []*db.Tag) error
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
	if result := r.DB.Order("created_at desc").Find(&pages); result.Error != nil {
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
	if result := r.DB.Model(obj).Order("created_at desc").Related(&pages); result.Error != nil {
		return nil, result.Error
	}
	return pages, nil
}

func (r *PageRepository) FindByTagId(id uint) ([]*db.Page, error) {
	tag := db.Tag{}
	if result := r.DB.First(&tag, id); result.Error != nil {
		return nil, result.Error
	}
	pages := []*db.Page{}

	r.DB.Model(&tag).Order("created_at desc").Association("Pages").Find(&pages)

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

func (r *PageRepository) Associate(page *db.Page, tags []*db.Tag) error {
	result := r.DB.Model(page).Association("Tags").Replace(tags)
	return result.Error
}

func (r *PageRepository) ClearAssociate(page *db.Page, tags []*db.Tag) error {
	result := r.DB.Model(page).Association("Tags").Append(tags)
	return result.Error
}
