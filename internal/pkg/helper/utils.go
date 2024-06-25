package helper

import (
	"github.com/AkbarFikri/Mooistudio-Technical-Test/internal/api/authentication/dto"
	"github.com/gin-gonic/gin"
)

func GetUserLoginData(c *gin.Context) dto.UserTokenData {
	getUser, _ := c.Get("user")
	user := getUser.(dto.UserTokenData)

	return user
}
