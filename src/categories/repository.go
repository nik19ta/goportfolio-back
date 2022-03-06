package categories

import "go-just-portfolio/models"

type СategoriesRepository interface {
	GetCategoriesByUserName(shortname string) ([]models.Category, error)
	DeleteCategory(user_uuid, category_uuid string) error
	EditCategory(user_uuid, category_uuid, title string) error
	AddCategory(user_uuid, title string) (*string, error)
}
