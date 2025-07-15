package controllers

import (
	"net/http"

	"github.com/Thomika1/TestOrga/model"
	"github.com/Thomika1/TestOrga/usecase"
	"github.com/gin-gonic/gin"
)

type userController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(usecase usecase.UserUsecase) userController {
	return userController{
		userUsecase: usecase,
	}
}

func (u *userController) GetUsers(ctx *gin.Context) {
	users, err := u.userUsecase.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, users)
}

func (u *userController) CreateUser(ctx *gin.Context) {
	var user model.User

	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	insertedUser, err := u.userUsecase.CreateUser(user)
	userResponse := model.UserResponse{
		ID:    insertedUser.ID,
		Email: insertedUser.Email,
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusCreated, userResponse)
}
