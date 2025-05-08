package helper

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	err = nil
	if userType != role {
		err = errors.New("Unauthorized to access this resource")
		return err
	}
	return err
}

func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	err = nil

	if userType == "ADMIN" {
		return nil
	}

	if userType == "USER" && uid == userId {
		return nil
	}

	if uid != userId {
		err = errors.New("Unauthorized to access this resource")
		return err
	}

	return nil
}
