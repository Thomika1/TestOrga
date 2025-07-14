package routes

import (
	"database/sql"

	"github.com/Thomika1/TestOrga/controllers"
	"github.com/Thomika1/TestOrga/repository"
	"github.com/Thomika1/TestOrga/usecase"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine, db *sql.DB) {

	userRepository := repository.NewUserRepository(db)
	UserUseCase := usecase.NewUserUsecase(userRepository)
	userController := controllers.NewUserController(UserUseCase)

	router.GET("/getUsers", userController.GetUsers)

	router.POST("/register", func(c *gin.Context) {

	})

	router.POST("/login", func(c *gin.Context) {

	})

	router.POST("/exam", func(c *gin.Context) {

	})

	router.GET("/exams", func(c *gin.Context) {

	})

	router.POST("/generate-plan", func(c *gin.Context) {

	})

	router.POST("/studyplan", func(c *gin.Context) {

	})

}
