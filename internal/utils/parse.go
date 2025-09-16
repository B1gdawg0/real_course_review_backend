package utils

import (
	"strconv"
	"strings"

	"github.com/B1gdawg0/real_course_review_backend/internal/dtos"
	m "github.com/B1gdawg0/real_course_review_backend/internal/model"
)

func ParseRate(rateStr string) dtos.RatingResponse {
	parts := strings.Split(rateStr, ",")
	rating := dtos.RatingResponse{}

	if len(parts) > int(m.Quality) {
		if val, err := strconv.ParseFloat(parts[m.Happiness], 64); err == nil {
			rating.Happiness = val
		}
		if val, err := strconv.ParseFloat(parts[m.Easiness], 64); err == nil {
			rating.Easiness = val
		}
		if val, err := strconv.ParseFloat(parts[m.Quality], 64); err == nil {
			rating.Quality = val
		}
	}

	return rating
}