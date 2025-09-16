package usecase

import "github.com/B1gdawg0/real_course_review_backend/internal/dtos"



type CourseUseCase interface {
  	GetAll() ([]dtos.CourseShortResponse, error)
	GetCourseById(id string) (*dtos.CourseFullResponse, error)
}