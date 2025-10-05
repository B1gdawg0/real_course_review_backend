package usecase

import "github.com/B1gdawg0/real_course_review_backend/internal/dtos"

type TagUseCase interface{
	GetAll() ([]dtos.TagResponse, error)
	GetTagById(id string) (*dtos.TagResponse,error)
}