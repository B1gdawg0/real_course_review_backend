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
    user, err := v.auth.VerifyUserById(userid)
    if err != nil {
        return err
    }
    
    review, err := v.review.VerifyReviewById(reviewid)
    if err != nil {
        return err
    }
    
    if !user {
        return errors.New("user not found")
    }
    
    if !review {
        return errors.New("review not found")
    }
    
    existingVote, err := v.repo.GetVoteByUserAndReviewId(userid, reviewid)
    if err != nil {
        return err
    }
    
    if existingVote != nil {
        return errors.New("user has already voted on this review")
    }
    
    if err := v.repo.AddVoteToReview(userid, reviewid, vote); err != nil {
        return err
    }
    
    return nil
}