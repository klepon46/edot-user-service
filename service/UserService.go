package service

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/klepon46/edot-user-service/common/response"
	"github.com/klepon46/edot-user-service/model"
	"github.com/klepon46/edot-user-service/repository"
	"time"
)

type IUserService interface {
	Login(ctx context.Context, user model.User) (*model.UserResponse, error)
	Register(ctx context.Context, user model.User) (int, error)
	Authenticate(ctx context.Context, user model.User) (int, error)
}

type UserService struct {
	repoRegistry repository.IRegistry
}

func NewUserService(repoRegistry repository.IRegistry) *UserService {
	return &UserService{
		repoRegistry: repoRegistry,
	}
}

func (u UserService) Login(ctx context.Context, user model.User) (*model.UserResponse, error) {
	err := u.repoRegistry.GetUserRepository().Login(ctx, user)
	if err != nil {
		return nil, err
	}
	secretKey := []byte("secret")
	claims := model.UserClaims{
		UserRequest: model.UserRequest{
			Email: user.Email,
			Phone: user.Phone,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "user-service",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return nil, &response.Err{Response: *response.UnprocessableEntity(ctx).WithMessage(response.MessageUnprocessableEntity)}
	}

	return &model.UserResponse{
		Email: user.Email,
		Phone: user.Phone,
		Token: tokenString,
	}, nil
}

func (u UserService) Register(ctx context.Context, user model.User) (int, error) {
	register, err := u.repoRegistry.GetUserRepository().Register(ctx, user)
	if err != nil {
		return 0, err
	}

	return register, nil
}

func (u UserService) Authenticate(ctx context.Context, user model.User) (int, error) {
	return 0, nil
}
