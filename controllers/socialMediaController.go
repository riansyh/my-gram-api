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

// CreateSocialMedia godoc
// @Summary Add SocialMedia
// @Description Add a comment to a photo
// @Tags SocialMedia
// @Accept json
// @Produce json
// @Success 201 {object} models.SocialMedia
// @Router /social-medias [post]
func CreateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	SocialMedia := models.SocialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID

	err := db.Debug().Create(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, SocialMedia)
}

// UpdateSocialMedia godoc
// @Summary Update SocialMedia
// @Description Update a comment using id
// @Tags SocialMedia
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer {JWT token}"
// @Param id path int true "Id of the social media"
// @Success 200 {object} models.SocialMedia
// @Router /social-medias/{id} [put]
func UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	SocialMedia := models.SocialMedia{}

	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID
	SocialMedia.ID = uint(socialMediaId)

	err := db.Model(&SocialMedia).Where("id = ?", socialMediaId).Updates(models.SocialMedia{Name: SocialMedia.Name, SocialMediaUrl: SocialMedia.SocialMediaUrl}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SocialMedia)
}

// GetSocialMedia godoc
// @Summary Get SocialMedia
// @Description Show a comment using id
// @Tags SocialMedia
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer {JWT token}"
// @Param id path int true "Id of the social media"
// @Success 200 {object} models.SocialMedia
// @Router /social-medias/{id} [get]
func GetSocialMediaById(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	SocialMedia := models.SocialMedia{}

	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	userID := uint(userData["id"].(float64))

	SocialMedia.UserID = userID
	SocialMedia.ID = uint(socialMediaId)

	err := db.First(&SocialMedia, "id = ?", socialMediaId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SocialMedia)
}

// DeleteSocialMedia godoc
// @Summary Delete a SocialMedia
// @Description Delete a comment using id
// @Tags SocialMedia
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer {JWT token}"
// @Param id path int true "Id of the social media"
// @Success 200 {string} string "SocialMedia successfully deleted"
// @Router /social-medias/{id} [delete]
func DeleteSocialMediaById(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	SocialMedia := models.SocialMedia{}

	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	userID := uint(userData["id"].(float64))

	SocialMedia.UserID = userID
	SocialMedia.ID = uint(socialMediaId)

	result := db.Where("id = ?", socialMediaId).Delete(&SocialMedia)

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
		"message": "Social Media successfully deleted",
	})
}

// GetAllSocialMedia godoc
// @Summary Get All Social Medias
// @Description Get all social media
// @Tags SocialMedia
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer {JWT token}"
// @Success 200 {object} models.SocialMedia
// @Router /social-medias [get]
func GetAllSocialMedias(c *gin.Context) {
	db := database.GetDB()
	products := []models.SocialMedia{}

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
