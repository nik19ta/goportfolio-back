package categories

import "errors"

var (
	ErrprojectNotFound    = errors.New("project not found")
	ErrInvalidAccessToken = errors.New("invalid access token")
	DataBaseError         = errors.New("database error")
)
