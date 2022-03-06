package http

import (
	"go-just-portfolio/src/categories"

	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc categories.UseCase) {
	h := NewHandler(uc)

	authEndpoints := router.Group("/api/categories")
	{
		//* Not Auth
		authEndpoints.GET("/", h.GetCategoriesByShortname)
		//* Auth
		authEndpoints.POST("/create", h.NewCategory)
		authEndpoints.DELETE("/delete", h.DeleteCategoryById)
		authEndpoints.PUT("/edit", h.UpdateCategory)
	}
}
