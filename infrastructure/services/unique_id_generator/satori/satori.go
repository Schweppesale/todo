package satori

import "github.com/satori/go.uuid"

type UniqueIdGenerator struct{}

func (u UniqueIdGenerator) Generate() string {
	return uuid.NewV4().String()
}
