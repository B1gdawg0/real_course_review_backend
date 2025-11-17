package dtos

type VoteReviewRequest struct{
	ReviewID string `json:"review"`
	UserID   string `json:"user"`
	Vote 	 int 	`json:"vote"`
}

type VoteShortReponse struct{
	UpVote int `json:"upvote"`
	DownVote int `json:"downvote"`
}