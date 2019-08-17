package repository

import (
	"github.com/hubbdevelopers/db"
	"github.com/jinzhu/gorm"
)

type Tag interface {
	Find() ([]*db.Tag, error)
	FindById(id uint) (*db.Tag, error)
	FindByName(name string) (*db.Tag, error)
	FindByPageId(id uint) ([]*db.Tag, error)
	Create(Tag *db.Tag) error
	Delete(Tag *db.Tag) error
}

type TagRepository struct {
	*gorm.DB
}

func NewTagRepository(db *gorm.DB) Tag {
	newRepo := &TagRepository{
		db,
	}
	return newRepo
}

func (r *TagRepository) Find() ([]*db.Tag, error) {
	tags := []*db.Tag{}
	if result := r.DB.Find(&tags); result.Error != nil {
		return nil, result.Error
	}
	return tags, nil
}

func (r *TagRepository) FindById(id uint) (*db.Tag, error) {
	tag := db.Tag{}
	if result := r.DB.First(&tag, id); result.Error != nil {
		return nil, result.Error
	}
	return &tag, nil
}

func (r *TagRepository) FindByName(name string) (*db.Tag, error) {
	tag := db.Tag{}
	if result := r.DB.Where("name = ?", name).First(&tag); result.Error != nil {
		return nil, result.Error
	}
	return &tag, nil
}

func (r *TagRepository) FindByPageId(id uint) ([]*db.Tag, error) {
	page := db.Page{}
	if result := r.DB.First(&page, id); result.Error != nil {
		return nil, result.Error
	}
	tags := []*db.Tag{}
	
	r.DB.Model(&page).Association("Tags").Find(&tags)

	return tags, nil
}

func (r *TagRepository) Create(tag *db.Tag) error {
	if result := r.DB.Create(tag); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *TagRepository) Delete(tag *db.Tag) error {
	if result := r.DB.Delete(tag); result.Error != nil {
		return result.Error
	}
	return nil
}
