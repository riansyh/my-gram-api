package router

import (
	"my-gram/controllers"
	"my-gram/middlewares"

	_ "my-gram/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// @title MyGram API
// @version 1.0
// @description This service allows users to store their photos and make comments on other users' photos. It was created as a final project for FGA Scalable Web Service with Golang.
// @contact.name Rian Febriansyah
// @contact.email rianfebriansyah22@gmail.com
// @contact.website https://riansyh.tech
// @host localhost:8081
// @BasePath /
func StartApp() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	userRouter := r.Group("/users")
	{
		// Create
		userRouter.POST("/register", controllers.UserRegister)
		// Login
		userRouter.POST("/login", controllers.UserLogin)
	}

	socialMediaRouter := r.Group("/social-medias")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		// Create
		socialMediaRouter.POST("/", controllers.CreateSocialMedia)
		// Read All
		socialMediaRouter.GET("/", controllers.GetAllSocialMedias)
		// Read
		socialMediaRouter.GET("/:socialMediaId", controllers.GetSocialMediaById)
		// Update
		socialMediaRouter.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
		// Delete
		socialMediaRouter.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMediaById)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		// Create
		photoRouter.POST("/", controllers.CreatePhoto)
		// Read All
		photoRouter.GET("/", controllers.GetAllPhotos)
		// Read
		photoRouter.GET("/:photoId", controllers.GetPhotoById)
		// Update
		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		// Delete
		photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhotoById)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		// Create
		commentRouter.POST("/", controllers.CreateComment)
		// Read All
		commentRouter.GET("/", controllers.GetAllComments)
		// Read
		commentRouter.GET("/:commentId", controllers.GetCommentById)
		// Update
		commentRouter.PUT("/:commentId", middlewares.CommentAuthorization(), controllers.UpdateComment)
		// Delete
		commentRouter.DELETE("/:commentId", middlewares.CommentAuthorization(), controllers.DeleteCommentById)
	}

	return r
}
