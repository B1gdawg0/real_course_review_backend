package repository

import m "github.com/B1gdawg0/real_course_review_backend/internal/model"

type TagRepository interface {
	GetAll() ([]m.Tag, error)
	GetTagById(id string) (*m.Tag,error)
}