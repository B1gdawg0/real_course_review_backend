package repository

import (
	m "github.com/B1gdawg0/real_course_review_backend/internal/model"
)

type CourseRepository interface {
    GetAll(filter m.CourseFilter) ([]m.Course, error)
    Search(q string, limit, offset int, filter m.CourseFilter) ([]m.Course, error)
    CountSearch(q string, filter m.CourseFilter) (int64, error)
    GetCourseById(id string) (*m.Course, error)
}