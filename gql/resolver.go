//go:generate go run github.com/99designs/gqlgen
package gql

import (
	"context"
	"errors"
	"time"

	"github.com/hubbdevelopers/auth"
	"github.com/hubbdevelopers/db"
	"github.com/hubbdevelopers/repository"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	UserRepo      repository.User
	PageRepo      repository.Page
	TagRepo       repository.Tag
	Authenticator auth.Authenticator
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

func (r *Resolver) Tag() TagResolver {
	return &tagResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input NewUser) (*db.User, error) {
	if result := r.Authenticator.IsAlreadyRegisteredUID(ctx, input.UID); result == false {
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

	if result := r.Authenticator.IsValidUID(ctx, user.UID); result == false {
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

	if result := r.Authenticator.IsValidUID(ctx, user.UID); result == false {
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

	if result := r.Authenticator.IsValidUID(ctx, user.UID); result == false {
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

	if input.Tags != nil {
		// create tag
		var tags []*db.Tag
		for _, tagName := range input.Tags {
			tag, err := r.TagRepo.FindByName(*tagName)
			if err != nil {
				tag = &db.Tag{
					Name: *tagName,
				}
			}
			tags = append(tags, tag)
		}
		r.PageRepo.Associate(page, tags)
	}

	return page, nil
}

func (r *mutationResolver) UpdatePage(ctx context.Context, id int, input UpdatePage) (*db.Page, error) {
	page, err := r.PageRepo.FindById(uint(id))
	if err != nil {
		return nil, err
	}

	user, err := r.UserRepo.FindById(page.UserID)
	if err != nil {
		return nil, err
	}

	if result := r.Authenticator.IsValidUID(ctx, user.UID); result == false {
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

	if input.Tags != nil {
		// create tag
		var tags []*db.Tag
		for _, tagName := range input.Tags {
			tag, err := r.TagRepo.FindByName(*tagName)
			if err != nil {
				tag = &db.Tag{
					Name: *tagName,
				}
			}
			tags = append(tags, tag)
		}
		r.PageRepo.Associate(page, tags)
	}

	return page, r.PageRepo.Update(page)
}

func (r *mutationResolver) DeletePage(ctx context.Context, id int) (bool, error) {
	page, err := r.PageRepo.FindById(uint(id))
	if err != nil {
		return false, err
	}

	user, err := r.UserRepo.FindById(page.UserID)
	if err != nil {
		return false, err
	}

	if result := r.Authenticator.IsValidUID(ctx, user.UID); result == false {
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

func (r *queryResolver) Tags(ctx context.Context) ([]*db.Tag, error) {
	tags, err := r.TagRepo.Find()

	if err != nil {
		return nil, err
	}

	return tags, nil
}

func (r *queryResolver) Tag(ctx context.Context, id *int) (*db.Tag, error) {
	tag, err := r.TagRepo.FindById(uint(*id))
	if err != nil {
		return nil, err
	}

	return tag, nil
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
	return obj.CreatedAt.Format(time.RFC3339), nil
}

func (r *userResolver) UpdatedAt(ctx context.Context, obj *db.User) (string, error) {
	return obj.UpdatedAt.Format(time.RFC3339), nil
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
	return obj.CreatedAt.Format(time.RFC3339), nil
}

func (r *pageResolver) UpdatedAt(ctx context.Context, obj *db.Page) (string, error) {
	return obj.UpdatedAt.Format(time.RFC3339), nil
}

func (r *pageResolver) Tags(ctx context.Context, obj *db.Page) ([]*db.Tag, error) {
	tags, err := r.TagRepo.FindByPageId(uint(obj.ID))
	if err != nil {
		return nil, err
	}
	return tags, nil
}

type tagResolver struct{ *Resolver }

func (r *tagResolver) ID(ctx context.Context, obj *db.Tag) (int, error) {
	return int(obj.ID), nil
}

func (r *tagResolver) CreatedAt(ctx context.Context, obj *db.Tag) (string, error) {
	return obj.CreatedAt.Format(time.RFC3339), nil
}

func (r *tagResolver) UpdatedAt(ctx context.Context, obj *db.Tag) (string, error) {
	return obj.UpdatedAt.Format(time.RFC3339), nil
}

func (r *tagResolver) Pages(ctx context.Context, obj *db.Tag) ([]*db.Page, error) {
	pages, err := r.PageRepo.FindByTagId(uint(obj.ID))
	if err != nil {
		return nil, err
	}
	return pages, nil
}
