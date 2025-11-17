package dtos

type VoteReviewRequest struct{
	ReviewID string `json:"review"`
	Vote 	 int 	`json:"vote"`
}

type VoteShortReponse struct{
	UpVote int `json:"upvote"`
	DownVote int `json:"downvote"`
	HasUserVoted int `json:"hasUserVoted"`
}