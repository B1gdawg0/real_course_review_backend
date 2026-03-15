package usecase

import "github.com/B1gdawg0/real_course_review_backend/internal/dtos"

type ReportUseCase interface{
	AddReportToReview(dtos.ReportRequest) (error)
	GetAllReports() ([]dtos.ReportShortResponse, error)
	SolveReport(id string, reviewID string, action string, reason string) error
}