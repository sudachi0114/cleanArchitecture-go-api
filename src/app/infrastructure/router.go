package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/sudachi0114/cleanArchitecture-go-api/src/app/interfaces/controllers"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	userController := controllers.NewUserController(NewSqlHandler())

	router.POST("/users", func(c *gin.Context) {
		userController.Create(c)
	})

	Router = router
}
