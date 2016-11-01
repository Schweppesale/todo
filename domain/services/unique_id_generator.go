package services

type UniqueIdGenerator interface {
	Generate() string
}
