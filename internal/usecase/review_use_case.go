package usecase

import (
    "math"
    "github.com/B1gdawg0/real_course_review_backend/internal/dtos"
    m "github.com/B1gdawg0/real_course_review_backend/internal/model"
    repo "github.com/B1gdawg0/real_course_review_backend/internal/repository"
    "github.com/B1gdawg0/real_course_review_backend/internal/utils"
)

type reviewUseCase struct {
    repo repo.ReviewRepository
}

func NewReviewUseCase(repo repo.ReviewRepository) ReviewUseCase {
    return &reviewUseCase{repo: repo}
}

func (r *reviewUseCase) GetReviewsByCourseID(id string, page, size int) ([]dtos.ReviewFullResponse, int, int, int64, int, error) {
    return r.getReviews(
        id,
        page,
        size,
        r.repo.GetReviewsByCourseId,
        r.repo.CountReviewsByID,
        "course",
    )
}

func (r *reviewUseCase) GetReviewsByUserID(id string, page, size int) ([]dtos.ReviewFullResponse, int, int, int64, int, error) {
    return r.getReviews(
        id,
        page,
        size,
        r.repo.GetReviewsByUserId,
        r.repo.CountReviewsByID,
        "user",
    )
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
        Score:       review.Score,
        UpCount:     review.UpCount,
        DownCount:   review.DownCount,
        ReportCount: review.ReportCount,
        IsAnonymous: review.IsAnonymous,
        Tags:        r.mapTagsToDTO(review.Tags), // Updated to handle normalized tags
        CreatedAt:   review.CreatedAt,
        UpdatedAt:   review.UpdatedAt,
    }

    if !review.IsAnonymous && review.User.ID != "" {
        dto.User = &dtos.UserShortResponse{
            ID:   review.User.ID,
            Name: review.User.Name,
        }
    }

    // Parse rate
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
    id string,
    page, size int,
    fetchFunc func(string, int, int) ([]m.Review, error),
    countFunc func(string, string) (int64, error),
    entityType string,
) ([]dtos.ReviewFullResponse, int, int, int64, int, error) {
    page, size, offset := r.validatePagination(page, size)

    reviews, err := fetchFunc(id, size, offset)
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