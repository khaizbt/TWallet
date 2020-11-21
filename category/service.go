package category

import "errors"

type Service interface {
	GetCategory(userID int) ([]Category, error)
	GetDetailCategory(input CategoryUserInput) (Category, error)
	CreateCategory(input CreateCategoryInput) (Category, error)
	UpdateCampaign(inputID CategoryUserInput, inputData CreateCategoryInput) (Category, error)
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

func (s *service) GetDetailCategory(input CategoryUserInput) (Category, error) {
	campaign, err := s.repository.FindByID(input.ID)

	if err != nil {
		return campaign, err
	}

	return campaign, nil
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

func (s *service) UpdateCampaign(inputID CategoryUserInput, inputData CreateCategoryInput) (Category, error) {
	campaign, err := s.repository.FindByID(inputID.ID)

	if err != nil {
		return campaign, err
	}

	if campaign.UserID != inputData.User.ID { //Handle ketika edit campaign user lain
		return campaign, errors.New("Not an owner of the campaign")
	}

	campaign.Name = inputData.Name
	campaign.Description = inputData.Description
	campaign.Type = inputData.Type

	updatedCampaign, err := s.repository.Update(campaign)

	if err != nil {
		return updatedCampaign, err
	}

	return updatedCampaign, nil
}
