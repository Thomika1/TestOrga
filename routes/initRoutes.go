package routes

import (
	"database/sql"

	"github.com/Thomika1/TestOrga/controllers"
	"github.com/Thomika1/TestOrga/repository"
	"github.com/Thomika1/TestOrga/usecase"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine, db *sql.DB) {
	// User
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controllers.NewUserController(userUsecase)
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", userController.GetUsers)             // GET /users/
		userRoutes.POST("/register", userController.CreateUser)  // POST /users/register
		userRoutes.GET("/:email", userController.GetUserByEmail) // GET /users/:email
		userRoutes.POST("/login", userController.UserLogin)      // POST /users/login
	}

	// Exam
	examRepository := repository.NewExamRepository(db)
	examUseCase := usecase.NewExamUsecase(examRepository)
	examController := controllers.NewExamController(examUseCase)
	examRoutes := router.Group("/exam")
	{
		examRoutes.POST("/insert", examController.RegisterExam)
	}

	// studyPlan
	router.GET("/exams", func(c *gin.Context) {

	})

	router.POST("/generate-plan", func(c *gin.Context) {

	})

	router.POST("/studyplan", func(c *gin.Context) {

	})

}
