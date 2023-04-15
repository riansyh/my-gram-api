package controllers

import (
	"my-gram/models"
	"my-gram/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

type jwtResponse struct {
	Jwt string `json:"jwt"`
}

type userController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *userController {
	return &userController{
		userService: userService,
	}
}

// UserRegister godoc
// @Summary User registration
// @Description Register user using email
// @Tags User
// @Accept json
// @Produce json
// @Param request body models.User true "register"
// @Success 201 {object} models.User
// @Router /users/register [post]
func (uc *userController) Register(c *gin.Context) {
	user := &models.User{}

	if err := c.ShouldBind(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	createdUser, err := uc.userService.Register(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       createdUser.ID,
		"email":    createdUser.Email,
		"username": createdUser.Username,
		"age":      createdUser.Age,
	})
}

// UserLogin godoc
// @Summary Login using user account
// @Description Authenticate user using email and password
// @Tags User
// @Accept json
// @Produce json
// @Param request body models.User true "login info, email and password"
// @Success 200 {object} jwtResponse
// @Router /users/login [post]
func (uc *userController) Login(c *gin.Context) {
	user := &models.User{}

	if err := c.ShouldBind(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	token, err := uc.userService.Login(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
