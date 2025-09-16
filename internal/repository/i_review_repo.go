package repository

import m "github.com/B1gdawg0/real_course_review_backend/internal/model"

type ReviewRepository interface {
	GetReviewsByCourseId(id string, limit, offset int) ([]m.Review, error)
	GetReviewsByUserId(id string, limit, offset int) ([]m.Review, error)
	CountReviewsByID(id string, section string) (int64, error)
}