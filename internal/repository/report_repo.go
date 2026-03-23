package repository

import (
	m "github.com/B1gdawg0/real_course_review_backend/internal/model"
	"gorm.io/gorm"
)

type reportRepo struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) ReportRepository {
	return &reportRepo{
		db: db,
	}
}

func (r *reportRepo) AddReportToReview(userid string, reviewid string, report m.ReportType, reason string) error {
	newReport := m.Report{
		UserID:     userid,
		ReviewID:   reviewid,
		ReportType: report,
		Reason:    reason,
	}

	if err := r.db.Create(&newReport).Error; err != nil {
		return err
	}

	return nil
}

func (r *reportRepo) GetAllReports() ([]m.Report, error) {
	var report []m.Report
	err := r.db.Where("rec_status = ?", true).Preload("Review.Course").Find(&report).Error
	return report, err
}

func (r *reportRepo) SolveReport(reviewID string, reason string) error {
	if err := r.db.Model(&m.Report{}).Where("review_id = ?", reviewID).Updates(map[string]interface{}{
		"rec_status":   false,
		"solve_reason": reason,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (r *reportRepo) RejectReport(id string) error {
	if err := r.db.Model(&m.Report{}).Where("id = ?", id).Update("rec_status", false).Error; err != nil {
		return err
	}
	return nil
}