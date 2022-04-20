package http

import (
	project "go-just-portfolio/src/project"

	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc project.UseCase) {
	h := NewHandler(uc)

	projectEndpoints := router.Group("/api/project")
	{
		//* Not Auth
		projectEndpoints.GET("/id", h.GetProjectById)           // ! (no)
		projectEndpoints.GET("/user", h.GetProjectsByShortname) // * (ok)
		//* Auth
		projectEndpoints.POST("/new", h.NewProject)                           // * (ok)
		projectEndpoints.POST("/create/description", func(c *gin.Context) {}) // ! (no)
		projectEndpoints.POST("/create/tag", func(c *gin.Context) {})         // ! (no)
		projectEndpoints.PUT("/photo", h.LoadPhotoPrewiew)                    // * (ok)
		projectEndpoints.PUT("/state", h.ProjectSetState)                     // * (ok)
		projectEndpoints.PUT("/title", h.RenameProject)                       // ~ (now)
		projectEndpoints.DELETE("/photo", func(c *gin.Context) {})            // ! (no)
		projectEndpoints.DELETE("/description", func(c *gin.Context) {})      // ! (no)
		projectEndpoints.DELETE("/tag", func(c *gin.Context) {})              // ! (no)
		projectEndpoints.DELETE("", h.DeleteprojectById)                      // * (ok)
	}
}
