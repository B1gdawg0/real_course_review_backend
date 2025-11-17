package usecase

// import "github.com/B1gdawg0/real_course_review_backend/internal/dtos"

type VoteUseCase interface{
	AddVoteToReview(userid string, reviewid string, vote int) (error)
}