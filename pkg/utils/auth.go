package utils

import (
	"errors"
	"go-just-portfolio/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func GetUserIdFromJWT(c *gin.Context) (*string, error) {
	if len(c.Request.Header["Authorization"]) == 0 {
		return nil, errors.New("Unauthorized")
	}

	userid, err := jwt.GetFieldFromJWT(c.Request.Header["Authorization"][0], "id")

	if err != nil {
		return nil, errors.New("Unauthorized")
	}

	return userid, nil
}
