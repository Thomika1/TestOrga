package controllers

import (
	"net/http"

	"github.com/Thomika1/TestOrga/model"
	"github.com/Thomika1/TestOrga/usecase"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type userController struct {
	userUsecase usecase.UserUsecase
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
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

	password, err := HashPassword(user.PasswordHash)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	//salva a senha encriptada no campo de user
	user.PasswordHash = password

	//func CheckPasswordHash(password, hash string) bool {
	//err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	//return err == nil
	//}

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
