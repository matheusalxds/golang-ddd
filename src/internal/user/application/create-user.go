package application

import (
	idGenerator "go-fx-project/src/internal/infra/id-generator"
	"go-fx-project/src/internal/infra/logger"
	"go-fx-project/src/internal/user/domain"
	"go-fx-project/src/internal/user/infra"

	"go.uber.org/zap"
)

var svcName = "UserService"
var fnName = "CreateUser"

type UserService interface {
	CreateUser(name, email string) (*domain.UserEntity, error)
}

type userService struct {
	repo        infra.UserRepo
	idGenerator idGenerator.IdGenerator
	logger      *zap.Logger
}

func NewUserService(repo infra.UserRepo, idGenerator idGenerator.IdGenerator, logger *zap.Logger) UserService {
	return &userService{repo: repo, idGenerator: idGenerator, logger: logger}
}

func (s *userService) CreateUser(name, email string) (*domain.UserEntity, error) {
	s.logger.Info(logger.MsgStart([]string{svcName, fnName}, map[string]interface{}{"name": name, "email": email}))

	id, err := s.idGenerator.GenerateUUID()
	if err != nil {
		return nil, err
	}

	user := &domain.UserEntity{Name: name, Email: email, ID: id}
	newUser, err := s.repo.CreateUser(user)

	if err != nil {
		return nil, err
	}

	s.logger.Info(logger.MsgEnd([]string{svcName, fnName}, map[string]interface{}{"newUser": newUser}))
	return newUser, nil
}
