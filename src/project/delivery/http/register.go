package http

import (
	project "go-just-portfolio/src/project"

	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc project.UseCase) {
	h := NewHandler(uc)

	authEndpoints := router.Group("/api/project")
	{
		//* Not Auth
		authEndpoints.GET("/id", h.GetProjectById)
		authEndpoints.GET("/user", h.GetProjectsByShortname)
		//* Auth
		authEndpoints.POST("/new", h.Newproject)
		authEndpoints.DELETE("/", h.DeleteprojectById)

		authEndpoints.PUT("/photo", h.LoadPhoto)
		authEndpoints.PUT("/prewiew", h.LoadPhotoPrewiew)

		authEndpoints.PUT("/state", h.ProjectSetState)
	}
}
