package project

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrprojectNotFound    = errors.New("project not found")
	ErrInvalidAccessToken = errors.New("invalid access token")
	DataBaseError         = errors.New("database error")
)
