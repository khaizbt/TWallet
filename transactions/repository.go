package transactions

import (
	"TWallet/category"
	"TWallet/user"

	"gorm.io/gorm"
)

type Repository interface {
	ListTransaction(UserID int) ([]Transaction, error)
	Save(transaction Transaction) (Transaction, error)
	FindByCategoryID(ID int) (category.Category, error)
	UpdateUser(user user.User) (user.User, error)
	FindByIDUser(ID int) (user.User, error)
	DetailTransaction(ID int) (Transaction, error)
	DeleteTransaction(ID int) (Transaction, error)
	// FindByID(ID int) (Category, error)
	// UpdateBalance(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) ListTransaction(UserID int) ([]Transaction, error) {
	var transaction []Transaction
	err := r.db.Where("user_id = ?", UserID).Preload("Category").Find(&transaction).Error //Get All List category by User Login

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) Save(transaction Transaction) (Transaction, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) UpdateBalance(user user.User) (user.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByCategoryID(ID int) (category.Category, error) {
	var category category.Category
	err := r.db.Preload("User").Where("id = ?", ID).Find(&category).Error

	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) UpdateUser(user user.User) (user.User, error) {
	err := r.db.Save(&user).Error //Menggunakan Save karena untuk update/menyimpan data yang sudah ada

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByIDUser(ID int) (user.User, error) { //Find ID User untuk Update Balance
	var user user.User

	err := r.db.Where("id = ?", ID).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) DetailTransaction(ID int) (Transaction, error) { //GetDetailTransaction
	var transaction Transaction
	err := r.db.Where("id = ?", ID).Preload("User").Preload("Category").Find(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) DeleteTransaction(ID int) (Transaction, error) {
	var transaction Transaction
	err := r.db.Delete(&transaction, ID).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
