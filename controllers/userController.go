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

func (p *userController) GetUser(ctx *gin.Context) {
	user := []model.User{
		{
			ID:           1,
			Email:        "thomazmbonfim@gmail.com",
			PasswordHash: "penislongo",
			CreatedAt:    "10-07-2025",
		},
	}
	ctx.JSON(http.StatusOK, user)
}
