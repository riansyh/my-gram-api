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

// CreateComment godoc
// @Summary Add Comment
// @Description Add a comment to a photo
// @Tags Comment
// @Accept json
// @Produce json
// @Success 201 {object} models.Comment
// @Router /comments [post]
func CreateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}
	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID

	err := db.First(&Photo, "id = ?", Comment.PhotoId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Photo not found",
		})
		return
	}

	err = db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Comment)
}

// UpdateComment godoc
// @Summary Update Comment
// @Description Update a comment using id
// @Tags Comment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer {JWT token}"
// @Param id path int true "Id of the comment"
// @Success 200 {object} models.Comment
// @Router /comments/{id} [put]
func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Comment.ID = uint(commentId)

	err := db.Model(&Comment).
		Where("id = ?", commentId).
		Updates(models.Comment{Message: Comment.Message}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Comment)
}

// GetComment godoc
// @Summary Get Comment
// @Description Show a comment using id
// @Tags Comment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer {JWT token}"
// @Param id path int true "Id of the comment"
// @Success 200 {object} models.Comment
// @Router /comments/{id} [get]
func GetCommentById(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	Comment.UserID = userID
	Comment.ID = uint(commentId)

	err := db.First(&Comment, "id = ?", commentId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Comment)
}

// DeleteComment godoc
// @Summary Delete a Comment
// @Description Delete a comment using id
// @Tags Comment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer {JWT token}"
// @Param id path int true "Id of the comment"
// @Success 200 {string} string "Comment successfully deleted"
// @Router /comments/{id} [delete]
func DeleteCommentById(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	Comment.UserID = userID
	Comment.ID = uint(commentId)

	result := db.Where("id = ?", commentId).Delete(&Comment)

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
		"message": "Comment successfully deleted",
	})
}

// GetAllComment godoc
// @Summary Get All Comments
// @Description Get all comments
// @Tags Comment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer {JWT token}"
// @Success 200 {object} models.Comment
// @Router /comments [get]
func GetAllComments(c *gin.Context) {
	db := database.GetDB()
	products := []models.Comment{}

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
