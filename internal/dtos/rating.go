package dtos

type RatingResponse struct {
	Happiness float64 `json:"happiness"`
	Easiness  float64 `json:"easiness"`
	Quality   float64 `json:"quality"`
}

type ReviewCountResponse struct {
	Happiness int `json:"happiness"`
	Easiness  int `json:"easiness"`
	Quality   int `json:"quality"`
}