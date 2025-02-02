package models

import (
	"github.com/google/uuid"
)

type UserEntity struct {
	ID    uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name  string    `gorm:"size:150;not null"`
	Email string    `gorm:"size:100;not null"`
}

func (user *UserEntity) TableName() string {
	return "users"
}
