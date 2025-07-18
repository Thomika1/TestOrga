package usecase

import (
	"fmt"

	"github.com/Thomika1/TestOrga/model"
	"github.com/Thomika1/TestOrga/repository"
)

type ExamUsecase struct {
	repository repository.ExamRepository
}

func NewExamUsecase(repo repository.ExamRepository) ExamUsecase {
	return ExamUsecase{repository: repo}
}

func (eu *ExamUsecase) RegisterExam(exam model.Exam) (model.Exam, error) {
	examId, examdate, err := eu.repository.RegisterExam(exam)
	if err != nil {
		fmt.Println(err)
		return model.Exam{}, err
	}

	exam.ID = examId
	exam.CreatedAt = examdate

	return exam, nil

}
