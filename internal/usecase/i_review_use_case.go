package usecase

import "github.com/B1gdawg0/real_course_review_backend/internal/dtos"


type ReviewUseCase interface {
    GetReviewsByCourseID(userid string, id string, page, size int) ([]dtos.ReviewFullResponse, int, int, int64, int, error)
    // GetReviewsByUserID(id string, page, size int) ([]dtos.ReviewFullResponse, int, int, int64, int, error)
	CreateReview(req dtos.CreateReviewRequest, userID string) (*dtos.ReviewFullResponse, error)
}