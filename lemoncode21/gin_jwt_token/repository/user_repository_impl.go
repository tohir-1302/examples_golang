package repository

import (
	"errors"
	"github.com/tohir-1302/gin_jwt_token/data/request"
	"github.com/tohir-1302/gin_jwt_token/helper"
	"github.com/tohir-1302/gin_jwt_token/model"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		Db: db,
	}
}

func (u *UserRepositoryImpl) Save(user model.User) {
	result := u.Db.Create(&user)
	helper.ErrorPanic(result.Error)
}

func (u *UserRepositoryImpl) Update(user model.User) {
	var updateUser = request.UpdateUserRequest{
		Username: user.Username,
		Password: user.Password,
		Id:       user.Id,
		Email:    user.Email,
	}

	result := u.Db.Model(&user).Updates(updateUser)
	helper.ErrorPanic(result.Error)
}

func (u *UserRepositoryImpl) Delete(userId int) {
	var user model.User
	result := u.Db.Where("id = ?", userId).Delete(&user)
	helper.ErrorPanic(result.Error)
}

func (u *UserRepositoryImpl) FindById(userId int) (model.User, error) {
	var user model.User
	result := u.Db.Where("id = ?", userId).First(&user)

	if result.Error != nil {
		return user, errors.New("invalid Username or password")
	}

	return user, nil
}

func (u *UserRepositoryImpl) FindAll() ([]model.User, error) {
	var users []model.User
	result := u.Db.Find(&users)
	helper.ErrorPanic(result.Error)
	return users, nil
}

func (u *UserRepositoryImpl) FindByUsername(username string) (model.User, error) {
	var user model.User
	result := u.Db.Where("username = ?", username).First(&user)
	if result != nil {
		return user, nil
	}
	return user, errors.New("user not found")
}
