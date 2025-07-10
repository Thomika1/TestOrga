package model

type Exam struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Subject   string `json:"subject"`
	ExamDate  string `json:"exam_date"` // formato "2006-01-02"
	CreatedAt string `json:"created_at"`
}

type ExamResponse struct {
	ID        int    `json:"id"`
	Subject   string `json:"subject"`
	ExamDate  string `json:"exam_date"` // formato "YYYY-MM-DD"
	CreatedAt string `json:"created_at"`
}
