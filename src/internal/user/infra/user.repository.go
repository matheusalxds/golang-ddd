package infra

import (
	"go-fx-project/src/internal/user/domain"

	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(user *domain.UserEntity) (*domain.UserEntity, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) CreateUser(user *domain.UserEntity) (*domain.UserEntity, error) {
	result := r.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
