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

func (r *reviewRepo) VerifyReviewById (id string)(bool, error){
	var count int64
    err := r.db.Model(&m.Review{}).Where("id = ?", id).Count(&count).Error
    if err != nil {
        return false, err
    }
    return count > 0, nil
}

func (r *reviewRepo) GetReviewsByCourseId(userid string, id string, limit, offset int) ([]m.Review, error) {
    return r.getReviewsWithVotes("course_id", id, userid, limit, offset)
}

// func (r *reviewRepo) GetReviewsByUserId(id string, limit, offset int) ([]m.Review, error) {
//     return r.getReviewsWithVotes("user_id", id, "", limit, offset)
// }

func (r *reviewRepo) getReviewsWithVotes(field, id, userID string, limit, offset int) ([]m.Review, error) {
    var reviews []m.Review

    query := r.db.
        Preload("Tags").
        Preload("User")

    if userID != "" {
        query = query.Select(`
            reviews.*,
            COALESCE(SUM(CASE WHEN votes.vote = 1 THEN 1 ELSE 0 END), 0) AS up_count,
            COALESCE(SUM(CASE WHEN votes.vote = -1 THEN 1 ELSE 0 END), 0) AS down_count,
            MAX(CASE WHEN votes.user_id = ? THEN votes.vote ELSE NULL END) AS user_vote
        `, userID)
    } else {
        query = query.Select(`
            reviews.*,
            COALESCE(SUM(CASE WHEN votes.vote = 1 THEN 1 ELSE 0 END), 0) AS up_count,
            COALESCE(SUM(CASE WHEN votes.vote = -1 THEN 1 ELSE 0 END), 0) AS down_count
        `)
    }

    err := query.
        Joins("LEFT JOIN votes ON votes.review_id = reviews.id").
        Where("reviews."+field+" = ?", id).
        Group("reviews.id").
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

func (r *reviewRepo) CreateReview(review *m.Review, course *m.Course) error {
    return r.db.Transaction(func(tx *gorm.DB) error {
        if err := tx.Create(review).Error; err != nil {
            return err
        }

        if len(review.Tags) > 0 {
            if err := tx.Model(review).Association("Tags").Replace(review.Tags); err != nil {
                return err
            }
        }

		if len(review.Tags) > 0 {
            if err := tx.Model(course).Association("Tags").Append(review.Tags); err != nil {
                return err
            }
        }

        if err := tx.Save(course).Error; err != nil {
            return err
        }

        return nil
    })
}
