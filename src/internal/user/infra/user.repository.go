package infra

import (
	"go-fx-project/src/internal/user/domain"

	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(user *domain.UserEntity) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) CreateUser(user *domain.UserEntity) error {
	return r.db.Create(user).Error
}
