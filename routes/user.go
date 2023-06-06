package routes

import (
	"fmt"

	"github.com/AdityaBkIIITA/gin-gorm-rest/auth"
	"github.com/AdityaBkIIITA/gin-gorm-rest/controller"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/user", controller.GetAllUsers)
	router.GET("/user/:id", controller.GetUser)
	router.POST("/user", controller.CreateUser)
	router.PATCH("/user/:id", controller.UpdateUser)
	router.DELETE("/user/:id", controller.DeleteUser)
	router.DELETE("/user", controller.DeleteAllUsers)

}

func AuthRouter(router *gin.Engine) {
	fmt.Println("Working auth!!")
	router.POST("/token", controller.Tokener)
	authRouter := router.Group("/user")
	authRouter.Use(auth.AuthMiddleware())
	{
		authRouter.GET("", controller.GetAllUsers)
		authRouter.GET("/:id", controller.GetUser)
		authRouter.POST("", controller.CreateUser)
		authRouter.PATCH("/:id", controller.UpdateUser)
		authRouter.DELETE("/:id", controller.DeleteUser)
		authRouter.DELETE("", controller.DeleteAllUsers)
	}
}
