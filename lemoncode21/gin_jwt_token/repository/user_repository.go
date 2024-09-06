package repository

import "github.com/tohir-1302/gin_jwt_token/model"

type UserRepository interface {
	Save(user model.User)
	Update(user model.User)
	Delete(userId int)
	FindById(userId int) (model.User, error)
	FindAll() ([]model.User, error)
	FindByUsername(username string) (model.User, error)
}
