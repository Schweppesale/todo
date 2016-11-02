package response

import "time"

type TaskResponse struct {
	UniqueId    string
	Title       string
	Description string
	CreatedOn   time.Time
	UpdatedOn   time.Time
}
