package http

import (
	"go-just-portfolio/models"
	jwt "go-just-portfolio/pkg/jwt"
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

type signInput struct {
	Shortname string `json:"shortname"`
	Mail      string `json:"mail"`
	Password  string `json:"password"`
	Fullname  string `json:"fullname"`
}

func (h *Handler) GetProjectById(c *gin.Context) {

}

func (h *Handler) GetProjectsByShortname(c *gin.Context) {

	var projects []models.ApiProject
	var err error
	if len(c.Request.Header["Authorization"]) != 0 {
		userid, err := jwt.GetFieldFromJWT(c.Request.Header["Authorization"][0], "id")

		if err != nil {
			projects, err = h.useCase.GetProjectsByShortname(c.Request.URL.Query()["shortname"][0], false, "")
		} else {
			projects, err = h.useCase.GetProjectsByShortname(c.Request.URL.Query()["shortname"][0], true, userid)
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

type SetState struct {
	ProjectUUID string `json:"project_uuid"`
	State       int    `json:"state"`
}

func (h *Handler) ProjectSetState(c *gin.Context) {
	inp := new(SetState)

	if err := c.BindJSON(inp); err != nil {
		log.Panicln(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userid, err := jwt.GetFieldFromJWT(c.Request.Header["Authorization"][0], "id")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "non token"})
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

	err = h.useCase.SetStateproject(inp.State, inp.ProjectUUID, userid)

	if err != nil {
		c.JSON(400, gin.H{"message": "Something went wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"state": inp.State, "project": inp.ProjectUUID})
}

type DeleteProject struct {
	UUID string `json:"uuid"`
}

func (h *Handler) DeleteprojectById(c *gin.Context) {
	inp := new(DeleteProject)

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userid, err := jwt.GetFieldFromJWT(c.Request.Header["Authorization"][0], "id")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "non token"})
		return
	}

	err = h.useCase.DeleteprojectById(inp.UUID, userid)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": true})
		return
	}

	c.JSON(http.StatusOK, gin.H{"deleted": inp.UUID})
}

type Newproject struct {
	CategoryUUID string `json:"category_uuid"`
}

func (h *Handler) Newproject(c *gin.Context) {
	//* Создаёт новый untitled проект
	inp := new(Newproject)

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userid, err := jwt.GetFieldFromJWT(c.Request.Header["Authorization"][0], "id")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "non token"})
		return
	}

	uuid, err := h.useCase.Newproject(userid, inp.CategoryUUID)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "something went wrong"})
	}

	c.JSON(http.StatusOK, gin.H{"project": uuid})
}

func (h *Handler) LoadPhoto(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"project": "n n n"})
}

type LoadPhotoPrewiew struct {
	PhotoType   string `json:"photo_type"`
	ProjectUUID string `json:"project_uui"`
}

func (h *Handler) LoadPhotoPrewiew(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}

	inp := new(LoadPhotoPrewiew)

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userid, err := jwt.GetFieldFromJWT(c.Request.Header["Authorization"][0], "id")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "non token"})
		return
	}

	h.useCase.LoadPhoto(file, userid, inp.ProjectUUID, inp.PhotoType)
}
