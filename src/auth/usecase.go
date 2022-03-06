package auth

import "go-just-portfolio/models"

type UseCase interface {
	SignUp(username, mail, password, fullname string) (*string, error)
	SignIn(username, password string) (*string, error)
	Profile(shortname string) (*models.User, error)
}
