package dtos

type CourseShortResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Semester    string  `json:"semester"`
	Code        string  `json:"code"`
	Credit      int     `json:"credit"`
	ReviewCount ReviewCountResponse     `json:"review_count"`
	Rate        RatingResponse `json:"rating"`
	Tags 		[]TagResponse `json:"tags"`
	AvgRate     float64             `json:"avg_rating_score"`
	Score       float64 `json:"score"`
}

type CourseFullResponse struct {
	CourseShortResponse
	Professors   []ProfessorShortResponse `json:"professors"`
}
