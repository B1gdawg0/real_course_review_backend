package utils

import (
	"github.com/B1gdawg0/real_course_review_backend/internal/dtos"
	m "github.com/B1gdawg0/real_course_review_backend/internal/model"
)

func RecalculateCourseReview(course *m.Course, req *dtos.CreateReviewRequest) (*m.Course, error) {
	reviewCount := ParseReviewCount(course.ReviewCount)

	reviewCount.Easiness++
	reviewCount.Happiness++
	reviewCount.Quality++

	rate, _ := ParseRate(course.Rate)

	rate.Easiness = ((rate.Easiness * float64(reviewCount.Easiness-1)) + req.Rate.Easiness) / float64(reviewCount.Easiness)
	rate.Happiness = ((rate.Happiness * float64(reviewCount.Happiness-1)) + req.Rate.Happiness) / float64(reviewCount.Happiness)
	rate.Quality = ((rate.Quality * float64(reviewCount.Quality-1)) + req.Rate.Quality) / float64(reviewCount.Quality)

	course.ReviewCount = ReviewCountToString(reviewCount)
	course.Rate = RateToString(rate)

	return course, nil
}

func ReduceRateCourseReview(course *m.Course, review *m.Review) (*m.Course, error) {
	rate, _ := ParseRate(course.Rate)
	reviewCount := ParseReviewCount(course.ReviewCount)

	dtoReview, _ := ParseRate(review.Rate)

	reviewCount.Easiness--
	reviewCount.Happiness--
	reviewCount.Quality--

	rate.Easiness = ((rate.Easiness * float64(reviewCount.Easiness+1)) - dtoReview.Easiness) / float64(reviewCount.Easiness)
	rate.Happiness = ((rate.Happiness * float64(reviewCount.Happiness+1)) - dtoReview.Happiness) / float64(reviewCount.Happiness)
	rate.Quality = ((rate.Quality * float64(reviewCount.Quality+1)) - dtoReview.Quality) / float64(reviewCount.Quality)	

	course.ReviewCount = ReviewCountToString(reviewCount)
	course.Rate = RateToString(rate)

	return course, nil
}

func MaxOfThreeInt(a, b, c int) int {
		if a >= b && a >= c {
			return a
		}
		if b >= a && b >= c {
			return b
		}
		return c
}