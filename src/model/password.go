package model

import uuid "github.com/satori/go.uuid"

type Password struct {
	Base
	Id       uuid.UUID `gorm:"primaryKey" json:"id"`
	Password string    `json:"password"`
	UserId uuid.UUID    `gorm:"type:uuid;references:Id"`
}
