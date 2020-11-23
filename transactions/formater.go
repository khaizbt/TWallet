package transactions

type TransactionFormater struct {
	ID         int              `json:"id"`
	UserID     int              `json:"user_id"`
	CategoryID int              `json:"category_id"`
	Name       string           `json:"name"`
	Nominal    int              `json:"nominal"`
	Category   CategoryFormater `json:"category"`
}

type TransactionDetailFormater struct {
	ID          int                  `json:"id"`
	Name        string               `json:"name"`
	Nominal     int                  `json:"nominal"`
	Description string               `json:"description"`
	UserID      int                  `json:"user_id"`
	CategoryID  int                  `json:"category_id"`
	User        CategoryUserFormater `json:"user"`
	Category    CategoryFormater     `json:"category"`
}

type CategoryFormater struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type CategoryUserFormater struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
	Balance  int    `json:"balance"`
}

func FormatTransaction(transaction Transaction) TransactionFormater { //Token akan didapatkan dari JWT
	formater := TransactionFormater{
		ID:         transaction.ID,
		UserID:     transaction.UserID,
		CategoryID: transaction.CategoryID,
		Name:       transaction.Name,
		Nominal:    transaction.Nominal,
	}

	category := transaction.Category
	formaterCategory := CategoryFormater{
		Name: category.Name,
		Type: category.Type,
	}

	formater.Category = formaterCategory

	return formater
}

func FormatTransactions(transactions []Transaction) []TransactionFormater {
	transactionsFormater := []TransactionFormater{}

	for _, transaction := range transactions {
		transactionFormater := FormatTransaction(transaction)                    //Pemanggilan Fungsi Format Campaign
		transactionsFormater = append(transactionsFormater, transactionFormater) //Perhatikan ada huruf s yang berarti jamak
	}

	return transactionsFormater
}

func FormatTransactionDetail(transaction Transaction) TransactionDetailFormater {
	formater := TransactionDetailFormater{
		ID:          transaction.ID,
		UserID:      transaction.UserID,
		CategoryID:  transaction.CategoryID,
		Name:        transaction.Name,
		Nominal:     transaction.Nominal,
		Description: transaction.Description,
	}

	user := transaction.User
	userFormater := CategoryUserFormater{
		Name:     user.Name,
		ImageURL: user.AvatarImage,
		Balance:  user.Balance,
	}

	category := transaction.Category

	categoryFormater := CategoryFormater{
		Name: category.Name,
		Type: category.Type,
	}

	formater.Category = categoryFormater

	formater.User = userFormater

	return formater
}

type TransactionFormaterCreate struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	CategoryID int    `json:"category_id"`
	Name       string `json:"name"`
	Nominal    int    `json:"nominal"`
}

func FormatTransactionCreate(transaction Transaction) TransactionFormaterCreate { //Token akan didapatkan dari JWT
	formater := TransactionFormaterCreate{
		ID:         transaction.ID,
		UserID:     transaction.UserID,
		CategoryID: transaction.CategoryID,
		Name:       transaction.Name,
		Nominal:    transaction.Nominal,
	}

	return formater
}
