package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tarunrana0222/user_project_go/controllers"
)

func ClientRoutes(c *gin.Engine) {
	client := c.Group("/client")
	{
		client.GET("/all", controllers.GetAllClients())
		client.GET("/token", controllers.GetAuthToken())
	}
}
