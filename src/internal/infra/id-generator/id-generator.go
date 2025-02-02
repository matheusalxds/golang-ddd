package idGenerator

import "github.com/google/uuid"

type IdGenerator interface {
	GenerateUUID() (uuid.UUID, error)
}

type idGenerator struct{}

func NewIdGenerator() IdGenerator {
	return &idGenerator{}
}

func (i *idGenerator) GenerateUUID() (uuid.UUID, error) {
	return uuid.NewV7()
}
