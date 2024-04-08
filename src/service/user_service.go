package service

import (
	"sample/src/data/request"
	"sample/src/data/response"

	"github.com/satori/go.uuid"
)

type UserService interface {
	Save(request.CreateUserRequest)
	FindById(userId uuid.UUID) response.UserResponse
	FindAll() []response.UserResponse
	Update(request.UpdateUserRequest)
	Delete(userId uuid.UUID)
}
