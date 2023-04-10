package router

import (
	"my-gram/controllers"
	"my-gram/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	socialMediaRouter := r.Group("/social-media")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		socialMediaRouter.POST("/", controllers.CreateSocialMedia)
		socialMediaRouter.GET("/", controllers.GetAllSocialMedias)
		socialMediaRouter.GET("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.GetSocialMediaById)
		socialMediaRouter.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMediaById)
	}

	return r
}
