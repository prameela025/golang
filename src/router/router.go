package router

import (
	"net/http"
	"sample/src/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter(UserController *controller.UserController) *gin.Engine {
	router := gin.Default()

	router.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello World")
	})

	userRouter := router.Group("/user")
	userRouter.GET("", UserController.FindAll)
	userRouter.GET("/:id", UserController.FindById)
	userRouter.POST("", UserController.Save)
	userRouter.PUT("", UserController.Update)
	userRouter.DELETE("/:id", UserController.Delete)

	return router
}