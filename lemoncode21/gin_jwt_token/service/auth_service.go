package service

import "github.com/tohir-1302/gin_jwt_token/data/request"

type AuthService interface {
	Login(user request.LoginRequest) (string, error)
	Register(user request.LoginRequest) (string, error)
}
