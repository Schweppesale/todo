package todo

// UUIDGenerator serves as a layer of abstraction around our UUID generator
type UUIDGenerator interface {
	Generate() string
}
