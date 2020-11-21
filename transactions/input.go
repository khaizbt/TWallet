package transactions

import (
	"TWallet/user"
)

type TransactionUserInput struct {
	Name        string `json:"name" binding:"required"`
	Nominal     int    `json:"nominal" binding:"required"`
	Description string `json:"description" binding:"required"`
	CategoryID  int    `json:"category_id", binding:"required"`
	User        user.User
	// Category    category.Category
}
