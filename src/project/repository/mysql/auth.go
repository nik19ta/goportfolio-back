package mysql

import (
	"go-just-portfolio/models"
	"go-just-portfolio/src/project"
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type ProjectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (p ProjectRepository) Newproject(user_uuid, category_uuid, title string) (*string, error) {
	uuid := uuid.New().String()

	project := models.Project{
		UUID:         uuid,
		UserUUID:     user_uuid,
		CategoryUUID: category_uuid,
		Name:         title,
		Prewiew:      "empty",
		State:        1,
	}

	res := p.db.Create(&project)

	if res.Error != nil {
		return nil, res.Error
	}

	return &project.UUID, nil
}

func (p *ProjectRepository) GetProjectsByShortname(shortname string) ([]models.Project, error) {
	var projects []models.Project
	var user models.User
	p.db.Where("shortname = ?", shortname).Find(&user)

	if user.UUID == "" {
		return nil, project.ErrUserNotFound
	}

	p.db.Where("user_uuid = ?", user.UUID).Find(&projects)

	return projects, nil
}

func (p ProjectRepository) SetStateproject(state int, uuid, user_id string) error {
	res := p.db.Model(&models.Project{}).Where("uuid = ?", uuid).Update("state", state)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (p ProjectRepository) CreateDescription(project_uuid, key, value, lang string) (*string, error) {
	return nil, nil
}

func (p ProjectRepository) DeleteprojectById(project_uuid, user_uuid string) error {
	query := "DELETE FROM projects WHERE (uuid = '" + project_uuid + "' and user_uuid = '" + user_uuid + "');"

	var result []models.Project
	res := p.db.Raw(query).Scan(&result)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (p ProjectRepository) LoadPhoto(file *multipart.FileHeader) error {
	return nil
}

func (p ProjectRepository) CreateTag(project_uuid string, tag string) (*string, error) {
	return nil, nil
}

func (p ProjectRepository) SavePhoto(project_uuid, photo_name, photo_type string) error {
	return nil
}
