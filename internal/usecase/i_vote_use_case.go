package usecase

type VoteUseCase interface{
	AddVoteToReview(userid string, reviewid string, vote int) (error)
}