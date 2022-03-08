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
		projectEndpoints.POST("/new", h.Newproject)                        // * (ok)
		projectEndpoints.PUT("/photo", h.LoadPhotoPrewiew)                 // * (ok)
		projectEndpoints.PUT("/state", h.ProjectSetState)                  // * (ok)
		projectEndpoints.PUT("/edit/title", func(c *gin.Context) {})       // ! (no)
		projectEndpoints.PUT("/edit/description", func(c *gin.Context) {}) // ! (no)
		projectEndpoints.PUT("/edit/tag", func(c *gin.Context) {})         // ! (no)
		projectEndpoints.DELETE("/", h.DeleteprojectById)                  // * (ok)
		projectEndpoints.DELETE("/photo", func(c *gin.Context) {})         // ! (no)
		projectEndpoints.DELETE("/description", func(c *gin.Context) {})   // ! (no)
		projectEndpoints.DELETE("/tag", func(c *gin.Context) {})           // ! (no)
	}
}
