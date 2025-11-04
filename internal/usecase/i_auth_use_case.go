package usecase

import "github.com/B1gdawg0/real_course_review_backend/internal/dtos"

type AuthUsecase interface {
	LoginOrRegister(req *dtos.AuthRequest) (*dtos.AuthResponse, error)
}