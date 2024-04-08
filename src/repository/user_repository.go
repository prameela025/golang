package repository

import (
	"sample/src/model"

	"github.com/satori/go.uuid"
)

type UserRepository interface {
	Save(user model.User) 
	FindById(userId uuid.UUID) (user model.User, err error)
	FindAll() []model.User
	Update(user model.User)
	Delete(userId uuid.UUID)
}
