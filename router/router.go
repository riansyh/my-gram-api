package router

import (
	"my-gram/controllers"
	"my-gram/database"
	"my-gram/middlewares"
	"my-gram/repositories"
	"my-gram/services"

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

// @securityDefinitions.apikey BearerAuth
// @in header -> Bearer
// @name Authorization

// @host localhost:8081
// @BasePath /
func StartApp() *gin.Engine {
	userRepository := repositories.NewUserRepository(database.GetDB())
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	socialMediaRepo := repositories.NewSocialMediaRepository(database.GetDB())
	socialMediaService := services.NewSocialMediaService(socialMediaRepo)
	socialMediaController := controllers.NewSocialMediaController(socialMediaService)

	photoRepo := repositories.NewPhotoRepository(database.GetDB())
	photoService := services.NewPhotoService(photoRepo)
	photoController := controllers.NewPhotoController(photoService)

	commentRepo := repositories.NewCommentRepository(database.GetDB())
	commentService := services.NewCommentService(commentRepo, photoRepo)
	commentController := controllers.NewCommentController(commentService)

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	userRouter := r.Group("/users")
	{
		// Create
		userRouter.POST("/register", userController.Register)
		// Login
		userRouter.POST("/login", userController.Login)
	}

	socialMediaRouter := r.Group("/social-medias")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		// Create
		socialMediaRouter.POST("/", socialMediaController.CreateSocialMedia)
		// Read All
		socialMediaRouter.GET("/", socialMediaController.GetAllSocialMedia)
		// Read
		socialMediaRouter.GET("/:socialMediaId", socialMediaController.GetSocialMediaById)
		// Update
		socialMediaRouter.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(), socialMediaController.UpdateSocialMedia)
		// Delete
		socialMediaRouter.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(), socialMediaController.DeleteSocialMedia)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		// Create
		photoRouter.POST("/", photoController.CreatePhoto)
		// Read All
		photoRouter.GET("/", photoController.GetAllPhoto)
		// Read
		photoRouter.GET("/:photoId", photoController.GetPhotoById)
		// Update
		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), photoController.UpdatePhoto)
		// Delete
		photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), photoController.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		// Create
		commentRouter.POST("/", commentController.CreateComment)
		// Read All
		commentRouter.GET("/", commentController.GetAllComment)
		// Read
		commentRouter.GET("/:commentId", commentController.GetCommentById)
		// Update
		commentRouter.PUT("/:commentId", middlewares.CommentAuthorization(), commentController.UpdateComment)
		// Delete
		commentRouter.DELETE("/:commentId", middlewares.CommentAuthorization(), commentController.DeleteComment)
	}

	return r
}
