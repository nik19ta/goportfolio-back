package postgres

import (
	"go-just-portfolio/models"
	"go-just-portfolio/src/project"
	"log"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"

	c "go-just-portfolio/pkg/custom"
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
		Contents:     "",
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

func (p ProjectRepository) CreateDescription(project_uuid, text string) (*string, error) {
	uuid := uuid.New().String()

	description := models.Description{
		UUID:        uuid,
		ProjectUUID: project_uuid,
		Text:        text,
	}

	res := p.db.Create(&description)

	if res.Error != nil {
		return nil, res.Error
	}

	return &uuid, nil
}

func (p ProjectRepository) AddDescriptionIdToContent(project_uuid, description_uuid, type_ string) error {
	query := "SELECT * FROM projects WHERE (uuid = '" + project_uuid + "');"

	log.Println(query)

	var ss models.Project

	res := p.db.Raw(query).Scan(&ss)

	if res.Error != nil {
		return res.Error
	}

	new_content := c.InsertIntoString(ss.Contents, type_+"&"+description_uuid)

	log.Println(ss.Contents)
	log.Println(new_content)

	query2 := "UPDATE projects SET contents = '" + new_content + "' WHERE (uuid = '" + project_uuid + "');"

	var ss2 models.Project

	res = p.db.Raw(query2).Scan(&ss2)

	if res.Error != nil {
		return res.Error
	}

	return nil
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

func (p ProjectRepository) AddDescription(project_uuid, text string) (*string, error) {
	uuid := uuid.New().String()

	newDescription := models.Description{
		UUID:        uuid,
		ProjectUUID: uuid,
		Text:        text,
	}

	res := p.db.Create(&newDescription)

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

func (p *ProjectRepository) RenameProject(user_uuid, uuid, title string) error {
	query := "UPDATE projects SET name = '" + title + "' WHERE (uuid = '" + uuid + "') AND (user_uuid = '" + user_uuid + "');"

	var ss models.Project

	res := p.db.Raw(query).Scan(&ss)

	return res.Error
}

func (p *ProjectRepository) GetProjectById(uuid string) models.InfoProjects {
	query := "SELECT * FROM projects WHERE (uuid = '" + uuid + "');"

	var info models.InfoProjects

	var ss models.Project

	p.db.Raw(query).Scan(&ss)

	var content_photos []models.Photo
	var content_descs []models.Description

	arr := c.GetArray(ss.Contents)

	for _, item := range arr {
		if item != "" {
			arr := strings.Split(item, "&")
			log.Println(item)
			log.Println(arr)

			content_type := arr[0]
			content_c := arr[1]

			if content_type == "photo" {
				photo_query := "SELECT * FROM photos WHERE (uuid = '" + content_c + "');"
				var photo_content models.Photo
				p.db.Raw(photo_query).Scan(&photo_content)
				content_photos = append(content_photos, photo_content)
			} else {
				description_query := "SELECT * FROM descriptions WHERE (uuid = '" + content_c + "');"
				var description_content models.Description
				p.db.Raw(description_query).Scan(&description_content)
				content_descs = append(content_descs, description_content)
			}
		}
	}

	info.Main = ss
	info.Photos = content_photos
	info.Descriptions = content_descs
	return info
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
