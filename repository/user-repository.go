package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/klepon46/edot-user-service/common/response"
	"github.com/klepon46/edot-user-service/model"
)

type IUserRepository interface {
	Login(ctx context.Context, user model.User) error
	Register(ctx context.Context, user model.User) (int, error)
	Authenticate(ctx context.Context, user model.User) (int, error)
}

type UserRepository struct {
	Db *sqlx.DB
}

func NewUserRepository(Db *sqlx.DB) *UserRepository {
	return &UserRepository{
		Db: Db,
	}
}

func (u UserRepository) Login(ctx context.Context, user model.User) error {
	//for _, u := range *u.Users {
	//	if (u.Phone == user.Phone || u.Email == user.Email) && u.Password == user.Password {
	//		return nil
	//	}
	//}
	//return &response.Err{Response: *response.NotFound(ctx).WithMessage(response.MessageNotFound)}

	err := u.Db.Get(&user, "SELECT * FROM user WHERE (email = ? OR phone = ? ) AND password = ?",
		user.Email, user.Phone, user.Password)

	if err != nil {
		return &response.Err{Response: *response.NotFound(ctx).WithMessage(response.MessageNotFound)}
	}

	return nil
}

func (u UserRepository) Register(ctx context.Context, user model.User) (int, error) {
	result, err := u.Db.Exec("INSERT INTO user ( email, phone, password) VALUES (?, ?, ?)",
		user.Email, user.Phone, user.Password)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (u UserRepository) Authenticate(ctx context.Context, user model.User) (int, error) {
	//TODO implement me
	panic("implement me")
}
