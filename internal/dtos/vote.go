package dtos

type VoteReviewRequest struct{
	ReviewID string `json:"review_id"`
	Vote 	 int 	`json:"vote"`
}

type VoteShortReponse struct{
	UpVote int `json:"upvote"`
	DownVote int `json:"downvote"`
	HasUserVoted int `json:"hasUserVoted"`
}