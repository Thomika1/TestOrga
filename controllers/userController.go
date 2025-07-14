package controllers

import (
	"net/http"

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
