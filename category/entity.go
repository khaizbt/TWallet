package category

import "time"

type Category struct {
	ID          int
	Name        string
	Type        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
