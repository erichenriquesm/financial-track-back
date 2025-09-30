package route

import (
	"financial-track/controller"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	userController := controller.NewUserController()
	api := r.Group("/")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", userController.RegisterUser)
		}
	}
}
