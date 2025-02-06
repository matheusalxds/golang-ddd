package application

import (
	"errors"
	"go-fx-project/src/internal/user/domain"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

type MockStruct struct {
	mock.Mock
}

func (m *MockStruct) GenerateUUID() (uuid.UUID, error) {
	args := m.Called()
	return args.Get(0).(uuid.UUID), args.Error(1)
}

func (m *MockStruct) CreateUser(user *domain.UserEntity) (*domain.UserEntity, error) {
	args := m.Called(user)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.UserEntity), args.Error(1)
	}
	return nil, args.Error(1)
}

type Setup struct {
	Logger  *zap.Logger
	Repo    *MockStruct
	IDGen   *MockStruct
	UUIDVal uuid.UUID
	Name    string
	Email   string
}

func setup() Setup {
	repo := new(MockStruct)
	idGen := new(MockStruct)
	logger := zap.NewNop()
	uuidVal, _ := uuid.Parse("123e4567-e89b-12d3-a456-426614174000")
	name := "any_name"
	email := "any_email@mail.com"

	idGen.On("GenerateUUID").Return(uuidVal, nil)
	repo.On("CreateUser", mock.AnythingOfType("*domain.UserEntity")).Return(&domain.UserEntity{
		Name:  name,
		Email: email,
		ID:    uuidVal,
	}, nil)

	return Setup{Logger: logger, Repo: repo, IDGen: idGen, UUIDVal: uuidVal, Name: name, Email: email}
}

func TestCreateUser(t *testing.T) {

	t.Run("should call idGen with correct params", func(t *testing.T) {
		s := setup()
		idGen := s.IDGen
		sut := NewUserService(s.Repo, idGen, s.Logger)

		sut.CreateUser("any_name", "any_email@mail")

		idGen.AssertNumberOfCalls(t, "GenerateUUID", 1)
	})

	t.Run("should return an error if idGen returns error", func(t *testing.T) {
		s := setup()
		newError := errors.New("error to generate uuid")
		idGen := new(MockStruct)
		idGen.On("GenerateUUID").Return(uuid.Nil, newError)
		sut := NewUserService(s.Repo, idGen, s.Logger)

		_, err := sut.CreateUser("any_name", "any_email")

		assert.Error(t, err)
		assert.Equal(t, err.Error(), newError.Error())
	})

	t.Run("should call repo.CreateUser with correct params", func(t *testing.T) {
		s := setup()
		repo := s.Repo
		sut := NewUserService(repo, s.IDGen, s.Logger)
		repo.On("CreateUser", &domain.UserEntity{Name: "any_name", Email: "any_email", ID: s.UUIDVal}).Return(nil)

		sut.CreateUser("any_name", "any_email")

		repo.AssertNumberOfCalls(t, "CreateUser", 1)
		repo.AssertCalled(t, "CreateUser", &domain.UserEntity{Name: "any_name", Email: "any_email", ID: s.UUIDVal})
	})

	t.Run("should return an error if repo returns error", func(t *testing.T) {
		s := setup()
		newError := errors.New("error to create user")
		repo := new(MockStruct)
		repo.On("CreateUser", mock.AnythingOfType("*domain.UserEntity")).Return(&domain.UserEntity{}, newError)
		sut := NewUserService(repo, s.IDGen, s.Logger)

		_, err := sut.CreateUser("any_name", "any_email")

		assert.Error(t, err)
		assert.Equal(t, err.Error(), newError.Error())
	})

	t.Run("should create user successfully", func(t *testing.T) {
		s := setup()
		name := "any_name"
		email := "any_email@example.com"
		repo := new(MockStruct)
		repo.On("CreateUser", mock.AnythingOfType("*domain.UserEntity")).Return(&domain.UserEntity{Name: name, Email: email, ID: s.UUIDVal}, nil)
		sut := NewUserService(repo, s.IDGen, s.Logger)

		user, _ := sut.CreateUser(name, email)

		assert.Equal(t, name, user.Name)
		assert.Equal(t, email, user.Email)
		assert.Equal(t, s.UUIDVal, user.ID)
	})
}
