package user

import (
	"time"
)

type User struct {
	ID          int
	Name        string
	Email       string
	Password    string
	AvatarImage string
	Role        string
	Balance     int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
