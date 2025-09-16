package repository

import (
	m "github.com/B1gdawg0/real_course_review_backend/internal/model"
)

type CourseRepository interface {
	GetAll() ([]m.Course, error)
	GetCourseById(id string) (*m.Course,error)
}