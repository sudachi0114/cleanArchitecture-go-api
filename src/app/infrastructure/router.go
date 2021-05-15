package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sudachi0114/cleanArchitecture-go-api/src/app/interfaces/controllers"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"healthz": "check ok",
		})
	})

	userController := controllers.NewUserController(NewSqlHandler())
	router.POST("/users", func(c *gin.Context) {
		userController.Create(c)
	})
	router.GET("/users", func(c *gin.Context) {
		userController.List(c)
	})

	Router = router
}
