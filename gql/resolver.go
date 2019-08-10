//go:generate go run github.com/99designs/gqlgen
package gql

import (
	"context"
	"errors"

	"github.com/hubbdevelopers/auth"
	"github.com/hubbdevelopers/db"
	"github.com/hubbdevelopers/repository"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	UserRepo repository.User
	PageRepo repository.Page
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
	if result := auth.CheckValidUID(ctx, input.UID); result == false {
		return nil, errors.New("Auth Checker Error")
	}

	user := &db.User{
		Name:      input.Name,
		AccountID: input.AccountID,
		UID:       input.UID,
	}

	return user, r.UserRepo.Create(user)
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (bool, error) {
	user, err := r.UserRepo.FindById(uint(id))
	if err != nil {
		return false, err
	}

	if result := auth.Check(ctx, *user); result == false {
		return false, errors.New("Auth Checker Error")
	}

	err = r.UserRepo.Delete(user)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input UpdateUser) (*db.User, error) {
	user, err := r.UserRepo.FindById(uint(id))
	if err != nil {
		return nil, err
	}

	if result := auth.Check(ctx, *user); result == false {
		return nil, errors.New("Auth Checker Error")
	}

	if input.Description != nil {
		user.Description = *input.Description
	}

	if input.Name != nil {
		user.Name = *input.Name
	}

	if input.Image != nil {
		user.Image = *input.Image
	}

	if input.Twitter != nil {
		user.Twitter = *input.Twitter
	}

	if input.Instagram != nil {
		user.Instagram = *input.Instagram
	}

	if input.Facebook != nil {
		user.Facebook = *input.Facebook
	}

	if input.Homepage != nil {
		user.Homepage = *input.Homepage
	}

	return user, r.UserRepo.Update(user)
}

func (r *mutationResolver) CreatePage(ctx context.Context, input NewPage) (*db.Page, error) {
	user, err := r.UserRepo.FindById(uint(input.UserID))
	if err != nil {
		return nil, err
	}

	if result := auth.Check(ctx, *user); result == false {
		return nil, errors.New("Auth Checker Error")
	}

	page := &db.Page{
		Text:   input.Text,
		UserID: uint(input.UserID),
		Name:   input.Name,
	}

	err = r.PageRepo.Create(page)
	if err != nil {
		return nil, err
	}

	return page, nil
}

func (r *mutationResolver) UpdatePage(ctx context.Context, id int, input UpdatePage) (*db.Page, error) {
	page, err := r.PageRepo.FindById(uint(id))
	if err != nil {
		return nil, err
	}

	if result := auth.Check(ctx, page.User); result == false {
		return nil, errors.New("Auth Checker Error")
	}

	if input.Name != nil {
		page.Name = *input.Name
	}

	if input.Text != nil {
		page.Text = *input.Text
	}

	if input.Image != nil {
		page.Image = *input.Image
	}

	return page, r.PageRepo.Update(page)
}

func (r *mutationResolver) DeletePage(ctx context.Context, id int) (bool, error) {
	page, err := r.PageRepo.FindById(uint(id))
	if err != nil {
		return false, err
	}

	if result := auth.Check(ctx, page.User); result == false {
		return false, errors.New("Auth Checker Error")
	}

	err = r.PageRepo.Delete(page)
	if err != nil {
		return false, err
	}

	return true, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Pages(ctx context.Context) ([]*db.Page, error) {
	pages, err := r.PageRepo.Find()
	if err != nil {
		return nil, err
	}

	return pages, nil
}

func (r *queryResolver) Page(ctx context.Context, id *int) (*db.Page, error) {
	page, err := r.PageRepo.FindById(uint(*id))
	if err != nil {
		return nil, err
	}

	return page, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*db.User, error) {
	users, err := r.UserRepo.Find()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *queryResolver) User(ctx context.Context, id *int, accountID *string, uid *string) (*db.User, error) {
	if id != nil {
		return r.UserRepo.FindById(uint(*id))
	} else if accountID != nil {
		return r.UserRepo.FindByAccountId(*accountID)
	} else if uid != nil {
		return r.UserRepo.FindByUId(*uid)
	}
	return nil, errors.New("cannot find user")
}

type userResolver struct{ *Resolver }

func (r *userResolver) ID(ctx context.Context, obj *db.User) (int, error) {
	return int(obj.ID), nil
}

func (r *userResolver) Pages(ctx context.Context, obj *db.User) ([]*db.Page, error) {
	pages, err := r.PageRepo.FindByUser(obj)
	if err != nil {
		return nil, err
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
	user, err := r.UserRepo.FindById(obj.UserID)
	if err != nil {
		return nil, err
	}

	return user, nil
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
