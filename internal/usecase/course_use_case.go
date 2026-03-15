package usecase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"net/http"

	"github.com/B1gdawg0/real_course_review_backend/internal/dtos"
	"github.com/B1gdawg0/real_course_review_backend/internal/model"
	repo "github.com/B1gdawg0/real_course_review_backend/internal/repository"
	"github.com/B1gdawg0/real_course_review_backend/internal/utils"
	"github.com/jinzhu/copier"
)

type courseUseCase struct {
	repo       repo.CourseRepository
	reviewRepo repo.ReviewRepository
	n8nBaseURL string
}

func NewCourseUseCase(repo repo.CourseRepository, reviewRepo repo.ReviewRepository, n8nBaseURL string) CourseUseCase {
	return &courseUseCase{repo: repo, reviewRepo: reviewRepo, n8nBaseURL: n8nBaseURL}
}

func (c *courseUseCase) GetAll(page, size int, filter model.CourseFilter) ([]dtos.CourseShortResponse, int, int, int64, int, error) {
	page, size, offset := c.validatePagination(page, size)

	entities, total, err := c.repo.GetAll(filter, size, offset)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	res := make([]dtos.CourseShortResponse, len(entities))
	for i, entity := range entities {
		err := copier.Copy(&res[i], &entity)
		if err != nil {
			return nil, 0, 0, 0, 0, err
		}

		res[i].Tags = make([]dtos.TagResponse, len(entity.Tags))
		for j, tag := range entity.Tags {
			res[i].Tags[j] = dtos.TagResponse{
				ID:   tag.ID,
				Name: tag.Name,
			}
		}

		res[i].Rate, res[i].AvgRate = utils.ParseRate(entity.Rate)
		res[i].ReviewCount = utils.ParseReviewCount(entity.ReviewCount)
	}

	totalPages := int(math.Ceil(float64(total) / float64(size)))

	return res, page, size, total, totalPages, nil
}

func (c *courseUseCase) GetCourseById(id string) (*dtos.CourseFullResponse, error) {
	entity, err := c.repo.GetCourseById(id)
	if err != nil {
		return nil, err
	}

	var res dtos.CourseFullResponse
	err = copier.Copy(&res, entity)
	if err != nil {
		return nil, err
	}

	res.Rate, res.AvgRate = utils.ParseRate(entity.Rate)
	res.ReviewCount = utils.ParseReviewCount(entity.ReviewCount)

	return &res, nil
}

func (c *courseUseCase) CompareCoursesById(userId, first, second string) ([]dtos.CourseCompareResponse, error) {
	course1, err := c.GetCourseById(first)
	if err != nil {
		return nil, err
	}
	course2, err := c.GetCourseById(second)
	if err != nil {
		return nil, err
	}

	var res1 dtos.CourseCompareResponse
	var res2 dtos.CourseCompareResponse

	res1.AvgRate = course1.AvgRate
	res2.AvgRate = course2.AvgRate

	res1.TotalReviewCount = utils.MaxOfThreeInt(course1.ReviewCount.Easiness, course1.ReviewCount.Happiness, course1.ReviewCount.Quality)
	res2.TotalReviewCount = utils.MaxOfThreeInt(course2.ReviewCount.Easiness, course2.ReviewCount.Happiness, course2.ReviewCount.Quality)

	res1.Rate = course1.Rate
	res2.Rate = course2.Rate

	if len(course1.Tags) > 0 {
		res1.Tag = course1.Tags
	}
	if len(course2.Tags) > 0 {
		res2.Tag = course2.Tags
	}

	res1.RawContent = dtos.RawCourseResponse{
		ID:          course1.ID,
		Name:        course1.Name,
		Description: course1.Description,
		Semester:    course1.Semester,
		Code:        course1.Code,
		Credit:      course1.Credit,
		CourseType:  course1.CourseType,
		Professors:  course1.Professors,
	}

	res2.RawContent = dtos.RawCourseResponse{
		ID:          course2.ID,
		Name:        course2.Name,
		Description: course2.Description,
		Semester:    course2.Semester,
		Code:        course2.Code,
		Credit:      course2.Credit,
		CourseType:  course2.CourseType,
		Professors:  course2.Professors,
	}

	res1.HotPicks = []dtos.ReviewShortResponse{}
	hotReviews1, err := c.reviewRepo.GetReviewsByCourseId(userId, first, 1, 0)
	if err != nil {
		return nil, err
	}
	for _, review := range hotReviews1 {
		shortReview := dtos.ReviewShortResponse{
			ID:          review.ID,
			Description: review.Description,
			AvgRate:     review.AvgRate,
			IsAnonymous: review.IsAnonymous,
			Tags:        c.mapTagsToDTO(review.Tags),
			Grade:       review.Grade,
			Year:        review.Year,
			Sec:         review.Sec,
			CreatedAt:   review.CreatedAt,
			UpdatedAt:   review.UpdatedAt,
		}

		if !review.IsAnonymous && review.User.ID != "" {
			shortReview.User = &dtos.UserShortResponse{
				ID:   review.User.ID,
				Name: review.User.Name,
			}
		}
		shortReview.Rate, _ = utils.ParseRate(review.Rate)
		shortReview.Votes = dtos.VoteShortReponse{
			UpVote:   review.UpCount,
			DownVote: review.DownCount,
			HasUserVoted: func() int {
				if review.UserVote == nil {
					return 0
				}
				return *review.UserVote
			}(),
		}
		res1.HotPicks = append(res1.HotPicks, shortReview)
	}

	res2.HotPicks = []dtos.ReviewShortResponse{}
	hotReviews2, err := c.reviewRepo.GetHotReviewsByCourseID(userId, second, 1, 0)
	if err != nil {
		return nil, err
	}
	for _, review := range hotReviews2 {
		shortReview := dtos.ReviewShortResponse{
			ID:          review.ID,
			Description: review.Description,
			AvgRate:     review.AvgRate,
			IsAnonymous: review.IsAnonymous,
			Tags:        c.mapTagsToDTO(review.Tags),
			Grade:       review.Grade,
			Year:        review.Year,
			Sec:         review.Sec,
			CreatedAt:   review.CreatedAt,
			UpdatedAt:   review.UpdatedAt,
		}

		if !review.IsAnonymous && review.User.ID != "" {
			shortReview.User = &dtos.UserShortResponse{
				ID:   review.User.ID,
				Name: review.User.Name,
			}
		}
		shortReview.Rate, _ = utils.ParseRate(review.Rate)
		shortReview.Votes = dtos.VoteShortReponse{
			UpVote:   review.UpCount,
			DownVote: review.DownCount,
			HasUserVoted: func() int {
				if review.UserVote == nil {
					return 0
				}
				return *review.UserVote
			}(),
		}
		res2.HotPicks = append(res2.HotPicks, shortReview)
	}

	return []dtos.CourseCompareResponse{res1, res2}, nil
}

