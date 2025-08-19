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

func (er *ExamRepository) GetsExams(user_id int) ([]model.Exam, error) {
	query, err := er.connection.Prepare("SELECT * FROM exams where user_id=$1")
	if err != nil {
		fmt.Println(err)
		return []model.Exam{}, err
	}

	rows, err := query.Query(user_id)
	if err != nil {
		fmt.Println(err)
		return []model.Exam{}, err
	}

	var examList []model.Exam
	var examObj model.Exam

	for rows.Next() {
		err = rows.Scan(
			&examObj.ID,
			&examObj.UserID,
			&examObj.Subject,
			&examObj.ExamDate,
			&examObj.CreatedAt,
			&examObj.Topics,
		)
		if err != nil {
			fmt.Println(err)
			return []model.Exam{}, err
		}
		examList = append(examList, examObj)
	}

	query.Close()
	return examList, nil
}

func (er *ExamRepository) UpdateExam(user_id int, newExam model.Exam) (model.Exam, error) {
	var exam model.Exam

	query, err := er.connection.Prepare("UPDATE exams SET subject=$1, exam_date=$2, topics=$3 WHERE user_id=$4 AND id=$5")
	if err != nil {
		fmt.Println(err)
		return model.Exam{}, err
	}

	_, err = query.Exec(exam.Subject, exam.ExamDate, exam.Topics, exam.UserID, exam.ID)
	if err != nil {
		fmt.Println(err)
		return model.Exam{}, err
	}

	return exam, nil
}
