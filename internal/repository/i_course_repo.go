package repository

import (
	m "github.com/B1gdawg0/real_course_review_backend/internal/model"
)

type CourseRepository interface {
    GetAll(filter m.CourseFilter, limit, offset int) ([]m.Course, int64, error)
    Search(q string, limit, offset int, filter m.CourseFilter) ([]m.Course, error)
    CountSearch(q string, filter m.CourseFilter) (int64, error)
    GetCourseById(id string) (*m.Course, error)
	UpdateCourseRecStatus(id string, recStatus bool) error
	UpdateCourse(course *m.Course) error
    BulkCreateCourses(courses []m.Course) error
	UpdateCourseById(id string, updatedCourse *m.Course) error
}