package service

import (
	"sample/src/data/request"
	"sample/src/data/response"
	"sample/src/helper"
	"sample/src/model"
	"sample/src/repository"

	"github.com/go-playground/validator/v10"
	"github.com/satori/go.uuid"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	validate       *validator.Validate
}

func NewUserServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{UserRepository: userRepository, validate: validate}
}

// Delete implements UserService.
func (u *UserServiceImpl) Delete(userId uuid.UUID) {
	u.UserRepository.Delete(userId)
}

// FindAll implements UserService.
func (u *UserServiceImpl) FindAll() []response.UserResponse {
	result := u.UserRepository.FindAll()
	var users []response.UserResponse
	for _, user := range result {
		users = append(users, response.UserResponse{
			Id:   user.Id,
			Name: user.Name,
		})
	}
	return users
}

// FindById implements UserService.
func (u *UserServiceImpl) FindById(userId uuid.UUID) response.UserResponse {
	_user, err := u.UserRepository.FindById(userId)
	helper.CheckErr(err)
	return response.UserResponse{
		Id:   _user.Id,
		Name: _user.Name,
	}
}

// Save implements UserService.
func (u *UserServiceImpl) Save(user request.CreateUserRequest) {
	err := u.validate.Struct(user)
	helper.CheckErr(err)
	userModel := model.User{
		Name: user.Name,
	}
	u.UserRepository.Save(userModel)
}

// Update implements UserService.
func (u *UserServiceImpl) Update(user request.UpdateUserRequest) {
	_user, err := u.UserRepository.FindById(user.Id)
	helper.CheckErr(err)
	_user.Name = user.Name
	u.UserRepository.Update(_user)
}
