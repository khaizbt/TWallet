package category

import "gorm.io/gorm"

type Repository interface {
	ListCategory() ([]Category, error)
	Save(category Category) (Category, error)
	Update(category Category) (Category, error)
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
