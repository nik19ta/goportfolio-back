package http

import (
	jwt "go-just-portfolio/pkg/jwt"
	categories "go-just-portfolio/src/categories"

	"net/http"

	gin "github.com/gin-gonic/gin"
)

type Handler struct {
	useCase categories.UseCase
}

func NewHandler(useCase categories.UseCase) *Handler {
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

func (h *Handler) GetCategoriesByShortname(c *gin.Context) {
	data, err := h.useCase.GetCategoriesByUserName(c.Request.URL.Query()["shortname"][0])

	if err != nil {
		c.JSON(400, gin.H{"error": "true"})
		return
	}

	c.JSON(http.StatusOK, data)
}

type CategoryUpdate struct {
	UUID  string `json:"uuid"`
	Title string `json:"title"`
}

func (h *Handler) UpdateCategory(c *gin.Context) {
	inp := new(CategoryUpdate)

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userid, err := jwt.GetFieldFromJWT(c.Request.Header["Authorization"][0], "id")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "non token"})
		return
	}

	err = h.useCase.EditCategory(userid, inp.UUID, inp.Title)

	if err != nil {
		c.JSON(400, gin.H{"error": "true"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"statue": true, "title": inp.Title})
}

type DeleteCategory struct {
	UUID string `json:"uuid"`
}

func (h *Handler) DeleteCategoryById(c *gin.Context) {
	inp := new(DeleteCategory)

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userid, err := jwt.GetFieldFromJWT(c.Request.Header["Authorization"][0], "id")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "non token"})
		return
	}

	err = h.useCase.DeleteCategory(userid, inp.UUID)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "true"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"delete": inp.UUID})
}

type NewCategory struct {
	Title string `json:"title"`
}

func (h *Handler) NewCategory(c *gin.Context) {
	//* Создаёт новый untitled проект
	inp := new(NewCategory)

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userid, err := jwt.GetFieldFromJWT(c.Request.Header["Authorization"][0], "id")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "non token"})
		return
	}

	uuid, err := h.useCase.AddCategory(userid, inp.Title)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "something went wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"created": uuid})
}
