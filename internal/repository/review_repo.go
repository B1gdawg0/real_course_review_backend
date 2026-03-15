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
    err := r.db.Model(&m.Review{}).Where("id = ? and rec_status = ?", id, true).Count(&count).Error
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
    decayExpr := `reviews.score / (1.0 + (EXTRACT(EPOCH FROM (NOW() - reviews.created_at)) / 86400.0) / 180.0) AS decayed_score`

    query := r.db.
        Preload("Tags").
        Preload("User")

    if userID != "" {
        query = query.Select(`
            reviews.*,
            COALESCE(SUM(CASE WHEN votes.vote = 1 THEN 1 ELSE 0 END), 0) AS up_count,
            COALESCE(SUM(CASE WHEN votes.vote = -1 THEN 1 ELSE 0 END), 0) AS down_count,
            MAX(CASE WHEN votes.user_id = ? THEN votes.vote ELSE NULL END) AS user_vote,
        `+decayExpr, userID)
    } else {
        query = query.Select(`
            reviews.*,
            COALESCE(SUM(CASE WHEN votes.vote = 1 THEN 1 ELSE 0 END), 0) AS up_count,
            COALESCE(SUM(CASE WHEN votes.vote = -1 THEN 1 ELSE 0 END), 0) AS down_count,
        `+decayExpr)
    }

    err := query.
        Joins("LEFT JOIN votes ON votes.review_id = reviews.id").
        Where("reviews."+field+" = ? and reviews.rec_status = ?", id, true).
        Group("reviews.id").
        Order("decayed_score DESC"). // we applied wilson here
        Limit(limit).
        Offset(offset).
        Find(&reviews).Error

    return reviews, err
}

func (r *reviewRepo) CountReviewsByID(id string, section string) (int64, error) {
	var count int64
	err := r.db.Model(&m.Review{}).Where(section+"_id = ? and rec_status = ?", id, true).Count(&count).Error
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

func (r *reviewRepo) GetHotReviewsByCourseID(userid string, id string, limit, offset int) ([]m.Review, error) {
		if limit <= 0 {
			limit = 1
		}

		var reviews []m.Review

		query := r.db.
			Preload("Tags").
			Preload("User")

		if userid != "" {
			query = query.Select(`
				reviews.*,
				COALESCE(SUM(CASE WHEN votes.vote = 1 THEN 1 ELSE 0 END), 0) AS up_count,
				COALESCE(SUM(CASE WHEN votes.vote = -1 THEN 1 ELSE 0 END), 0) AS down_count,
				MAX(CASE WHEN votes.user_id = ? THEN votes.vote ELSE NULL END) AS user_vote,
				COALESCE(SUM(CASE WHEN votes.vote <> 0 THEN 1 ELSE 0 END), 0) AS interactions
			`, userid)
		} else {
			query = query.Select(`
				reviews.*,
				COALESCE(SUM(CASE WHEN votes.vote = 1 THEN 1 ELSE 0 END), 0) AS up_count,
				COALESCE(SUM(CASE WHEN votes.vote = -1 THEN 1 ELSE 0 END), 0) AS down_count,
				COALESCE(SUM(CASE WHEN votes.vote <> 0 THEN 1 ELSE 0 END), 0) AS interactions
			`)
		}

		err := query.
			Joins("LEFT JOIN votes ON votes.review_id = reviews.id").
			Where("reviews.course_id = ? and reviews.rec_status = ?", id, true).
			Group("reviews.id").
			Order("interactions DESC").
			Limit(limit).
			Offset(offset).
			Find(&reviews).Error

		return reviews, err
}

func (r *reviewRepo) UpdateScoreForReview(id string, score float64) error {
	return r.db.Model(&m.Review{}).Where("id = ?", id).Update("score", score).Error 
}

func (r *reviewRepo) UpdateReviewRecStatus(id string, rec_status bool) error {
    return r.db.Model(&m.Review{}).Where("id = ?", id).Update("rec_status", rec_status).Error
}