package http

import (
	"go-just-portfolio/models"
	"go-just-portfolio/pkg/utils"
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

func (h *Handler) GetCategoriesByShortname(c *gin.Context) {
	data, err := h.useCase.GetCategoriesByUserName(c.Request.URL.Query()["shortname"][0])

	if err != nil {
		c.JSON(400, gin.H{"error": "true"})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *Handler) UpdateCategory(c *gin.Context) {
	inp := new(models.CategoryInp)

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userid, err := utils.GetUserIdFromJWT(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	err = h.useCase.EditCategory(*userid, inp.UUID, inp.Title)

	if err != nil {
		c.JSON(400, gin.H{"error": "true"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "title": inp.Title, "uuid": inp.UUID})
}

func (h *Handler) DeleteCategoryById(c *gin.Context) {
	inp := new(models.CategoryInp)

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userid, err := utils.GetUserIdFromJWT(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	err = h.useCase.DeleteCategory(*userid, inp.UUID)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "true"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"delete": inp.UUID})
}

func (h *Handler) NewCategory(c *gin.Context) {
	inp := new(models.CategoryInp)

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userid, err := utils.GetUserIdFromJWT(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	uuid, err := h.useCase.AddCategory(*userid, inp.Title)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "something went wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"created": uuid})
}
