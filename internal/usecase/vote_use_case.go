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
    if vote != 1 && vote != -1 {
        return errors.New("invalid vote value")
    }

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

    // 🟢 
    if existingVote == nil {
        return v.repo.AddVoteToReview(userid, reviewid, vote)
    }

    // 🔴
    if existingVote.Vote == vote {
        return v.repo.DeleteVote(userid, reviewid)
    }

    // 🟡 
    return v.repo.UpdateVote(userid, reviewid, vote)
}
