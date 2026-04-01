package usecase

import (
	"github.com/B1gdawg0/real_course_review_backend/internal/dtos"
	m "github.com/B1gdawg0/real_course_review_backend/internal/model"
)

type CourseUseCase interface {
	GetAll(page, size int, filter m.CourseFilter) ([]dtos.CourseShortResponse, int, int, int64, int, error)
	GetCourseById(id string) (*dtos.CourseFullResponse, error)
	Search(q string, limit, offset int, filter m.CourseFilter) ([]dtos.CourseShortResponse, int, int, int64, int, error)
	CompareCoursesById(userId, first, second string) ([]dtos.CourseCompareResponse, error)
	GetAISummaryWithN8N(firstCourseId, secondCourseId string) (*dtos.CourseAISummaryResponse, error)
	UpdateCourseRecStatus(id string, recStatus bool) error
	BulkCreateCourses(rows [][]string) (*dtos.BulkCreateCoursesResponse, error)
	UpdateCourseById(id string, updatedCourse *dtos.UpdateCourseRequest) error
}