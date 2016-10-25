package entities

type Task struct {
	unique_id  string
	title      string
	created_on int
	updated_on int
}

func(t Task) UniqueId() string {
	return t.unique_id
}