package http

import (
	"go-just-portfolio/src/categories"

	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc categories.UseCase) {
	h := NewHandler(uc)

	categoriesEndpoints := router.Group("/api/categories")
	{
		//* Not Auth
		categoriesEndpoints.GET("/", h.GetCategoriesByShortname) // * (ok)
		//* Auth
		categoriesEndpoints.POST("/create", h.NewCategory)          // * (ok)
		categoriesEndpoints.DELETE("/delete", h.DeleteCategoryById) // * (ok)
		categoriesEndpoints.PUT("/edit", h.UpdateCategory)          // * (ok)
	}
}
