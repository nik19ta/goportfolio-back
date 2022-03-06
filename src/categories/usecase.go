package categories

import "go-just-portfolio/models"

type UseCase interface {
	//* No Auth
	GetCategoriesByUserName(shortname string) ([]models.Category, error)
	//* Auth
	AddCategory(user_uuid, title string) (*string, error)
	DeleteCategory(user_uuid, category_uuid string) error
	EditCategory(user_uuid, category_uuid, title string) error
}
