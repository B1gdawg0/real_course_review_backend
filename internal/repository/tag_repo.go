package repository

import (
	m "github.com/B1gdawg0/real_course_review_backend/internal/model"
	"gorm.io/gorm"
)

type tagRepo struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return &tagRepo{
		db: db,
	}
}

func (t *tagRepo) GetAll() ([]m.Tag, error) {
	var tags []m.Tag
	err := t.db.Find(&tags).Error
	return tags, err
}

func (t *tagRepo) GetTagById(id string) (*m.Tag, error) {
	var tag m.Tag
	err := t.db.First(&tag, "id = ?", id).Error
	return &tag, err
}
