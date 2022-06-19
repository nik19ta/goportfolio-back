package http

import (
	"go-just-portfolio/models"
	jwt "go-just-portfolio/pkg/jwt"
	utils "go-just-portfolio/pkg/utils"
	project "go-just-portfolio/src/project"
	"log"

	"net/http"

	gin "github.com/gin-gonic/gin"
)

type Handler struct {
	useCase project.UseCase
}

func NewHandler(useCase project.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) GetProjectsByShortname(c *gin.Context) {
	var projects []models.ApiProject
	var err error
	if len(c.Request.Header["Authorization"]) != 0 {
		userid, err := jwt.GetFieldFromJWT(c.Request.Header["Authorization"][0], "id")

		if err != nil {
			projects, err = h.useCase.GetProjectsByShortname(c.Request.URL.Query()["shortname"][0], false, "")
		} else {
			projects, err = h.useCase.GetProjectsByShortname(c.Request.URL.Query()["shortname"][0], true, *userid)
		}

	} else {
		projects, err = h.useCase.GetProjectsByShortname(c.Request.URL.Query()["shortname"][0], false, "")
	}

	if err != nil {
		c.JSON(400, gin.H{"error": "true"})
		return
	}

	c.JSON(http.StatusOK, projects)
}

func (h *Handler) GetProjectById(c *gin.Context) {
	id := c.Request.URL.Query()["id"][0]

	project := h.useCase.GetProject(id)

	c.JSON(http.StatusOK, project)
}

func (h *Handler) ProjectSetState(c *gin.Context) {
	inp := new(models.SetStateInp)

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userid, err := utils.GetUserIdFromJWT(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	if inp.State > 2 {
		c.JSON(400, gin.H{"message": "state be only 0,1,2,3"})
		return
	}
	if inp.State < 0 {
		c.JSON(400, gin.H{"message": "state be only 0,1,2,3"})
		return
	}

	err = h.useCase.SetStateproject(inp.State, inp.ProjectUUID, *userid)

	if err != nil {
		c.JSON(400, gin.H{"message": "Something went wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"state": inp.State, "project": inp.ProjectUUID})
}

func (h *Handler) DeleteprojectById(c *gin.Context) {
	inp := new(models.DelProjectInp)

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userid, err := utils.GetUserIdFromJWT(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	err = h.useCase.DeleteprojectById(inp.UUID, *userid)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": true})
		return
	}

	c.JSON(http.StatusOK, gin.H{"deleted": inp.UUID})
}

//* Создаёт новый untitled проект
func (h *Handler) NewProject(c *gin.Context) {
	inp := new(models.NewProjectInp)

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userid, err := utils.GetUserIdFromJWT(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	uuid, err := h.useCase.Newproject(*userid, inp.CategoryUUID)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "something went wrong"})
	}

	c.JSON(http.StatusOK, gin.H{"project": uuid})
}

//* Создаёт новый untitled проект
func (h *Handler) RenameProject(c *gin.Context) {
	inp := new(models.RenameProjectInp)

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userid, err := utils.GetUserIdFromJWT(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	rename_err := h.useCase.RenameProject(*userid, inp.UUID, inp.Title)

	if rename_err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "something went wrong"})
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "uuid": inp.UUID, "title": inp.Title})
}

//* Добовляет текстовое описание к проекту
func (h *Handler) AddDescription(c *gin.Context) {
	inp := new(models.AddDescription)

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	_, err := utils.GetUserIdFromJWT(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	uuid, err := h.useCase.AddDescription(inp.UUID, inp.Test)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "something went wrong"})
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "description_id": uuid})
}

func (h *Handler) LoadPhotoPrewiew(c *gin.Context) {
	log.Println("UPDATE PHOTO !!!!!!!")
	file, err := c.FormFile("file")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}

	photo_type := c.PostForm("photo_type")
	project_uuid := c.PostForm("project_uuid")

	log.Println(photo_type)
	log.Println(project_uuid)

	userid, err := utils.GetUserIdFromJWT(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	uuid, err := h.useCase.LoadPhoto(file, *userid, project_uuid, photo_type)

	log.Println(err)
	log.Println(uuid)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "error upload file"})
		return
	}

	c.JSON(201, gin.H{"upload": uuid})
}
