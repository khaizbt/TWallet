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

type IDUserInput struct {
	ID int `uri:"id" binding:"required"`
}

type TransactionFilterDate struct {
	StartDate string `json:"start_date" binding:"required"`
	EndDate   string `json:"end_date" binding:"required"`
}
