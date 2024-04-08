package model

import ( 
	"time" 
	"github.com/satori/go.uuid" 
)

type Base struct {
	CreatedAt time.Time      `json:"-" gorm:"autoCreateTime"`
	CreatedBy *uuid.UUID      `json:"-" gorm:"type:uuid"`
	ModifiedAt *time.Time      `json:"-" gorm:"index"`
	ModifiedBy *uuid.UUID      `json:"-" gorm:"type:uuid"`
	DeletedAt *time.Time `json:"-" gorm:"index"`
}
