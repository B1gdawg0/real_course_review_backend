package usecase

import (
	"fmt"
	"math"

	"github.com/B1gdawg0/real_course_review_backend/internal/dtos"
	m "github.com/B1gdawg0/real_course_review_backend/internal/model"
	repo "github.com/B1gdawg0/real_course_review_backend/internal/repository"
	"github.com/B1gdawg0/real_course_review_backend/internal/utils"
)

type reviewUseCase struct {
    repo repo.ReviewRepository
	courseRepo repo.CourseRepository
}

func NewReviewUseCase(repo repo.ReviewRepository, courseRepo repo.CourseRepository) ReviewUseCase {
    return &reviewUseCase{repo: repo, courseRepo: courseRepo}
}

func (r *reviewUseCase) GetReviewsByCourseID(userid string,id string, page, size int) ([]dtos.ReviewFullResponse, int, int, int64, int, error) {
    return r.getReviews(
		userid,
        id,
        page,
        size,
        r.repo.GetReviewsByCourseId,
        r.repo.CountReviewsByID,
        "course",
    )
}

func (r *reviewUseCase) GetHotReviewsByCourseID(userid string, id string) (dtos.ReviewFullResponse, int, int, int64, int, error) {
	reviews, err := r.repo.GetHotReviewsByCourseID(userid, id, 1, 0)
	if err != nil {
		return dtos.ReviewFullResponse{}, 0, 0, 0, 0, err
	}

	if len(reviews) == 0 {
		return dtos.ReviewFullResponse{}, 0, 0, 0, 0, nil
	}

	return r.mapReviewToDTO(reviews[0]), 1, 1, 1, 1, nil
}

// func (r *reviewUseCase) GetReviewsByUserID(id string, page, size int) ([]dtos.ReviewFullResponse, int, int, int64, int, error) {
//     return r.getReviews(
//         id,
//         page,
//         size,
//         r.repo.GetReviewsByUserId,
//         r.repo.CountReviewsByID,
//         "user",
//     )
// }

func (r *reviewUseCase) CreateReview(req dtos.CreateReviewRequest, userID string) (*dtos.ReviewFullResponse, error) {
    var tags []m.Tag
    for _, t := range req.Tags {
		tags = append(tags, m.Tag{ID: t})
    }

	course, err := r.courseRepo.GetCourseById(req.CourseID)
	if err != nil{
		return nil, err
	}

	course, err = utils.RecalculateCourseReview(course, &req)
	if err != nil{
		return nil, err
	}

	review := &m.Review{
        UserID:      userID,
        CourseID:    course.ID,
        Description: req.Comment,
        Rate:        fmt.Sprintf("%f,%f,%f",req.Rate.Happiness,req.Rate.Easiness,req.Rate.Quality),
        Tags:        tags,
		Grade:       req.Grade,
		Year:        req.Year,
		Sec:         req.Sec,
        IsAnonymous: req.IsAnonymous,
        ProfessorID: req.ProfessorID,
    }

    if err := r.repo.CreateReview(review, course); err != nil {
        return nil, err
    }

    dto := r.mapReviewToDTO(*review)
    return &dto, nil
}


// ------------------ Helper Functions ------------------

func (r *reviewUseCase) validatePagination(page, size int) (int, int, int) {
    if page < 1 {
        page = 1
    }
    if size <= 0 {
        size = 10
    }
    offset := (page - 1) * size
    return page, size, offset
}

// Map Tag entities to TagResponse DTOs
func (r *reviewUseCase) mapTagsToDTO(tags []m.Tag) []dtos.TagResponse {
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

// Map a single Review entity to DTO
func (r *reviewUseCase) mapReviewToDTO(review m.Review) dtos.ReviewFullResponse {
    dto := dtos.ReviewFullResponse{
        ID:          review.ID,
        Status:      review.Status,
        Description: review.Description,
        AvgRate:     review.AvgRate,
        Score:       review.DecayedScore, // magic happened at query level
        ReportCount: review.ReportCount,
        IsAnonymous: review.IsAnonymous,
        Tags:        r.mapTagsToDTO(review.Tags), // Updated to handle normalized tags
		Grade:       review.Grade,
		Year: 		 review.Year,
		Sec:         review.Sec,
        CreatedAt:   review.CreatedAt,
        UpdatedAt:   review.UpdatedAt,
        Professor:   dtos.ProfessorNameResponse{
            ID: review.Professor.ID,
            Name: review.Professor.Name,
        },
    }

    if !review.IsAnonymous && review.User.ID != "" {
        dto.User = &dtos.UserShortResponse{
            ID:   review.User.ID,
            Name: review.User.Name,
        }
    }

    dto.Votes = dtos.VoteShortReponse{
        UpVote:   review.UpCount,
        DownVote: review.DownCount,
        HasUserVoted: func() int {
			if review.UserVote == nil {
				return 0
			}
			return *review.UserVote
		}(),
    }

    dto.Rate, dto.AvgRate = utils.ParseRate(review.Rate)
    return dto
}

func (r *reviewUseCase) mapReviewsToDTOs(reviews []m.Review) []dtos.ReviewFullResponse {
    reviewDtos := make([]dtos.ReviewFullResponse, len(reviews))
    for i, review := range reviews {
        reviewDtos[i] = r.mapReviewToDTO(review)
    }
    return reviewDtos
}

func (r *reviewUseCase) getReviews(
	userid string,
    id string,
    page, size int,
    fetchFunc func(string, string, int, int) ([]m.Review, error),
    countFunc func(string, string) (int64, error),
    entityType string,
) ([]dtos.ReviewFullResponse, int, int, int64, int, error) {
    page, size, offset := r.validatePagination(page, size)

    reviews, err := fetchFunc(userid, id, size, offset)
    if err != nil {
        return nil, 0, 0, 0, 0, err
    }

    total, err := countFunc(id, entityType)
    if err != nil {
        return nil, 0, 0, 0, 0, err
    }

    totalPages := int(math.Ceil(float64(total) / float64(size)))
    reviewDtos := r.mapReviewsToDTOs(reviews)

    return reviewDtos, page, size, total, totalPages, nil
}

func (r *reviewUseCase) UpdateReviewRecStatus(id string, rec_status bool) error {
    return r.repo.UpdateReviewRecStatus(id, rec_status)
}



























// func (r *reviewUseCase) AddTagsToReview(reviewID string, tagIDs []string) error {
//     return r.repo.AddTagsToReview(reviewID, tagIDs)
// }

// func (r *reviewUseCase) RemoveTagsFromReview(reviewID string, tagIDs []string) error {
//     return r.repo.RemoveTagsFromReview(reviewID, tagIDs)
// }

// func (r *reviewUseCase) GetReviewsByTags(tagNames []string, page, size int) ([]dtos.ReviewFullResponse, int, int, int64, int, error) {
//     page, size, offset := r.validatePagination(page, size)

//     reviews, err := r.repo.GetReviewsByTags(tagNames, size, offset)
//     if err != nil {
//         return nil, 0, 0, 0, 0, err
//     }

//     total, err := r.repo.CountReviewsByTags(tagNames)
//     if err != nil {
//         return nil, 0, 0, 0, 0, err
//     }

//     totalPages := int(math.Ceil(float64(total) / float64(size)))
//     reviewDtos := r.mapReviewsToDTOs(reviews)

//     return reviewDtos, page, size, total, totalPages, nil
// }