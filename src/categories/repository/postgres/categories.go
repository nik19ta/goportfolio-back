package postgres

import (
	"go-just-portfolio/models"
	"go-just-portfolio/src/categories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type СategoriesRepository struct {
	db *gorm.DB
}

func NewСategoriesRepository(db *gorm.DB) *СategoriesRepository {
	return &СategoriesRepository{db: db}
}

func (c *СategoriesRepository) AddCategory(user_uuid, title string) (*string, error) {
	uuid := uuid.New().String()

	category := models.Category{
		UUID:     uuid,
		UserUUID: user_uuid,
		Name:     title,
	}

	result := c.db.Create(&category)

	if result.Error != nil {
		return nil, categories.DataBaseError
	}

	return &uuid, nil
}

func (c *СategoriesRepository) GetCategoriesByUserName(shortname string) ([]models.Category, error) {
	query := "SELECT uuid, user_uuid, name FROM categories WHERE user_uuid = (SELECT uuid FROM users WHERE shortname = '" + shortname + "');"

	var categories []models.Category
	c.db.Raw(query).Scan(&categories)

	return categories, nil
}

func (c *СategoriesRepository) DeleteCategory(user_uuid, category_uuid string) error {
	query := "DELETE FROM categories WHERE (uuid = '" + category_uuid + "' and user_uuid = '" + user_uuid + "');"

	var result []models.Category
	res := c.db.Raw(query).Scan(&result)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (c *СategoriesRepository) EditCategory(user_uuid, category_uuid, title string) error {
	query := "UPDATE categories SET name = '" + title + "'  WHERE (uuid = '" + category_uuid + "' and user_uuid = '" + user_uuid + "')"

	var result []models.Category
	res := c.db.Raw(query).Scan(&result)

	if res.Error != nil {
		return res.Error
	}

	return nil
}
