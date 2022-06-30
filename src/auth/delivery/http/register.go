package http

import (
	"go-just-portfolio/src/auth"

	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc auth.UseCase) {
	h := NewHandler(uc)

	authEndpoints := router.Group("/api/auth")
	{
		authEndpoints.POST("/sign-up", h.SignUp)
		authEndpoints.POST("/sign-in", h.SignIn)
		authEndpoints.GET("/profile", h.Profile)

		authEndpoints.GET("/github/get_link", h.GetLinkGithub)
		authEndpoints.GET("/github/callback", h.GitHubCallback)
	}
}
