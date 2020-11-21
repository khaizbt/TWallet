package category

import (
	"TWallet/user"
	"time"
)

type Category struct {
	ID          int
	UserID      int
	Name        string
	Type        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        user.User
}
