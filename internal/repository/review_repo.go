package repository

import (
	m "github.com/B1gdawg0/real_course_review_backend/internal/model"
	"gorm.io/gorm"
)

type reviewRepo struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &reviewRepo{
		db: db,
	}
}

func (r *reviewRepo) GetReviewsByCourseId(id string, limit, offset int) ([]m.Review, error) {
	var reviews []m.Review
	err := r.db.
		Preload("Tags").
		Preload("User").
		Preload("Votes").
		Where("course_id = ?", id).
		Limit(limit).
		Offset(offset).
		Find(&reviews).Error

	return reviews, err
}

func (r *reviewRepo) GetReviewsByUserId(id string, limit, offset int) ([]m.Review, error) {
	var reviews []m.Review
	err := r.db.
		Preload("Tags").
		Preload("User").
		Preload("Votes").
		Where("user_id = ?", id).
		Limit(limit).
		Offset(offset).
		Find(&reviews).Error

	return reviews, err
}

func (r *reviewRepo) CountReviewsByID(id string, section string) (int64, error) {
	var count int64
	err := r.db.Model(&m.Review{}).Where(section+"_id = ?", id).Count(&count).Error
	return count, err
}

func (r *reviewRepo) CreateReview(review *m.Review) error {
    return r.db.Transaction(func(tx *gorm.DB) error {
        if err := tx.Create(review).Error; err != nil {
            return err
        }

        if len(review.Tags) > 0 {
            if err := tx.Model(review).Association("Tags").Replace(review.Tags); err != nil {
                return err
            }
        }

        return nil
    })
}