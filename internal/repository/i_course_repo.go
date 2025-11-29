package repository

import (
	m "github.com/B1gdawg0/real_course_review_backend/internal/model"
)

type CourseRepository interface {
	GetAll() ([]m.Course, error)
	GetCourseById(id string) (*m.Course,error)
	Search(q string, limit, offset int) ([]m.Course, error)
    CountSearch(q string) (int64, error)
}