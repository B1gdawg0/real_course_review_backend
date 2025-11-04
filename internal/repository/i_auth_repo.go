package repository

import (
	m "github.com/B1gdawg0/real_course_review_backend/internal/model"
)

type AuthRepository interface {
	GetUserByEmail(email string) (*m.User, error)
	CreateUser(user *m.User) error
}