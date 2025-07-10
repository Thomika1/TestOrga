package model

type StudyPlan struct {
	ID           int    `json:"id"`
	UserID       int    `json:"user_id"`
	Day          string `json:"day"` // ex: "2025-07-07"
	Subject      string `json:"subject"`
	StudyContent string `json:"study_content"`
	CreatedAt    string `json:"created_at"`
}

type StudyPlanResponse struct {
	ID           int    `json:"id"`
	Day          string `json:"day"` // ex: "2025-07-07"
	Subject      string `json:"subject"`
	StudyContent string `json:"study_content"`
	CreatedAt    string `json:"created_at"`
}
