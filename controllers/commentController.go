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

type CommentController interface {
	CreateComment(c *gin.Context)
	UpdateComment(c *gin.Context)
	DeleteComment(c *gin.Context)
	GetCommentById(c *gin.Context)
	GetAllComment(c *gin.Context)
}

type commentController struct {
	service services.CommentService
}

func NewCommentController(service services.CommentService) *commentController {
	return &commentController{service: service}
}

// CreateComment godoc
// @Summary Add Comment
// @Description Add a comment to a photo
// @Tags Comment
// @Accept json
// @Produce json
// @Success 201 {object} models.Comment
// @Router /comments [post]
func (s *commentController) CreateComment(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	userID := c.MustGet("userData").(jwt.MapClaims)["id"].(float64)

	Comment := models.Comment{}
	Comment.UserID = uint(userID)

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	err := s.service.CreateComment(&Comment, uint(userID))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": Comment,
	})
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
func (s *commentController) UpdateComment(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	userID := c.MustGet("userData").(jwt.MapClaims)["id"].(float64)

	Comment := models.Comment{}
	commentId, _ := strconv.Atoi(c.Param("commentId"))

	Comment.UserID = uint(userID)
	Comment.ID = uint(commentId)

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	err := s.service.UpdateComment(&Comment, uint(userID), uint(commentId))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": Comment,
	})
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
func (s *commentController) GetCommentById(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)

	Comment := models.Comment{}
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	Comment.UserID = userID
	Comment.ID = uint(commentId)

	result, err := s.service.GetCommentById(Comment.ID, userID)

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
func (s *commentController) DeleteComment(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	userID := c.MustGet("userData").(jwt.MapClaims)["id"].(float64)

	Comment := models.Comment{}
	commentId, _ := strconv.Atoi(c.Param("commentId"))

	Comment.UserID = uint(userID)
	Comment.ID = uint(commentId)

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	err := s.service.DeleteComment(&Comment, uint(userID), uint(commentId))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
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
func (s *commentController) GetAllComment(c *gin.Context) {
	result, err := s.service.GetAllComment()

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
