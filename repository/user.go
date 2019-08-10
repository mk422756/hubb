package repository

import (
	"github.com/hubbdevelopers/db"
	"github.com/jinzhu/gorm"
)

type User interface {
	Find() ([]*db.User, error)
	FindById(id uint) (*db.User, error)
	FindByAccountId(accountID string) (*db.User, error)
	FindByUId(uid string) (*db.User, error)
	Create(user *db.User) error
	Update(user *db.User) error
	Delete(user *db.User) error
}

type UserRepository struct {
	*gorm.DB
}

func NewUserRepository(db *gorm.DB) User {
	newRepo := &UserRepository{
		db,
	}
	return newRepo
}

func (r *UserRepository) Find() ([]*db.User, error) {
	users := []*db.User{}
	if result := r.DB.Find(&users); result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *UserRepository) FindById(id uint) (*db.User, error) {
	user := db.User{}
	if result := r.DB.First(&user, id); result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) FindByAccountId(accountID string) (*db.User, error) {
	user := db.User{}
	if result := r.DB.Where("account_id = ?", accountID).First((&user)); result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) FindByUId(uid string) (*db.User, error) {
	user := db.User{}
	if result := r.DB.Where("uid = ?", uid).First((&user)); result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) Create(user *db.User) error {
	if result := r.DB.Create(user); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepository) Update(user *db.User) error {
	if result := r.DB.Save(user); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepository) Delete(user *db.User) error {
	if result := r.DB.Delete(user); result.Error != nil {
		return result.Error
	}
	return nil
}
