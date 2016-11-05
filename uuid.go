package todo

type UUIDGenerator interface {
	Generate() string
}
