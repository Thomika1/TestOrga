package controllers

import (
	"net/http"

	"github.com/Thomika1/TestOrga/model"
	"github.com/Thomika1/TestOrga/usecase"
	"github.com/gin-gonic/gin"
)

type examController struct {
	examUsecase usecase.ExamUsecase
}

func NewExamController(usecase usecase.ExamUsecase) examController {
	return examController{
		examUsecase: usecase,
	}
}

func (e *examController) RegisterExam(ctx *gin.Context) {
	var exam model.Exam

	err := ctx.BindJSON(&exam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid format"})
		return
	}

	user_id, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "usuário não autenticado"})
		return
	}

	exam.UserID = user_id.(int)

	insertedExam, err := e.examUsecase.RegisterExam(exam)
	var examResponse = model.ExamResponse{
		ID:        insertedExam.ID,
		Subject:   insertedExam.Subject,
		ExamDate:  insertedExam.ExamDate,
		Topics:    insertedExam.Topics,
		CreatedAt: insertedExam.CreatedAt,
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Exam inserted", "exam": examResponse})

}

func (e *examController) GetExams(ctx *gin.Context) {
	user_id, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "usuário não autenticado"})
		return
	}

	exams, err := e.examUsecase.GetExams(user_id.(int))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Exams selected", "exams": exams})
}
