package repository

import (
	"regexp"

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
    err := c.db.Preload("Tags").Find(&courses).Error
    return courses, err
}

func (c *courseRepo) GetCourseById(id string) (*m.Course, error) {
	var course m.Course
	err := c.db.Preload("Tags").Preload("Professors").Where("id = ?", id).First(&course).Error
	return &course, err
}

func (c *courseRepo) Search(q string, limit int, offset int) ([]m.Course, error) {
    var courses []m.Course
    if q == "" {
        return courses, nil
    }
    
    searchPattern := "%" + q + "%"
    isNumeric := regexp.MustCompile(`^\d+$`).MatchString(q)
    
    if isNumeric {
        err := c.db.
            Preload("Tags").
            Where("name ILIKE ? OR code ILIKE ? OR description ILIKE ?", 
                searchPattern, searchPattern, searchPattern).
            Order("score DESC").
            Limit(limit).
            Offset(offset).
            Find(&courses).Error
        return courses, err
    }

    err := c.db.
        Preload("Tags").
        Where("fts @@ plainto_tsquery('english', ?) OR name ILIKE ? OR code ILIKE ?", 
            q, searchPattern, searchPattern).
        Order(gorm.Expr("ts_rank(fts, plainto_tsquery('english', ?)) DESC, score DESC", q)).
        Limit(limit).
        Offset(offset).
        Find(&courses).Error
    
    return courses, err
}

func (r *courseRepo) CountSearch(q string) (int64, error) {
    if q == "" {
        return 0, nil
    }
    
    var count int64
    searchPattern := "%" + q + "%"
    isNumeric := regexp.MustCompile(`^\d+$`).MatchString(q)
    
    if isNumeric {
        err := r.db.Model(&m.Course{}).
            Where("name ILIKE ? OR code ILIKE ? OR description ILIKE ?", 
                searchPattern, searchPattern, searchPattern).
            Count(&count).Error
        return count, err
    }

    err := r.db.Model(&m.Course{}).
        Where("fts @@ plainto_tsquery('english', ?) OR name ILIKE ? OR code ILIKE ?", 
            q, searchPattern, searchPattern).
        Count(&count).Error
    
    return count, err
}