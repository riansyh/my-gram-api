package controllers

import (
	"my-gram/helpers"
	"my-gram/models"
	"my-gram/services"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type PhotoController interface {
	CreatePhoto(c *gin.Context)
	UpdatePhoto(c *gin.Context)
	DeletePhoto(c *gin.Context)
	GetPhotoById(c *gin.Context)
	GetAllPhoto(c *gin.Context)
}

type photoController struct {
	service services.PhotoService
}

func NewPhotoController(service services.PhotoService) *photoController {
	return &photoController{service: service}
}

// CreatePhoto godoc
// @Summary Add Photo
// @Description Add a photo to a photo
// @Tags Photo
// @Accept json
// @Produce json
// @Success 201 {object} models.Photo
// @Router /photos [post]
func (s *photoController) CreatePhoto(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	userID := c.MustGet("userData").(jwt.MapClaims)["id"].(float64)

	Photo := models.Photo{}
	Photo.UserID = uint(userID)

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	err := s.service.CreatePhoto(&Photo)

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
func (s *photoController) UpdatePhoto(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	userID := c.MustGet("userData").(jwt.MapClaims)["id"].(float64)

	Photo := models.Photo{}
	photoId, _ := strconv.Atoi(c.Param("photoId"))

	Photo.UserID = uint(userID)
	Photo.ID = uint(photoId)

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	err := s.service.UpdatePhoto(&Photo, uint(userID), uint(photoId))

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
func (s *photoController) GetPhotoById(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)

	Photo := models.Photo{}
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	Photo.UserID = userID
	Photo.ID = uint(photoId)

	socmed, err := s.service.GetPhotoById(Photo.ID, userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, socmed)
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
func (s *photoController) DeletePhoto(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	userID := c.MustGet("userData").(jwt.MapClaims)["id"].(float64)

	Photo := models.Photo{}
	photoId, _ := strconv.Atoi(c.Param("photoId"))

	Photo.UserID = uint(userID)
	Photo.ID = uint(photoId)

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	err := s.service.DeletePhoto(&Photo, uint(userID), uint(photoId))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
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
func (s *photoController) GetAllPhoto(c *gin.Context) {
	result, err := s.service.GetAllPhoto()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
