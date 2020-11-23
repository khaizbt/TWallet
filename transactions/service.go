package transactions

import (
	"TWallet/category"
	"TWallet/user"

	"github.com/pkg/errors"
)

type Service interface {
	GetTransaction(UserID int) ([]Transaction, error)
	GetTransactionByDate(UserID int, StartDate string, EndDate string) ([]Transaction, error)
	SaveTransaction(input TransactionUserInput) (Transaction, error)
	CheckTypeCategory(UserID int, ID int) (category.Category, error)
	UpdateUser(ID int, data int) (user.User, error)
	GetDetailTransaction(input IDUserInput, UserID int) (Transaction, error)
	DeleteTransaction(input IDUserInput, UserID int) (Transaction, error)
}

type service struct { //Internal Struct
	repository Repository //Buat manggil struct repository melalui Interface repository
}

func NewService(repository Repository) *service { //
	return &service{repository} //mengirim ke interface harus pake pointer(ke alamat)
}

func (s *service) GetTransaction(UserID int) ([]Transaction, error) {
	transaction, err := s.repository.ListTransaction(UserID)

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (s *service) GetTransactionByDate(UserID int, StartDate string, EndDate string) ([]Transaction, error) {
	transaction, err := s.repository.ListTransactionFilter(UserID, StartDate, EndDate)

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (s *service) SaveTransaction(input TransactionUserInput) (Transaction, error) {
	transaction := Transaction{
		Name:        input.Name,
		CategoryID:  input.CategoryID,
		Nominal:     input.Nominal,
		Description: input.Description,
		UserID:      input.User.ID,
	}

	newTransaction, err := s.repository.Save(transaction)

	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil

}

func (s *service) CheckTypeCategory(UserID int, ID int) (category.Category, error) {

	category, err := s.repository.FindByCategoryID(ID)
	if err != nil {
		return category, err
	}

	if UserID != category.UserID {
		return category, errors.New("You Don't Have Access to this Category")
	}

	return category, nil
}

func (s *service) UpdateUser(ID int, data int) (user.User, error) {
	user, err := s.repository.FindByIDUser(ID) //Mendapatkan ID User

	if err != nil {
		return user, err
	}

	user.Balance = data

	updatedUser, err := s.repository.UpdateUser(user)

	if err != nil {
		return updatedUser, err
	}

	return updatedUser, err

}

func (s *service) GetDetailTransaction(input IDUserInput, UserID int) (Transaction, error) {
	transaction, err := s.repository.DetailTransaction(input.ID)

	if transaction.UserID != UserID {
		return transaction, errors.New("You not an owner this transaction")
	}

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (s *service) DeleteTransaction(input IDUserInput, UserID int) (Transaction, error) {
	transaction, err := s.repository.DeleteTransaction(input.ID)

	if transaction.UserID != UserID {
		return transaction, errors.New("You not an owner this transaction")
	}

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
