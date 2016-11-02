package satori

import "github.com/satori/go.uuid"

type UuidService struct{}

func NewUuidService() UuidService {
	return UuidService{}
}

func (u UuidService) Generate() string {
	return uuid.NewV4().String()
}
