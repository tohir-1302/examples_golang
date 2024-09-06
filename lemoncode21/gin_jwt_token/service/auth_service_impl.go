package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/tohir-1302/gin_jwt_token/data/request"
	"github.com/tohir-1302/gin_jwt_token/repository"
)

type AuthServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewAuthServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) *AuthService {
	return &AuthServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}

func (a *AuthServiceImpl) Login(user request.LoginRequest) (string, error) {
	panic("implement me")
}

func (a *AuthServiceImpl) Register(user request.CreateUserRequest) (string, error) {
	panic("implement me")
}
