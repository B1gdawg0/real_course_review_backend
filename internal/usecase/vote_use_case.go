package usecase

import (
	"errors"
	repo "github.com/B1gdawg0/real_course_review_backend/internal/repository"
)

type voteUseCase struct {
	repo repo.VoteRepository
	auth repo.AuthRepository
	review repo.ReviewRepository
}

func NewVoteUseCase(repo repo.VoteRepository, auth repo.AuthRepository, review repo.ReviewRepository) VoteUseCase {
	return &voteUseCase{repo: repo, auth: auth, review: review}
}

func (v *voteUseCase) AddVoteToReview(userid string, reviewid string, vote int) error {
    userExists, err := v.auth.VerifyUserById(userid)
    if err != nil {
        return err
    }
    if !userExists {
        return errors.New("user not found")
    }
    reviewExists, err := v.review.VerifyReviewById(reviewid)
    if err != nil {
        return err
    }
    if !reviewExists {
        return errors.New("review not found")
    }
    
    existingVote, err := v.repo.GetVoteByUserAndReviewId(userid, reviewid)
    if err != nil {
        return err
    }
    
    if existingVote != nil {
        if existingVote.Vote == vote {
            return errors.New("can't vote the same vote")
        }
        return v.repo.UpdateVote(userid, reviewid, vote)
    }
    return v.repo.AddVoteToReview(userid, reviewid, vote)
}