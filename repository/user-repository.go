package repository

import (
	"context"
	"github.com/klepon46/edot-user-service/common/response"
	"github.com/klepon46/edot-user-service/model"
)

type IUserRepository interface {
	Login(ctx context.Context, user model.User) error
	Register(ctx context.Context, user model.User) (int, error)
	Authenticate(ctx context.Context, user model.User) (int, error)
}

type UserRepository struct {
	Users *[]model.User
}

func NewUserRepository(users *[]model.User) *UserRepository {
	return &UserRepository{
		Users: users,
	}
}

func (u UserRepository) Login(ctx context.Context, user model.User) error {
	for _, u := range *u.Users {
		if (u.Phone == user.Phone || u.Email == user.Email) && u.Password == user.Password {
			return nil
		}
	}
	return &response.Err{Response: *response.NotFound(ctx).WithMessage(response.MessageNotFound)}
}

func (u UserRepository) Register(ctx context.Context, user model.User) (int, error) {
	user.ID = len(*u.Users) + 1
	*u.Users = append(*u.Users, user)
	return user.ID, nil
}

func (u UserRepository) Authenticate(ctx context.Context, user model.User) (int, error) {
	//TODO implement me
	panic("implement me")
}
