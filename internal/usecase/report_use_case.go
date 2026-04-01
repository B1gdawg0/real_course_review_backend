package usecase

import (
	"errors"

	"github.com/B1gdawg0/real_course_review_backend/internal/dtos"
	m "github.com/B1gdawg0/real_course_review_backend/internal/model"
	repo "github.com/B1gdawg0/real_course_review_backend/internal/repository"
	"github.com/B1gdawg0/real_course_review_backend/internal/utils"
)

type reportUseCase struct {
	repo repo.ReportRepository
	reviewUseCase repo.ReviewRepository
	courseUseCase repo.CourseRepository
}

func NewReportUseCase(repo repo.ReportRepository, reviewUseCase repo.ReviewRepository, courseUseCase repo.CourseRepository) ReportUseCase {
	return &reportUseCase{repo: repo, reviewUseCase: reviewUseCase, courseUseCase: courseUseCase}
}

func (r *reportUseCase) AddReportToReview(rq dtos.ReportRequest) error {
	// TODO: add check review. Prevent duplicate report (don't know. maybe important)

	if rq.ReportType < 1 || rq.ReportType > 4 {
		return errors.New("invalid report type")
	}

	if err := r.repo.AddReportToReview(rq.UserID, rq.ReviewID, rq.ReportType, rq.ReportReason); err != nil {
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
			ReviewMessage: rep.Review.Description,
			ReportType: int(rep.ReportType),
			CourseID:   rep.Review.CourseID,
			CourseName: rep.Review.Course.Name,
			ReportReason: rep.Reason,
			CreatedAt: rep.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return result, nil
}

func (r *reportUseCase) SolveReport(id string, reviewID string, action string, reason string) error {
	if action != "SOLVE" && action != "REJECT" {
		return errors.New("invalid action")
	}

	if action == "REJECT" {
		return r.repo.RejectReport(id)
	}

	if err := r.reviewUseCase.UpdateReviewRecStatus(reviewID, false); err != nil {
		return err
	}

	if err := r.repo.SolveReport(reviewID, reason); err != nil {
		return err
	}

	review, err := r.reviewUseCase.GetReviewById(reviewID)
	if err != nil {
		return err
	}

	course, err := utils.ReduceRateCourseReview(&review.Course, review)
	if err != nil {
		return err
	}

	reviewTagMap := make(map[string]struct{})
	for _, t := range review.Tags {
		reviewTagMap[t.ID] = struct{}{}
	}

	filtered := []m.Tag{}
	for _, tag := range course.Tags {
		if _, exists := reviewTagMap[tag.ID]; !exists {
			filtered = append(filtered, tag)
		}
	}

	course.Tags = filtered

	if err := r.courseUseCase.UpdateCourse(course); err != nil {
		return err
	}

	return nil
}