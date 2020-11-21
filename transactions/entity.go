package transactions

type Transaction struct {
	ID          int
	UserID      int
	CategoryID  int
	Name        string
	Nominal     int
	Description string
}

type User struct {
	ID      int
	Balance int
}
