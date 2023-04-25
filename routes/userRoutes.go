package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tarunrana0222/user_project_go/controllers"
	"github.com/tarunrana0222/user_project_go/middlewares"
)

func UserRoutes(router *gin.Engine) {
	user := router.Group("/user", middlewares.Authenticate())
	{
		user.GET("/all", controllers.GetAllUsers())
		user.GET("/:userId", controllers.GetSingleUsers())
		user.POST("/create", controllers.CreateUser())
		user.PATCH("/update", controllers.UpdateUser())
		user.DELETE("/delete/:userId", controllers.DeleteUser())

	}
}
