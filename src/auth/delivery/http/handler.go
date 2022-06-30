package http

import (
	"encoding/json"
	"fmt"
	"go-just-portfolio/pkg/config"
	auth "go-just-portfolio/src/auth"
	"net/http"

	"go-just-portfolio/models"

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

func (h *Handler) SignUp(c *gin.Context) {
	inp := new(models.SignInp)

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
	inp := new(models.SignInp)

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

func (h *Handler) GetLinkGithub(c *gin.Context) {
	link := "https://github.com/login/oauth/authorize?client_id=" + config.GetConfig().GITHUB_CLIENT_ID

	c.Redirect(http.StatusFound, link)
}

func (h *Handler) GitHubCallback(c *gin.Context) {
	code := c.Query("code")

	reqURL := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s", config.GetConfig().GITHUB_CLIENT_ID, config.GetConfig().GITHUB_SECRET, code)
	req, err := http.NewRequest(http.MethodPost, reqURL, nil)
	req.Header.Set("accept", "application/json")
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true})
	}
	defer res.Body.Close()

	var t models.OAuthAccessResponse
	if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true})
	}

	request_user_data, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true})
	}
	request_user_data.Header.Set("Authorization", "token "+t.AccessToken)
	response_user_data, err := http.DefaultClient.Do(request_user_data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true})
	}
	defer response_user_data.Body.Close()

	var responseUserData models.OAuthUser
	if err := json.NewDecoder(response_user_data.Body).Decode(&responseUserData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true})
	}

	request_user_emails, err := http.NewRequest("GET", "https://api.github.com/user/emails", nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true})
	}

	request_user_emails.Header.Set("Authorization", "token "+t.AccessToken)
	response_user_email, err := http.DefaultClient.Do(request_user_emails)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true})
	}
	defer response_user_email.Body.Close()

	var responseUserEmail []models.OAuthEmail
	if err := json.NewDecoder(response_user_email.Body).Decode(&responseUserEmail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true})
		return
	}
	token, err := h.useCase.GetHubCheckUser(responseUserData.Id)
	if err == nil {
		c.Redirect(http.StatusFound, config.GetConfig().REDIRECT+"?code="+*token)
		return
	}
	ttt, e := h.useCase.GithSignUp(responseUserData.Login, responseUserEmail[0].Email, responseUserData.Name, responseUserData.Id, responseUserData.AvatarUrl)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true})
		return
	}
	c.Redirect(http.StatusFound, config.GetConfig().REDIRECT+"?code="+*ttt)
	return
}
