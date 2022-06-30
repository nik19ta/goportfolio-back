package usecase

import (
	"go-just-portfolio/models"
	"go-just-portfolio/src/auth"
)

type AuthUseCase struct {
	userRepo auth.UserRepository
}

func NewAuthUseCase(userRepo auth.UserRepository) *AuthUseCase {
	return &AuthUseCase{userRepo: userRepo}
}

func (a *AuthUseCase) SignUp(username, mail, password, fullname string) (*string, error) {
	user := &models.User{
		Shortname: username,
		Mail:      mail,
		Password:  password,
		Fullname:  fullname,
		Type:      "user",
	}

	token, err := a.userRepo.CreateUser(user)

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (a *AuthUseCase) GetHubCheckUser(github_id int64) (*string, error) {
	token, err := a.userRepo.GetHubCheckUser(github_id)

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (a *AuthUseCase) GithSignUp(username, mail, fullname string, github_id int64, avatar_url string) (*string, error) {
	user := &models.User{
		Shortname: username,
		Mail:      mail,
		Fullname:  fullname,
		Type:      "github",
		ServiceId: github_id,
		Avatar:    avatar_url,
	}

	token, err := a.userRepo.CreateUser(user)

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (a *AuthUseCase) SignIn(username, password string) (*string, error) {
	token, err := a.userRepo.GetUserToken(username, password)
	if err != nil {
		return nil, auth.ErrUserNotFound
	}

	return token, nil
}

func (a *AuthUseCase) Profile(shortname string) (*models.User, error) {
	user, err := a.userRepo.GetUserInfo(shortname)
	if err != nil {
		return nil, auth.ErrUserNotFound
	}

	return user, nil
}
