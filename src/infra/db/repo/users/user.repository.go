package repo

import (
	"fmt"
	"go-fx-project/src/infra/db/models"

	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(user *models.UserEntity) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) CreateUser(user *models.UserEntity) error {
	fmt.Println("newUser >>", user)
	return r.db.Create(user).Error
}
