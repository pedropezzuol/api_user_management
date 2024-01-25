package routes

import "github.com/gin-gonic/gin"

func InitializeRouter() {
	router := gin.Default()

	InitializeRoutes(router)

	router.Run(":8080")
}
