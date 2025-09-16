package dtos

type RatingResponse struct {
	Happiness float64 `json:"happiness"`
	Easiness  float64 `json:"easiness"`
	Quality   float64 `json:"quality"`
}