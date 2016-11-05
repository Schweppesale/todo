package todo

// UUIDGenerator serves as a layer of abstraction when creating a UUID
type UUIDGenerator interface {
	Generate() string
}
