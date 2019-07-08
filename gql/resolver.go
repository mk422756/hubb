//go:generate go run github.com/99designs/gqlgen
package gql

import (
	"context"
	"time"

	"github.com/hubbdevelopers/db"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}

func (r *Resolver) Page() PageResolver {
	return &pageResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input NewUser) (*db.User, error) {
	dbOrm := db.GetDB()
	user := &db.User{
		Name:      input.Name,
		AccountID: input.AccountID,
		UID:       input.UID,
	}

	if result := dbOrm.Create(user); result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input UpdateUser) (*db.User, error) {
	dbOrm := db.GetDB()
	user := db.User{}
	if result := dbOrm.First(&user, id); result.Error != nil {
		return nil, result.Error
	}

	tx := dbOrm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if input.Description != nil {
		if result := tx.Model(&user).Update("Description", *input.Description); result.Error != nil {
			tx.Rollback()
			return nil, result.Error
		}
	}

	if input.Name != nil {
		if result := tx.Model(&user).Update("Name", *input.Name); result.Error != nil {
			tx.Rollback()
			return nil, result.Error
		}
	}

	if input.Birthday != nil {
		date, err := time.Parse("2006-01-02", *input.Birthday)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		if result := tx.Model(&user).Update("Birthday", date); result.Error != nil {
			tx.Rollback()
			return nil, result.Error
		}
	}

	if input.Twitter != nil {
		if result := tx.Model(&user).Update("Twitter", *input.Twitter); result.Error != nil {
			tx.Rollback()
			return nil, result.Error
		}
	}

	if input.Instagram != nil {
		if result := tx.Model(&user).Update("Instagram", *input.Instagram); result.Error != nil {
			tx.Rollback()
			return nil, result.Error
		}
	}

	if input.Facebook != nil {
		if result := tx.Model(&user).Update("Facebook", *input.Facebook); result.Error != nil {
			tx.Rollback()
			return nil, result.Error
		}
	}

	if input.Homepage != nil {
		if result := tx.Model(&user).Update("Homepage", *input.Homepage); result.Error != nil {
			tx.Rollback()
			return nil, result.Error
		}
	}

	return &user, tx.Commit().Error
}

func (r *mutationResolver) CreatePage(ctx context.Context, input NewPage) (*db.Page, error) {
	dbOrm := db.GetDB()
	page := &db.Page{
		Text:   input.Text,
		UserID: uint(input.UserID),
		Name:   input.Name,
	}
	if result := dbOrm.Create(page); result.Error != nil {
		return nil, result.Error
	}

	return page, nil
}

func (r *mutationResolver) UpdatePage(ctx context.Context, id int, input UpdatePage) (*db.Page, error) {
	dbOrm := db.GetDB()
	page := db.Page{}
	if result := dbOrm.First(&page, id); result.Error != nil {
		return nil, result.Error
	}

	tx := dbOrm.Begin()
	if input.Name != nil {
		if result := tx.Model(&page).Update("Name", *input.Name); result.Error != nil {
			tx.Rollback()
			return nil, result.Error
		}
	}

	if input.Text != nil {
		if result := tx.Model(&page).Update("Text", *input.Text); result.Error != nil {
			tx.Rollback()
			return nil, result.Error
		}
	}

	return &page, tx.Commit().Error
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Pages(ctx context.Context) ([]*db.Page, error) {
	dbOrm := db.GetDB()
	pages := []*db.Page{}
	if result := dbOrm.Find(&pages); result.Error != nil {
		return nil, result.Error
	}

	return pages, nil
}

func (r *queryResolver) Page(ctx context.Context, id *int) (*db.Page, error) {
	dbOrm := db.GetDB()
	page := db.Page{}
	if result := dbOrm.First(&page, *id); result.Error != nil {
		return nil, result.Error
	}

	return &page, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*db.User, error) {
	dbOrm := db.GetDB()
	users := []*db.User{}
	if result := dbOrm.Find(&users); result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (r *queryResolver) User(ctx context.Context, id *int, accountID *string, uid *string) (*db.User, error) {
	dbOrm := db.GetDB()
	user := db.User{}
	if id != nil {
		if result := dbOrm.First(&user, *id); result.Error != nil {
			return nil, result.Error
		}
	} else if accountID != nil {
		if result := dbOrm.Where("account_id = ?", *accountID).First((&user)); result.Error != nil {
			return nil, result.Error
		}
	} else if uid != nil {
		if result := dbOrm.Where("uid = ?", *uid).First((&user)); result.Error != nil {
			return nil, result.Error
		}
	}

	return &user, nil
}

type userResolver struct{ *Resolver }

func (r *userResolver) ID(ctx context.Context, obj *db.User) (int, error) {
	return int(obj.ID), nil
}

func (r *userResolver) Birthday(ctx context.Context, obj *db.User) (string, error) {
	return obj.Birthday.Format("2006-01-02"), nil
}

func (r *userResolver) Pages(ctx context.Context, obj *db.User) ([]*db.Page, error) {
	dbOrm := db.GetDB()
	pages := []*db.Page{}
	if result := dbOrm.Model(obj).Related(&pages); result.Error != nil {
		return nil, result.Error
	}
	return pages, nil
}

func (r *userResolver) CreatedAt(ctx context.Context, obj *db.User) (string, error) {
	return obj.CreatedAt.String(), nil
}

func (r *userResolver) UpdatedAt(ctx context.Context, obj *db.User) (string, error) {
	return obj.UpdatedAt.String(), nil
}

type pageResolver struct{ *Resolver }

func (r *pageResolver) User(ctx context.Context, obj *db.Page) (*db.User, error) {

	dbOrm := db.GetDB()
	user := db.User{}
	if result := dbOrm.First(&user, obj.UserID); result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *pageResolver) ID(ctx context.Context, obj *db.Page) (int, error) {
	return int(obj.ID), nil
}

func (r *pageResolver) CreatedAt(ctx context.Context, obj *db.Page) (string, error) {
	return obj.CreatedAt.String(), nil
}

func (r *pageResolver) UpdatedAt(ctx context.Context, obj *db.Page) (string, error) {
	return obj.UpdatedAt.String(), nil
}
