package http

import (
	auth "go-just-portfolio/src/auth"
	"log"
	"net/http"

	gin "github.com/gin-gonic/gin"
)

type Handler struct {
	useCase auth.UseCase
}

func NewHandler(useCase auth.UseCase) *Handler {
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

func (h *Handler) SignUp(c *gin.Context) {
	inp := new(signInput)

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := h.useCase.SignUp(inp.Shortname, inp.Mail, inp.Password, inp.Fullname)
	if err != nil {
		if err == auth.ErrUserAlreadyExist {
			c.JSON(http.StatusConflict, gin.H{"conflict": "user's mail or username already exists"})
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *Handler) Profile(c *gin.Context) {

	log.Println(len(c.Request.URL.Query()["shortname"]) == 0)
	if len(c.Request.URL.Query()["shortname"]) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	user, err := h.useCase.Profile(c.Request.URL.Query()["shortname"][0])
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
	}

	c.JSON(http.StatusOK, gin.H{
		"mail":      user.Mail,
		"avatar":    user.Avatar,
		"shortname": user.Shortname,
		"fullname":  user.Fullname,
	})
}

func (h *Handler) SignIn(c *gin.Context) {
	inp := new(signInput)

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := h.useCase.SignIn(inp.Mail, inp.Password)

	if err != nil {
		if err == auth.ErrUserNotFound {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
