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

/*
	Метод позволяет получить проект по его uuid

	! Метод требует авторизации через bearer token если этот проект приватный (state == 1)
*/

func (p *projectUseCase) GetProjectById(uuid string) (*models.Project, error) {
	var project models.Project
	return &project, nil
}

/*
		Получить проекты пользователя по имени
		Есть три вида проектов state: 0, 1, 2
		* Если state = 0 - публичный для всех (Default)
	  * Если state = 1 - приватный для всех
	 	* Если state = 2 - доступный только по api

		Если авторизованный пользователь посылает запрос и его uuid
		совпадает с user_uuid проекта, значит он может получить приватный проект

		! Метод требует авторизации через bearer token если нужно получить приватные проекты
*/

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

/*
	Создаёт новый проект с именем untitled, возвращает uuid по
	которому можно добовлять материал в проект:
	* Фото, тектс, теги
	А также изменять:
	* Превью, название

	! Метод требует авторизации через bearer token
*/
func (p *projectUseCase) Newproject(user_uuid, category_uuid string) (*string, error) {
	uuid, err := p.userRepo.Newproject(user_uuid, category_uuid, "untitled")

	if err != nil {
		return nil, err
	}

	return uuid, nil
}

/*
	Удаляет проект по его uuid

	* Удалить проект может только его создать

	! При удаление проекта также удаляёться связанные с ним записи в таблицах
	! tags
	! descriptions
	! photo

	Все записи удаляёться по uuid проекта, для того что бы база данных не засорялась

	! Метод требует авторизации через bearer token
*/

func (p *projectUseCase) DeleteprojectById(uuid, user_id string) error {
	err := p.userRepo.DeleteprojectById(uuid, user_id)

	if err != nil {
		return err
	}

	return nil
}

/*
	Можно изменить видимоть проекта, от 0 до 2

	! Метод требует авторизации через bearer token
*/
func (p *projectUseCase) SetStateproject(state int, uuid, user_id string) error {
	err := p.userRepo.SetStateproject(state, uuid, user_id)

	if err != nil {
		return err
	}

	return nil
}

/*
	Можно изменить имя проекта, строка

	! Метод требует авторизации через bearer token
*/

func (p *projectUseCase) RenameProject(user_uuid, uuid, title string) error {
	err := p.userRepo.RenameProject(user_uuid, uuid, title)
	return err
}

/*
	Получить проект по id

	! Метод не требует авторизации через bearer token
*/

func (p *projectUseCase) GetProject(uuid string) models.InfoProjects {
	data := p.userRepo.GetProjectById(uuid)
	return data
}

/*
	Метод подгрузки фото

	Фото емеет три типа
	Превью самого проекта, записываеться в поле prewie проекта
	* 1. prewiew
	Фото записываеться в таблицу photos, имеют uuid на проект к которому они принаджлежать:
	* 2. mobile - это скриншот мобильного телефона или фото которые нужно просто поместить в ряд
	* 3. desktop - это большое фото на весь экран, его нельяза поместить в одну линию с другими

	В таблици тип фото представлен в int
	* 0: desktop
	* 1: mobile

	Фото сохроняеться в директорию ./images/ а в базу данных записываеться имя фото и расширение
	! Достпуные расширения фото (.png, .jpg, .jpeg)

	Для того что бы получить фото нужно послать запрос на:
	https://host/images/image_name.extention

	* функция: filesystem.SaveUploadedFile() сохроняет картинку
	* а функция p.userRepo.SavePhoto() сохроняет запись о картинки в бд

	! Метод требует авторизации через bearer token
*/

func (p *projectUseCase) LoadPhoto(file *multipart.FileHeader, user_uuid, project_uuid string, photo_type string) (*string, error) {

	extension := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + extension
	path := "./images/" + newFileName

	err := filesystem.SaveUploadedFile(file, path)

	if err != nil {
		return nil, err
	}

	if photo_type == "prewiew" {
		err = p.userRepo.UpdatePrewiew(project_uuid, newFileName)
	} else {
		photo_id, _ := p.userRepo.SavePhoto(project_uuid, newFileName, photo_type)

		_ = p.userRepo.AddDescriptionIdToContent(project_uuid, *photo_id, "photo")
	}

	if err != nil {
		return &newFileName, err
	}

	return &newFileName, nil
}

func (p *projectUseCase) AddDescription(project_uuid, text string) (*string, error) {
	uuid, err := p.userRepo.AddDescription(project_uuid, text)

	if err != nil {
		return nil, err
	}

	_ = p.userRepo.AddDescriptionIdToContent(project_uuid, *uuid, "text")

	return uuid, nil
}
