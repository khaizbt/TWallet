package category

import (
	"TWallet/user"
)

type CategoryUserInput struct {
	ID int `uri:"id" binding:"required"` //Ambil id dari URL
}
type CreateCategoryInput struct {
	Name        string `json:"name" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Description string `json:"description" binding:"required"`
	User        user.User
}
