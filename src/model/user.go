package model

import (
	"github.com/satori/go.uuid"
)

type User struct {
	Base
	Id        uuid.UUID    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name      string       `gorm:"type:varchar(255)"`
}
