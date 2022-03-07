package mysql

import (
	"go-just-portfolio/models"
	"go-just-portfolio/src/project"

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
	uuid := uuid.New().String()

	description := models.Description{
		UUID:        uuid,
		ProjectUUID: project_uuid,
		Key:         key,
		Value:       value,
		Language:    lang,
	}

	res := p.db.Create(&description)

	if res.Error != nil {
		return nil, res.Error
	}

	return &uuid, nil
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

func (p ProjectRepository) CreateTag(project_uuid string, tag string) (*string, error) {
	uuid := uuid.New().String()

	newTag := models.Tags{
		UUID:        uuid,
		ProjectUUID: project_uuid,
		Name:        tag,
	}

	res := p.db.Create(&newTag)

	if res.Error != nil {
		return nil, res.Error
	}

	return &uuid, nil
}

func (p ProjectRepository) UpdatePrewiew(project_uuid, prewiew_name string) error {
	query := "UPDATE projects SET prewiew = '" + prewiew_name + "' WHERE (uuid = '" + project_uuid + "');"

	var ss models.Project

	p.db.Raw(query).Scan(&ss)

	return nil
}

func (p ProjectRepository) SavePhoto(project_uuid, photo_name, photo_type string) (*string, error) {
	uuid := uuid.New().String()

	var image_type int

	if photo_type == "mobile" {
		image_type = 1
	} else {
		image_type = 0
	}

	photo := models.Photo{
		UUID:        uuid,
		ProjectUUID: project_uuid,
		Src:         photo_name,
		Type:        image_type,
	}

	res := p.db.Create(&photo)

	if res.Error != nil {
		return nil, res.Error
	}
	return &uuid, nil
}
