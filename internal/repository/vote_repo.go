package repository

import (
	"errors"

	m "github.com/B1gdawg0/real_course_review_backend/internal/model"
	"gorm.io/gorm"
)

type VoteRepo struct {
	db *gorm.DB
}

func NewVoteRepository(db *gorm.DB) VoteRepository {
	return &VoteRepo{
		db: db,
	}
}

func (v *VoteRepo) GetVoteByUserAndReviewId(userid string, reviewid string) (*m.Vote, error) {
	var vote m.Vote
    err := v.db.Where("user_id = ? AND review_id = ?", userid, reviewid).First(&vote).Error
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, nil
        }
        return nil, err
    }
    return &vote, nil
}

func (v *VoteRepo) AddVoteToReview(userid string, reviewid string, vote int) error {
	newVote := m.Vote{
		UserID:   userid,
		ReviewID: reviewid,
		Vote:     vote,
	}

	err := v.db.Create(&newVote).Error
	if err != nil {
		return err
	}

	return nil
}

func (v *VoteRepo) GetVotesByReviewId(id string) (int, int, error) {
	var result struct {
		Upvotes   int
		Downvotes int
	}

	err := v.db.Model(&m.Vote{}).
		Select("SUM(CASE WHEN vote = 1 THEN 1 ELSE 0 END) as upvotes, SUM(CASE WHEN vote = -1 THEN 1 ELSE 0 END) as downvotes").
		Where("review_id = ?", id).
		Scan(&result).Error

	if err != nil {
		return 0, 0, err
	}

	return result.Upvotes, result.Downvotes, nil
}

func (v *VoteRepo) UpdateVote(userid string, reviewid string, vote int) error {
	return v.db.Model(&m.Vote{}).
		Where("user_id = ? AND review_id = ?", userid, reviewid).
		Update("vote", vote).Error
}

func (r *VoteRepo) DeleteVote(userID, reviewID string) error {
    return r.db.
        Where("user_id = ? AND review_id = ?", userID, reviewID).
        Delete(&m.Vote{}).Error
}
