package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/B1gdawg0/real_course_review_backend/internal/dtos"
	m "github.com/B1gdawg0/real_course_review_backend/internal/model"
)

func ParseRate(rateStr string) (dtos.RatingResponse, float64) {
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

	return rating, (rating.Easiness+rating.Happiness+rating.Quality)/3
}

func ParseReviewCount(reviewStr string) dtos.ReviewCountResponse{
	parts := strings.Split(reviewStr, ",")
	reviewCount := dtos.ReviewCountResponse{}

	if len(parts) > int(m.Quality) {
		if val, err := strconv.ParseFloat(parts[m.Happiness], 64); err == nil {
			reviewCount.Happiness = int(val)
		}
		if val, err := strconv.ParseFloat(parts[m.Easiness], 64); err == nil {
			reviewCount.Easiness = int(val)
		}
		if val, err := strconv.ParseFloat(parts[m.Quality], 64); err == nil {
			reviewCount.Quality = int(val)
		}
	}

	return reviewCount
}


func RateToString(r dtos.RatingResponse) string {
	return fmt.Sprintf("%.2f,%.2f,%.2f",
		r.Happiness,
		r.Easiness,
		r.Quality,
	)
}

func ReviewCountToString(r dtos.ReviewCountResponse) string {
	return fmt.Sprintf("%d,%d,%d",
		r.Happiness,
		r.Easiness,
		r.Quality,
	)
}

func ParseCommaSeparated(s string) []string {
    if s == "" {
        return nil
    }
    parts := strings.Split(s, ",")
    result := make([]string, 0, len(parts))
    for _, p := range parts {
        if trimmed := strings.TrimSpace(p); trimmed != "" {
            result = append(result, trimmed)
        }
    }
    return result
}