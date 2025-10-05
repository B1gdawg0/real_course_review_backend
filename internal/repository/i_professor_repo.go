package repository

import (
	m "github.com/B1gdawg0/real_course_review_backend/internal/model"
)

type ProfessorRepository interface {
	GetAll() ([]m.Professor, error)
	GetProfessorById(id string) (*m.Professor,error)
}