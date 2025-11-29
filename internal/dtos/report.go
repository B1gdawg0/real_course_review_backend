package dtos

import m "github.com/B1gdawg0/real_course_review_backend/internal/model"

type ReportShortResponse struct {
	ID         string `json:"id"`
	ReviewID   string `json:"review_id"`
	UserID     string `json:"user_id"`
	ReportType int    `json:"report_type"`
	ReportMessage string    `json:"report_message"`
	CreatedAt  string `json:"created_at"`
}

type ReportRequest struct {
	UserID 	   string `json:"user_id,omitempty"`
	ReviewID   string `json:"review_id" validate:"required,uuid"`
	ReportType m.ReportType    `json:"report_type" validate:"required,oneof=1 2 3 4"`
}