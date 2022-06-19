package project

import (
	"go-just-portfolio/models"
	"mime/multipart"
)

type UseCase interface {
	//* No Auth
	GetProjectById(uuid string) (*models.Project, error)
	GetProjectsByShortname(shortname string, auth bool, user_id string) ([]models.ApiProject, error)
	GetProject(uuid string) models.InfoProjects
	//* Auth
	Newproject(user_uuid, category_uuid string) (*string, error)
	RenameProject(user_uuid, uuid, title string) error
	DeleteprojectById(uuid, user_id string) error
	SetStateproject(state int, uuid, user_id string) error
	LoadPhoto(file *multipart.FileHeader, user_uuid, project_uuid string, photo_type string) (*string, error)
	AddDescription(project_uuid, text string) (*string, error)
}
