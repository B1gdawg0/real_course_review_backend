package usecase

import "github.com/B1gdawg0/real_course_review_backend/internal/dtos"

type ProfessorUseCase interface{
	GetAll() ([]dtos.ProfessorShortResponse, error)
	GetProfessorById(id string) (*dtos.ProfessorFullResponse,error)
}