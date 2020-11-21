package category

type CategoryFormater struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
}

func FormatCategory(category Category) CategoryFormater { //Token akan didapatkan dari JWT
	formater := CategoryFormater{
		ID:     category.ID,
		UserID: category.UserID,
		Name:   category.Name,
		Type:   category.Type,
	}

	return formater
}

func FormatCategories(categories []Category) []CategoryFormater {
	categoriesFormater := []CategoryFormater{}

	for _, category := range categories {
		categoryFormater := FormatCategory(category)                      //Pemanggilan Fungsi Format Campaign
		categoriesFormater = append(categoriesFormater, categoryFormater) //Perhatikan ada huruf s yang berarti jamak
	}

	return categoriesFormater
}

type CategoryDetailFormater struct {
	ID          int                  `json:"id"`
	Name        string               `json:"name"`
	Type        string               `json:"name"`
	Description string               `json:"description"`
	UserID      int                  `json:"user_id"`
	User        CategoryUserFormater `json:"user"`
}

type CategoryUserFormater struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func FormatCampaignDetail(category Category) CategoryDetailFormater {
	formater := CategoryDetailFormater{
		ID:          category.ID,
		UserID:      category.UserID,
		Name:        category.Name,
		Type:        category.Type,
		Description: category.Description,
	}

	user := category.User
	userFormater := CategoryUserFormater{
		Name:     user.Name,
		ImageURL: user.AvatarImage,
	}

	formater.User = userFormater

	return formater
}
