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

func applyFilters(db *gorm.DB, f m.CourseFilter) *gorm.DB {
    db = db.Where("rec_status = ?", true)

    if f.Semester != "" {
        db = db.Where("semester = ?", f.Semester)
    }
    if f.CourseType != "" {
        db = db.Where("course_type = ?", f.CourseType)
    }
    if len(f.TagIDs) > 0 {
        db = db.
            Joins("JOIN course_tags ON course_tags.course_id = courses.id").
            Where("course_tags.tag_id IN ?", f.TagIDs).
            Group("courses.id").
            Having("COUNT(DISTINCT course_tags.tag_id) = ?", len(f.TagIDs))
    }
    return db
}

func (c *courseRepo) GetAll(filter m.CourseFilter, limit, offset int) ([]m.Course, int64, error) {
	var courses []m.Course
	var total int64

	base := applyFilters(c.db.Model(&m.Course{}), filter)

	if err := base.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := applyFilters(c.db.Preload("Tags"), filter).
		Order(`(
            COALESCE(NULLIF(split_part(rate, ',', 1), '')::float, 0) +
            COALESCE(NULLIF(split_part(rate, ',', 2), '')::float, 0) +
            COALESCE(NULLIF(split_part(rate, ',', 3), '')::float, 0)
        ) / 3 DESC`).
		Limit(limit).
		Offset(offset).
		Find(&courses).Error

	return courses, total, err
}

func (c *courseRepo) GetCourseById(id string) (*m.Course, error) {
	var course m.Course
	err := c.db.Preload("Tags").Preload("Professors").Where("id = ? AND rec_status = ?", id, true).First(&course).Error
	return &course, err
}

func (c *courseRepo) Search(q string, limit, offset int, filter m.CourseFilter) ([]m.Course, error) {
	var courses []m.Course
	if q == "" {
		return courses, nil
	}
	searchPattern := "%" + q + "%"
	isNumeric := regexp.MustCompile(`^\d+$`).MatchString(q)

	base := applyFilters(c.db.Preload("Tags"), filter)

	if isNumeric {
		err := base.
			Where("name ILIKE ? OR code ILIKE ? OR description ILIKE ?",
				searchPattern, searchPattern, searchPattern).
			Order("score DESC").
			Limit(limit).Offset(offset).
			Find(&courses).Error
		return courses, err
	}

	err := base.
		Where("fts @@ plainto_tsquery('english', ?) OR name ILIKE ? OR code ILIKE ?",
			q, searchPattern, searchPattern).
		Order(gorm.Expr("ts_rank(fts, plainto_tsquery('english', ?)) DESC, score DESC", q)).
		Limit(limit).Offset(offset).
		Find(&courses).Error
	return courses, err
}

func (r *courseRepo) CountSearch(q string, filter m.CourseFilter) (int64, error) {
	if q == "" {
		return 0, nil
	}
	var count int64
	searchPattern := "%" + q + "%"
	isNumeric := regexp.MustCompile(`^\d+$`).MatchString(q)

	base := applyFilters(r.db.Model(&m.Course{}), filter)

	if isNumeric {
		err := base.
			Where("name ILIKE ? OR code ILIKE ? OR description ILIKE ?",
				searchPattern, searchPattern, searchPattern).
			Count(&count).Error
		return count, err
	}

	err := base.
		Where("fts @@ plainto_tsquery('english', ?) OR name ILIKE ? OR code ILIKE ?",
			q, searchPattern, searchPattern).
		Count(&count).Error
	return count, err
}

func (c *courseRepo) UpdateCourseRecStatus(id string, recStatus bool) error {
	return c.db.
		Model(&m.Course{}).
		Where("id = ?", id).
		Update("rec_status", recStatus).
		Error
}

func (c *courseRepo) UpdateCourse(course *m.Course) error {
	return c.db.Model(&m.Course{}).
		Where("id = ?", course.ID).
		Select("rate", "review_count").
		Updates(map[string]interface{}{
			"rate":         course.Rate,
			"review_count": course.ReviewCount,
		}).Error
}

func (c *courseRepo) BulkCreateCourses(courses []m.Course) error {
	tx := c.db.Begin()

	if err := tx.Create(&courses).Error; err != nil {
		tx.Rollback()
		return err
	}

	for i := range courses {
		if len(courses[i].Professors) > 0 {
			if err := tx.Model(&courses[i]).
				Association("Professors").
				Append(courses[i].Professors); err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit().Error
}
