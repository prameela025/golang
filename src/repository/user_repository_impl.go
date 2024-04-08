package repository

import (
	"errors"
	"sample/src/data/request"
	"sample/src/helper"
	"sample/src/model"

	"github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}

// Delete implements UserRepository.
func (u *UserRepositoryImpl) Delete(userId uuid.UUID) {
	var user model.User
	result := u.Db.Where("id = ?", userId).Delete(&user)
	helper.CheckErr(result.Error)
}

// FindAll implements UserRepository.
func (u *UserRepositoryImpl) FindAll() []model.User {
	var users []model.User
	result := u.Db.Find(&users)
	helper.CheckErr(result.Error)
	return users
}

// FindById implements UserRepository.
func (u *UserRepositoryImpl) FindById(userId uuid.UUID) (user model.User, err error) {
	var _user model.User
	result := u.Db.Find(&user, userId)
	if result.Error != nil {
		return _user, nil
	}
	return _user, errors.New("user not found")
}

// Save implements UserRepository.
func (u *UserRepositoryImpl) Save(user model.User) {
	result := u.Db.Create(&user)
	helper.CheckErr(result.Error)
}

// Update implements UserRepository.
func (u *UserRepositoryImpl) Update(user model.User) {
	var _user = request.UpdateUserRequest{
		Id:   user.Id,
		Name: user.Name,
	}
	result := u.Db.Model(&user).Updates(_user)
	helper.CheckErr(result.Error)
}
