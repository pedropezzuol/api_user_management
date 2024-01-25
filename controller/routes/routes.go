package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pedropezzuol/api_test/controller"
)

func InitializeRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		//return all users
		api.GET("/users", controller.FindAllUsers)

		//return user
		api.GET("/user/:id", controller.FindUser)

		//create user
		api.POST("/user", controller.CreateUser)

		//delete user
		api.DELETE("/user/:id", controller.DeleteUser)

		//update user
		api.PUT("/user/:id", controller.UpdateUser)
	}
}
