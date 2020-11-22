package transactions

import (
	"TWallet/category"
	"TWallet/user"
	"time"
)

type Transaction struct {
	ID          int
	UserID      int
	CategoryID  int
	Name        string
	Nominal     int
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        user.User
	Category    category.Category
}