func (c *courseUseCase) Search(
	keyword string,
	page, size int,
	filter model.CourseFilter,
) ([]dtos.CourseShortResponse, int, int, int64, int, error) {

	page, size, offset := c.validatePagination(page, size)

	entities, err := c.repo.Search(keyword, size, offset, filter)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	total, err := c.repo.CountSearch(keyword, filter)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	res := make([]dtos.CourseShortResponse, len(entities))
	for i, entity := range entities {
		err := copier.Copy(&res[i], &entity)
		if err != nil {
			return nil, 0, 0, 0, 0, err
		}

		res[i].Tags = make([]dtos.TagResponse, len(entity.Tags))
		for j, tag := range entity.Tags {
			res[i].Tags[j] = dtos.TagResponse{
				ID:   tag.ID,
				Name: tag.Name,
			}
		}

		res[i].Rate, res[i].AvgRate = utils.ParseRate(entity.Rate)
		res[i].ReviewCount = utils.ParseReviewCount(entity.ReviewCount)
	}

	totalPages := int(math.Ceil(float64(total) / float64(size)))

	return res, page, size, total, totalPages, nil
}

func (c *courseUseCase) validatePagination(page, size int) (int, int, int) {
	if page < 1 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}
	offset := (page - 1) * size
	return page, size, offset
}

func (c *courseUseCase) mapTagsToDTO(tags []model.Tag) []dtos.TagResponse {
	tagDtos := make([]dtos.TagResponse, len(tags))
	for i, tag := range tags {
		tagDtos[i] = dtos.TagResponse{
			ID:   tag.ID,
			Name: tag.Name,
			// Don't include CourseID in response as it's not relevant for tags
		}
	}
	return tagDtos
}

func (c *courseUseCase) GetAISummaryWithN8N(firstCourseId, secondCourseId string) (*dtos.CourseAISummaryResponse, error) {
    payload := map[string]string{
        "class-id-1": firstCourseId,
        "class-id-2": secondCourseId,
    }

    jsonPayload, err := json.Marshal(payload)
    if err != nil {
        return nil, err
    }

    res, err := http.Post(c.n8nBaseURL+"/compare", "application/json", bytes.NewBuffer(jsonPayload))
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("failed to call n8n webhook: %s", res.Status)
    }

    // FIX: Decode into a single object, not a slice
    var response dtos.CourseAISummaryResponse
    if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
        return nil, fmt.Errorf("decode error: %w", err)
    }

    return &response, nil
}