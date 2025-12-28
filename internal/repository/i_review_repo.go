package repository

import m "github.com/B1gdawg0/real_course_review_backend/internal/model"

type ReviewRepository interface {
	GetReviewsByCourseId(userid string, id string, limit, offset int) ([]m.Review, error)
	// GetReviewsByUserId(id string, limit, offset int) ([]m.Review, error)
	VerifyReviewById(id string)(bool, error)
	CountReviewsByID(id string, section string) (int64, error)
	CreateReview(review *m.Review, course *m.Course) error
	GetHotReviewsByCourseID(userid string, id string, limit, offset int) ([]m.Review, error)
}