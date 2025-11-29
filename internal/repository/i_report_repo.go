package repository

import m "github.com/B1gdawg0/real_course_review_backend/internal/model"

type ReportRepository interface{
	AddReportToReview(userid string, reviewid string, report m.ReportType) (error)
	GetAllReports() ([]m.Report, error)
}