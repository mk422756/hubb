//go:generate go run github.com/99designs/gqlgen
package gql

import (
	"context"

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

	dbOrm.Create(user)

	return user, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input UpdateUser) (*db.User, error) {
	dbOrm := db.GetDB()
	user := db.User{}
	dbOrm.First(&user, id)

	if input.Description != nil {
		dbOrm.Model(&user).Update("Description", *input.Description)
	}

	if input.Name != nil {
		dbOrm.Model(&user).Update("Name", *input.Name)
	}

	return &user, nil
}

func (r *mutationResolver) CreatePage(ctx context.Context, input NewPage) (*db.Page, error) {
	dbOrm := db.GetDB()
	page := &db.Page{
		Text:   input.Text,
		UserID: uint(input.UserID),
		Name:   input.Name,
	}
	dbOrm.Create(page)

	return page, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Pages(ctx context.Context) ([]*db.Page, error) {
	dbOrm := db.GetDB()
	pages := []*db.Page{}
	dbOrm.Find(&pages)

	return pages, nil
}

func (r *queryResolver) Page(ctx context.Context, id *int) (*db.Page, error) {
	dbOrm := db.GetDB()
	page := db.Page{}
	dbOrm.First(&page, *id)

	return &page, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*db.User, error) {
	dbOrm := db.GetDB()
	users := []*db.User{}
	dbOrm.Find(&users)

	return users, nil
}

func (r *queryResolver) User(ctx context.Context, id *int, accountID *string) (*db.User, error) {
	dbOrm := db.GetDB()
	user := db.User{}
	if id != nil {
		dbOrm.First(&user, *id)
	} else if accountID != nil {
		dbOrm.Where("account_id = ?", *accountID).First((&user))
	}

	return &user, nil
}

type userResolver struct{ *Resolver }

func (r *userResolver) ID(ctx context.Context, obj *db.User) (int, error) {
	return int(obj.ID), nil
}

func (r *userResolver) AccountID(ctx context.Context, obj *db.User) (string, error) {
	return obj.AccountID, nil
}

func (r *userResolver) Pages(ctx context.Context, obj *db.User) ([]*db.Page, error) {
	dbOrm := db.GetDB()
	pages := []*db.Page{}
	dbOrm.Model(obj).Related(&pages)
	return pages, nil
}

type pageResolver struct{ *Resolver }

func (r *pageResolver) User(ctx context.Context, obj *db.Page) (*db.User, error) {

	dbOrm := db.GetDB()
	user := db.User{}
	dbOrm.First(&user, obj.UserID)
	return &user, nil
}

func (r *pageResolver) ID(ctx context.Context, obj *db.Page) (int, error) {
	return int(obj.ID), nil
}
