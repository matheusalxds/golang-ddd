package application

import (
	idGenerator "go-fx-project/src/internal/infra/id-generator"
	"go-fx-project/src/internal/user/domain"
	"go-fx-project/src/internal/user/infra"
)

type UserService interface {
	CreateUser(name, email string) (*domain.UserEntity, error)
}

type userService struct {
	repo        infra.UserRepo
	idGenerator idGenerator.IdGenerator
}

func NewUserService(repo infra.UserRepo, idGenerator idGenerator.IdGenerator) UserService {
	return &userService{repo: repo, idGenerator: idGenerator}
}

func (s *userService) CreateUser(name, email string) (*domain.UserEntity, error) {
	id, err := s.idGenerator.GenerateUUID()
	if err != nil {
		return nil, err
	}
	user := &domain.UserEntity{Name: name, Email: email, ID: id}
	if err := s.repo.CreateUser(user); err != nil {

	}
	return user, nil
}
