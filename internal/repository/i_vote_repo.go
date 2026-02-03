package repository

import m "github.com/B1gdawg0/real_course_review_backend/internal/model"

type VoteRepository interface{
	AddVoteToReview(userid string, reviewid string, vote int) (error)
	UpdateVote(userid string, reviewid string, vote int) (error)
	DeleteVote(userID, reviewID string) error
	GetVotesByReviewId(id string) (int, int, error)
	GetVoteByUserAndReviewId(userid string, reviewid string) (*m.Vote, error)
}