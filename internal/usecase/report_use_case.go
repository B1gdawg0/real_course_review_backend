package usecase

import (
	"errors"

	"github.com/B1gdawg0/real_course_review_backend/internal/dtos"
	repo "github.com/B1gdawg0/real_course_review_backend/internal/repository"
)

type reportUseCase struct {
	repo repo.ReportRepository
}

func NewReportUseCase(repo repo.ReportRepository) ReportUseCase {
	return &reportUseCase{repo: repo}
}

func (r *reportUseCase) AddReportToReview(rq dtos.ReportRequest) error {
	// TODO: add check review. Prevent duplicate report (don't know. maybe important)

	if rq.ReportType < 1 || rq.ReportType > 4 {
		return errors.New("invalid report type")
	}

	if err := r.repo.AddReportToReview(rq.UserID, rq.ReviewID, rq.ReportType); err != nil {
		return err
	}

	return nil
}

func (r *reportUseCase) GetAllReports() ([]dtos.ReportShortResponse, error) {
	reports, err := r.repo.GetAllReports()
	if err != nil {
		return nil, err
	}

	result := make([]dtos.ReportShortResponse, 0, len(reports))

	for _, rep := range reports {
		result = append(result, dtos.ReportShortResponse{
			ID:         rep.ID,
			UserID:     rep.UserID,
			ReviewID:   rep.ReviewID,
			ReportMessage: rep.ReportType.String(),
			ReportType: int(rep.ReportType),
		})
	}

	return result, nil
}