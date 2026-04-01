package dtos

type CourseShortResponse struct {
	ID          string              `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description,omitempty"`
	Semester    string              `json:"semester"`
	Code        string              `json:"code"`
	Credit      int                 `json:"credit"`
	CourseType  string              `json:"course_type"`
	ReviewCount ReviewCountResponse `json:"review_count"`
	Rate        RatingResponse      `json:"rating"`
	Tags        []TagResponse       `json:"tags"`
	AvgRate     float64             `json:"avg_rating_score"`
	// Score       float64             `json:"score"`
}

type CourseFullResponse struct {
	CourseShortResponse
	Professors []ProfessorShortResponse `json:"professors"`
}

type CourseCompareResponse struct {
	// TODO: Explore oppotinity for Hoter field by compare user click count
	AvgRate          float64               `json:"avg_rating_score"`
	TotalReviewCount int                   `json:"total_review_count"`
	Rate             RatingResponse        `json:"rating"`
	Tag              []TagResponse         `json:"tags"`
	HotPicks         []ReviewShortResponse `json:"hot_picks"`
	RawContent       RawCourseResponse     `json:"raw_content"`
}

type RawCourseResponse struct {
	ID          string                   `json:"id"`
	Name        string                   `json:"name"`
	Description string                   `json:"description,omitempty"`
	Semester    string                   `json:"semester"`
	Code        string                   `json:"code"`
	Credit      int                      `json:"credit"`
	CourseType  string                   `json:"course_type"`
	Professors  []ProfessorShortResponse `json:"professors"`
}

type CourseAISummaryResponse struct {
	Output string `json:"output"`
}

type UpdateCourseRecStatusRequest struct {
	ID       string `json:"id" validate:"required"`
	RecStatus bool   `json:"rec_status" validate:"required"`
}

type BulkCreateCoursesResponse struct {
	SuccessCount int      `json:"success_count"`
	FailedCount  int      `json:"failed_count"`
	FailedRows   []int    `json:"failed_rows"`
}

type UpdateCourseRequest struct {
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	Semester    *string  `json:"semester,omitempty"`
	Code        *string  `json:"code,omitempty"`
	Credit      *int     `json:"credit,omitempty"`
	CourseType  *string  `json:"course_type,omitempty"`
}