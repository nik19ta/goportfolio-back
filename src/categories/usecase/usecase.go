package usecase

import (
	"go-just-portfolio/models"
	"go-just-portfolio/src/categories"
)

type СategoriesUseCase struct {
	userRepo categories.СategoriesRepository
}

func NewСategoriesUseCase(userRepo categories.СategoriesRepository) *СategoriesUseCase {
	return &СategoriesUseCase{userRepo: userRepo}
}

//* No Auth
func (r *СategoriesUseCase) GetCategoriesByUserName(shortname string) ([]models.Category, error) {
	res, err := r.userRepo.GetCategoriesByUserName(shortname)

	if err != nil {
		return nil, err
	}

	return res, nil
}

//* Auth
func (r *СategoriesUseCase) AddCategory(user_uuid, title string) (*string, error) {
	uuid, err := r.userRepo.AddCategory(user_uuid, title)

	if err != nil {
		return nil, err
	}
	return uuid, nil
}

func (r *СategoriesUseCase) DeleteCategory(user_uuid, category_uuid string) error {
	err := r.userRepo.DeleteCategory(user_uuid, category_uuid)

	if err != nil {
		return err
	}

	return nil
}

func (r *СategoriesUseCase) EditCategory(user_uuid, category_uuid, title string) error {
	err := r.userRepo.EditCategory(user_uuid, category_uuid, title)

	if err != nil {
		return err
	}

	return nil
}
