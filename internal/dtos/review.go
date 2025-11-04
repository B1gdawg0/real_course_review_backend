package dtos

import "time"

type ReviewFullResponse struct {
	ID          string              `json:"id"`
	Status      string              `json:"status"`
	Description string              `json:"description"`
	AvgRate     float64             `json:"avg_rating_score"`
	Score       float64             `json:"score"`
	UpCount     int                 `json:"up_count"`
	DownCount   int                 `json:"down_count"`
	ReportCount int                 `json:"report_count"`
	IsAnonymous bool                `json:"is_anonymous"`
	Grade       string				`json:"reviewer_grade"`
	Year		string 				`json:"course_year"`
	Sec			string 				`json:"section_id"`
	Tags        []TagResponse       `json:"tags"`
	User        *UserShortResponse  `json:"user,omitempty"`
	Rate  		RatingResponse    	`json:"rating"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
}

type CreateReviewRequest struct {
    CourseID    string          `json:"course"`
	Comment     string          `json:"comment"`
	Rate        RatingResponse  `json:"rating"`
	Tags        []string       `json:"tag"`
	ProfessorID string          `json:"professor"`
	Grade       string			`json:"grade"`
	Year		string 			`json:"year"`
	Sec			string 			`json:"sec"`
	User 	    string 			`json:"user"`
}