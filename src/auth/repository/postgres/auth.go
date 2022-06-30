package postgres

import (
	"go-just-portfolio/models"
	"go-just-portfolio/src/auth"

	jwt "go-just-portfolio/pkg/jwt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r UserRepository) CreateUser(user *models.User) (*string, error) {
	uuid := uuid.New().String()
	user.UUID = uuid

	is_user := models.User{}

	r.db.First(&is_user, "mail = ?", user.Mail)

	if is_user.Mail != "" {
		return nil, auth.ErrUserAlreadyExist
	}

	r.db.First(&is_user, "shortname = ?", user.Shortname)

	if is_user.Shortname != "" {
		return nil, auth.ErrUserAlreadyExist
	}

	result := r.db.Create(&user)
	token, _ := jwt.MakeJWT(user.Shortname, user.Mail, user.UUID)

	if result.Error != nil {
		return nil, auth.DataBaseError
	}

	return &token, nil
}

func (r UserRepository) GetUserToken(mail, password string) (*string, error) {
	user := models.User{}
	result := r.db.Where(&models.User{Mail: mail, Password: password, Type: "user"}).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	token, err := jwt.MakeJWT(user.Shortname, user.Mail, user.UUID)

	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (r UserRepository) GetHubCheckUser(github_id int64) (*string, error) {
	user := models.User{}
	result := r.db.Where(&models.User{ServiceId: github_id, Type: "github"}).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	token, err := jwt.MakeJWT(user.Shortname, user.Mail, user.UUID)

	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (r UserRepository) GetUserInfo(shortname string) (*models.User, error) {
	user := models.User{}
	result := r.db.Where(&models.User{Shortname: shortname}).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
