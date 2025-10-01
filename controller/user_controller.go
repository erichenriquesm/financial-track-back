package controller

import (
	"financial-track/model"
	"financial-track/repository"
	"financial-track/usecase"
	"financial-track/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

var userRepository *repository.UserRepository = repository.NewUserRepository()
var userUseCase *usecase.UserUseCase = usecase.NewUserUseCase(userRepository)

func (uc *UserController) RegisterUser(c *gin.Context) {
	var input model.RegisterUserInput
	errs := utils.ValidateJSON(c, &input)
	if errs != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Fields are invalid or missing",
			"errors":  errs,
		})
		return
	}

	createdUser, err := userUseCase.RegisterUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    createdUser,
	})
}

func (uc *UserController) LoginUser(c *gin.Context) {
	var input model.LoginUserInput
	errs := utils.ValidateJSON(c, &input)
	if errs != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Fields are invalid or missing",
			"errors":  errs,
		})
		return
	}

	token, err := userUseCase.LoginUser(input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}
