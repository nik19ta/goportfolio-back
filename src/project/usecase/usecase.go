package usecase

import (
	"go-just-portfolio/models"
	"go-just-portfolio/src/project"
	"mime/multipart"
	"path/filepath"

	filesystem "go-just-portfolio/pkg/filesystem"

	"github.com/google/uuid"
)

type projectUseCase struct {
	userRepo project.ProjectRepository
}

func NewprojectUseCase(userRepo project.ProjectRepository) *projectUseCase {
	return &projectUseCase{userRepo: userRepo}
}

//* No Auth
func (p *projectUseCase) GetProjectById(uuid string) (*models.Project, error) {
	var project models.Project
	return &project, nil
}

func (p *projectUseCase) GetProjectsByShortname(shortname string, auth bool, user_id string) ([]models.ApiProject, error) {
	data, err := p.userRepo.GetProjectsByShortname(shortname)

	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return []models.ApiProject{}, nil
	}

	var sort_projects []models.ApiProject

	for _, project := range data {

		if project.State == 1 {
			if auth {
				if user_id == project.UserUUID {
					sort_projects = append(sort_projects, models.ApiProject{
						UUID:         project.UUID,
						CategoryUUID: project.CategoryUUID,
						Name:         project.Name,
						Prewiew:      project.Prewiew,
						State:        project.State,
					})
				}
			}
		} else {
			sort_projects = append(sort_projects, models.ApiProject{
				UUID:         project.UUID,
				CategoryUUID: project.CategoryUUID,
				Name:         project.Name,
				Prewiew:      project.Prewiew,
				State:        project.State,
			})
		}

	}

	return sort_projects, nil
}

//* Auth
func (p *projectUseCase) Newproject(user_uuid, category_uuid string) (*string, error) {
	uuid, err := p.userRepo.Newproject(user_uuid, category_uuid, "untitled")

	if err != nil {
		return nil, err
	}

	return uuid, nil
}
func (p *projectUseCase) DeleteprojectById(uuid, user_id string) error {
	err := p.userRepo.DeleteprojectById(uuid, user_id)

	if err != nil {
		return err
	}

	return nil
}
func (p *projectUseCase) SetStateproject(state int, uuid, user_id string) error {
	err := p.userRepo.SetStateproject(state, uuid, user_id)

	if err != nil {
		return err
	}

	return nil
}
func (p *projectUseCase) LoadPhoto(file *multipart.FileHeader, user_uuid, project_uuid string, photo_type string) error {

	extension := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + extension
	path := "./images/" + newFileName

	err := filesystem.SaveUploadedFile(file, path)

	if err != nil {
		return err
	}

	if photo_type == "prewiew" {
		p.userRepo.SavePhoto(project_uuid, newFileName, "prewiew")
	} else {
		p.userRepo.SavePhoto(project_uuid, newFileName, photo_type)
	}

	return nil
}

// func (a *projectUseCase) SignUp(username, mail, password, fullname string) (*string, error) {
// 	user := &models.User{
// 		Shortname: username,
// 		Mail:      mail,
// 		Password:  password,
// 		Fullname:  fullname,
// 		Type:      "user",
// 	}

// 	token, err := a.userRepo.CreateUser(user)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return token, nil
// }

// func (a *projectUseCase) SignIn(username, password string) (*string, error) {
// 	token, err := a.userRepo.GetUserToken(username, password)
// 	if err != nil {
// 		return nil, auth.ErrUserNotFound
// 	}

// 	return token, nil
// }

// func (a *projectUseCase) Profile(shortname string) (*models.User, error) {
// 	user, err := a.userRepo.GetUserInfo(shortname)
// 	if err != nil {
// 		return nil, auth.ErrUserNotFound
// 	}

// 	return user, nil
// }
