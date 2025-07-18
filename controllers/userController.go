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

func (u *userController) GetUserByEmail(ctx *gin.Context) {

	email := ctx.Param("email")
	if email == "" {
		response := model.Response{
			Message: "Email must exist",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := u.userUsecase.GetUserById(email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})

		return
	}

	if user == nil {
		response := model.Response{
			Message: "User not found",
		}
		ctx.JSON(http.StatusNotFound, response)
	}

	ctx.JSON(http.StatusOK, user)
}

func (u *userController) UserLogin(ctx *gin.Context) {
	var req model.LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Format"})
		return
	}

	token, err := u.userUsecase.UserLogin(req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login realizado com sucesso",
		"token":   token,
	})

}
