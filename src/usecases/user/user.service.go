package usecases

import (
	"go-fx-project/src/infra/db/models"
	repo "go-fx-project/src/infra/db/repo/users"
	idGenerator "go-fx-project/src/infra/id-generator"
)

type UserService interface {
	CreateUser(name, email string) (*models.UserEntity, error)
}

type userService struct {
	repo        repo.UserRepo
	idGenerator idGenerator.IdGenerator
}

func NewUserService(repo repo.UserRepo, idGenerator idGenerator.IdGenerator) UserService {
	return &userService{repo: repo, idGenerator: idGenerator}
}

func (s *userService) CreateUser(name, email string) (*models.UserEntity, error) {
	id, err := s.idGenerator.GenerateUUID()
	if err != nil {
		return nil, err
	}
	user := &models.UserEntity{Name: name, Email: email, ID: id}
	if err := s.repo.CreateUser(user); err != nil {

	}
	return user, nil
}
