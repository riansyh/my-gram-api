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

type SocialMediaController interface {
	CreateSocialMedia(c *gin.Context)
	UpdateSocialMedia(c *gin.Context)
	DeleteSocialMedia(c *gin.Context)
	GetSocialMediaById(c *gin.Context)
	GetAllSocialMedia(c *gin.Context)
}

type socialMediaController struct {
	service services.SocialMediaService
}

func NewSocialMediaController(service services.SocialMediaService) *socialMediaController {
	return &socialMediaController{service: service}
}

// CreateSocialMedia godoc
// @Summary Add SocialMedia
// @Description Add a comment to a photo
// @Tags SocialMedia
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body models.SocialMedia true "Social Media"
// @Success 201 {object} models.SocialMedia
// @Router /social-medias [post]
func (s *socialMediaController) CreateSocialMedia(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	userID := c.MustGet("userData").(jwt.MapClaims)["id"].(float64)

	SocialMedia := models.SocialMedia{}
	SocialMedia.UserID = uint(userID)

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	err := s.service.CreateSocialMedia(&SocialMedia)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": SocialMedia,
	})
}

// UpdateSocialMedia godoc
// @Summary Update SocialMedia
// @Description Update a comment using id
// @Tags SocialMedia
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body models.SocialMedia true "Social Media"
// @Param id path int true "Id of the social media"
// @Success 200 {object} models.SocialMedia
// @Router /social-medias/{id} [put]
func (s *socialMediaController) UpdateSocialMedia(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	userID := c.MustGet("userData").(jwt.MapClaims)["id"].(float64)

	SocialMedia := models.SocialMedia{}
	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))

	SocialMedia.UserID = uint(userID)
	SocialMedia.ID = uint(socialMediaId)

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	err := s.service.UpdateSocialMedia(&SocialMedia, uint(userID), uint(socialMediaId))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": SocialMedia,
	})
}

// GetSocialMedia godoc
// @Summary Get SocialMedia
// @Description Show a comment using id
// @Tags SocialMedia
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Id of the social media"
// @Success 200 {object} models.SocialMedia
// @Router /social-medias/{id} [get]
func (s *socialMediaController) GetSocialMediaById(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)

	SocialMedia := models.SocialMedia{}
	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	userID := uint(userData["id"].(float64))

	SocialMedia.UserID = userID
	SocialMedia.ID = uint(socialMediaId)

	result, err := s.service.GetSocialMediaById(SocialMedia.ID, userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

// DeleteSocialMedia godoc
// @Summary Delete a SocialMedia
// @Description Delete a comment using id
// @Tags SocialMedia
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Id of the social media"
// @Success 200 {string} string "SocialMedia successfully deleted"
// @Router /social-medias/{id} [delete]
func (s *socialMediaController) DeleteSocialMedia(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	userID := c.MustGet("userData").(jwt.MapClaims)["id"].(float64)

	SocialMedia := models.SocialMedia{}
	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))

	SocialMedia.UserID = uint(userID)
	SocialMedia.ID = uint(socialMediaId)

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	err := s.service.DeleteSocialMedia(&SocialMedia, uint(userID), uint(socialMediaId))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Social media successfully deleted",
	})
}

// GetAllSocialMedia godoc
// @Summary Get All Social Medias
// @Description Get all social media
// @Tags SocialMedia
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.SocialMedia
// @Router /social-medias [get]
func (s *socialMediaController) GetAllSocialMedia(c *gin.Context) {
	result, err := s.service.GetAllSocialMedia()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
