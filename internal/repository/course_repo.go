package repository

import (
	m "github.com/B1gdawg0/real_course_review_backend/internal/model"
	"gorm.io/gorm"
)

type courseRepo struct {
	db *gorm.DB
}

func NewClassRepository(db *gorm.DB) CourseRepository {
	return &courseRepo{
		db: db,
	}
}

func (c *courseRepo) GetAll() ([]m.Course, error) {
	var courses []m.Course
	err := c.db.Find(&courses).Error
	return courses, err
}

func (c *courseRepo) GetCourseById(id string) (*m.Course, error) {
	var course m.Course
	err := c.db.Preload("Professors").Where("id = ?", id).First(&course).Error
	return &course, err
}
