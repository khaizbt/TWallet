package category

import "gorm.io/gorm"

type Repository interface {
	ListCategory(userID int) ([]Category, error)
	Save(category Category) (Category, error)
	Update(category Category) (Category, error)
	FindByID(ID int) (Category, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) ListCategory(UserID int) ([]Category, error) {
	var category []Category
	err := r.db.Where("user_id = ?", UserID).Find(&category).Error //Get All List category by User Login

	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) Save(category Category) (Category, error) {
	err := r.db.Create(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) FindByID(ID int) (Category, error) {
	var category Category
	err := r.db.Preload("User").Where("id = ?", ID).Find(&category).Error

	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) Update(category Category) (Category, error) {
	err := r.db.Save(&category).Error

	if err != nil {
		return category, err
	}

	return category, nil
}
