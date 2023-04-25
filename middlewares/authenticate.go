package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tarunrana0222/user_project_go/helpers"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var xToken string
		if authArr := strings.Split(c.GetHeader("authorization"), " "); len(authArr) != 2 {
			c.JSON(401, gin.H{"errors": "authentication failed : Headers missing"})
			c.Abort()
			return
		} else {
			xToken = authArr[1]
		}
		clientId, err := helpers.ValidateToken(xToken)
		if err != nil {
			c.JSON(401, gin.H{"message": "authentication failed", "error": err.Error()})
			c.Abort()
			return
		}
		c.Set("clientId", clientId)
		c.Next()
	}
}
