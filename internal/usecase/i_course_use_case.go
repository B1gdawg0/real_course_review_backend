package usecase

import "github.com/B1gdawg0/real_course_review_backend/internal/dtos"

type CourseUseCase interface {
	GetAll(page, size int) ([]dtos.CourseShortResponse, int, int, int64, int, error)
	GetCourseById(id string) (*dtos.CourseFullResponse, error)
	Search(q string, limit, offset int) ([]dtos.CourseShortResponse, int, int, int64, int, error)
	CompareCoursesById(userId, first, second string) ([]dtos.CourseCompareResponse, error)
}