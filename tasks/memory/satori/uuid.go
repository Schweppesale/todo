package satori

import "github.com/satori/go.uuid"

type UUIDGenerator struct{}

func NewUUIDGenerator() UUIDGenerator {
	return UUIDGenerator{}
}

func (u UUIDGenerator) Generate() string {
	return uuid.NewV4().String()
}
