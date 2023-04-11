package controllers

import (
	"my-gram/database"
	"my-gram/helpers"
	"my-gram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// CreatePhoto godoc
// @Summary Add Photo
// @Description Add a photo to a photo
// @Tags Photo
// @Accept json
// @Produce json
// @Success 201 {object} models.Photo
// @Router /photos [post]
func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Photo)
}

// UpdatePhoto godoc
// @Summary Update Photo
// @Description Update a photo using id
// @Tags Photo
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer {JWT token}"
// @Param id path int true "Id of the photo"
// @Success 200 {object} models.Photo
// @Router /photos/{id} [put]
func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.ID = uint(photoId)

	err := db.Model(&Photo).
		Where("id = ?", photoId).
		Updates(models.Photo{Title: Photo.Title, Caption: Photo.Caption, PhotoUrl: Photo.PhotoUrl}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Photo)
}

// GetPhoto godoc
// @Summary Get Photo
// @Description Show a photo using id
// @Tags Photo
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer {JWT token}"
// @Param id path int true "Id of the photo"
// @Success 200 {object} models.Photo
// @Router /photos/{id} [get]
func GetPhotoById(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	Photo.UserID = userID
	Photo.ID = uint(photoId)

	err := db.First(&Photo, "id = ?", photoId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Photo)
}

// DeletePhoto godoc
// @Summary Delete a Photo
// @Description Delete a photo using id
// @Tags Photo
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer {JWT token}"
// @Param id path int true "Id of the photo"
// @Success 200 {string} string "Photo successfully deleted"
// @Router /photos/{id} [delete]
func DeletePhotoById(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	Photo.UserID = userID
	Photo.ID = uint(photoId)

	result := db.Where("id = ?", photoId).Delete(&Photo)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Data not found",
			"message": "Data with selected id is not found",
		})
		return
	}

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Photo successfully deleted",
	})
}

// GetAllPhoto godoc
// @Summary Get All Photos
// @Description Get all photos
// @Tags Photo
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer {JWT token}"
// @Success 200 {object} models.Photo
// @Router /photos [get]
func GetAllPhotos(c *gin.Context) {
	db := database.GetDB()
	products := []models.Photo{}

	err := db.Find(&products).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, products)
}
