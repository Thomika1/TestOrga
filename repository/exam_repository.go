package repository

import (
	"database/sql"
	"fmt"

	"github.com/Thomika1/TestOrga/model"
)

type ExamRepository struct {
	connection *sql.DB
}

func NewExamRepository(connection *sql.DB) ExamRepository {
	return ExamRepository{
		connection: connection,
	}
}

func (er *ExamRepository) RegisterExam(exam model.Exam) (int, string, error) {
	var id int
	var date string
	query, err := er.connection.Prepare("INSERT INTO exams(user_id,subject,topics,exam_date) VALUES($1,$2,$3,$4) RETURNING id, created_at")
	if err != nil {
		fmt.Println(err)
		return 0, "", err
	}
	err = query.QueryRow(exam.UserID, exam.Subject, exam.Topics, exam.ExamDate).Scan(&id, &date)
	if err != nil {
		fmt.Println(err)
		return 0, "", err
	}
	query.Close()
	return id, date, nil

}
