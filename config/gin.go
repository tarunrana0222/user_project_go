package configs

import (
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func GetRouter() *gin.Engine {
	Router = gin.Default()
	Router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	Router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	return Router
}
