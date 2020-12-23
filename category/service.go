package category

import "errors"

type Service interface {
	GetCategory(userID int) ([]Category, error)
	GetDetailCategory(input CategoryUserInput, UserID int) (Category, error)
	CreateCategory(input CreateCategoryInput) (Category, error)
	UpdateCategory(inputID CategoryUserInput, inputData CreateCategoryInput) (Category, error)
	CheckTypeCategory(ID int) (Category, error)
	DeleteCategory(input CategoryUserInput, UserID int) (Category, error)
}

type service struct {
	repository Repository //Service mempunyai Ketergantungan ke repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCategory(userID int) ([]Category, error) {
	category, err := s.repository.ListCategory(userID)
	if err != nil {
		return category, err
	}

	return category, nil
}

func (s *service) GetDetailCategory(input CategoryUserInput, UserID int) (Category, error) {
	category, err := s.repository.FindByID(input.ID)

	if category.UserID != UserID {
		return category, errors.New("You not an owner this category")
	}

	if err != nil {
		return category, err
	}

	return category, nil
}

func (s *service) CreateCategory(input CreateCategoryInput) (Category, error) {
	category := Category{
		Name:        input.Name,
		Type:        input.Type,
		Description: input.Description,
		UserID:      input.User.ID,
	}

	newCategory, err := s.repository.Save(category)

	if err != nil {
		return newCategory, err
	}

	return newCategory, nil
}

func (s *service) UpdateCategory(inputID CategoryUserInput, inputData CreateCategoryInput) (Category, error) {
	category, err := s.repository.FindByID(inputID.ID)

	if err != nil {
		return category, err
	}

	if category.UserID != inputData.User.ID { //Handle ketika edit category user lain
		return category, errors.New("Not an owner of the category")
	}

	category.Name = inputData.Name
	category.Description = inputData.Description
	category.Type = inputData.Type

	updatedCategory, err := s.repository.Update(category)

	if err != nil {
		return updatedCategory, err
	}

	return updatedCategory, nil
}

func (s *service) CheckTypeCategory(ID int) (Category, error) {

	category, err := s.repository.FindByID(ID)

	if err != nil {
		return category, err
	}

	return category, nil
}

func (s *service) DeleteCategory(input CategoryUserInput, UserID int) (Category, error) {
	category, err := s.repository.DeleteCategory(input.ID)

	if category.UserID != UserID { //Handle ketika edit campaign user lain
		return category, errors.New("Not an owner of the campaign")
	}

	if err != nil {
		return category, err
	}

	return category, nil
}
