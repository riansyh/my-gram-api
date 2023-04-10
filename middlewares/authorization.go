package middlewares

import (
	"my-gram/database"
	"my-gram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		Id, err := strconv.Atoi(c.Param("socialMediaId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid parameter",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		userRole := userData["role"]
		SocialMedia := models.SocialMedia{}

		err = db.Select("user_id").First(&SocialMedia, uint(Id)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "data doesn't exists",
			})
			return
		}

		if SocialMedia.UserID != userID && userRole != "admin" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to access this data",
			})
			return
		}

		c.Next()
	}
}

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		Id, err := strconv.Atoi(c.Param("photoId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid parameter",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		userRole := userData["role"]
		Photo := models.Photo{}

		err = db.Select("user_id").First(&Photo, uint(Id)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "data doesn't exists",
			})
			return
		}

		if Photo.UserID != userID && userRole != "admin" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to access this data",
			})
			return
		}

		c.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		Id, err := strconv.Atoi(c.Param("commentId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid parameter",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		userRole := userData["role"]
		Comment := models.Comment{}

		err = db.Select("user_id").First(&Comment, uint(Id)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "data doesn't exists",
			})
			return
		}

		if Comment.UserID != userID && userRole != "admin" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to access this data",
			})
			return
		}

		c.Next()
	}
}
