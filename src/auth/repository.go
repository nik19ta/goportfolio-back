package auth

import (
	"go-just-portfolio/models"
)

type UserRepository interface {
	CreateUser(user *models.User) (*string, error)
	GetUserToken(mail, password string) (*string, error)
	GetUserInfo(shortname string) (*models.User, error)
	GetHubCheckUser(github_id int64) (*string, error)
}
