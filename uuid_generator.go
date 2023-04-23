package id_genius

import (
	"github.com/google/uuid"
)

type uuidGenerator struct{}

func newUuidGenerator() *uuidGenerator {
	return &uuidGenerator{}
}

func (u *uuidGenerator) GenerateNewId() (string, error) {
	return uuid.NewString(), nil
}
