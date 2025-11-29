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

func (r *reportRepo) AddReportToReview(userid string, reviewid string, report m.ReportType) error {
	newReport := m.Report{
		UserID:     userid,
		ReviewID:   reviewid,
		ReportType: report,
	}

	if err := r.db.Create(&newReport).Error; err != nil {
		return err
	}

	return nil
}

func (r *reportRepo) GetAllReports() ([]m.Report, error) {
	var report []m.Report
	err := r.db.Find(&report).Error
	return report, err
}