package auth

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrUserAlreadyExist   = errors.New("user already exist")
	ErrInvalidAccessToken = errors.New("invalid access token")
	DataBaseError         = errors.New("database error")
)
