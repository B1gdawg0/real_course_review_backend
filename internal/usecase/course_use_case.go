package usecase

import (
	"math"
	"sort"

	"github.com/B1gdawg0/real_course_review_backend/internal/dtos"
	repo "github.com/B1gdawg0/real_course_review_backend/internal/repository"
	"github.com/B1gdawg0/real_course_review_backend/internal/utils"
	"github.com/jinzhu/copier"
)

type courseUseCase struct {
	repo repo.CourseRepository
}

func NewCourseUseCase(repo repo.CourseRepository) CourseUseCase {
	return &courseUseCase{repo: repo}
}

func (c *courseUseCase) GetAll() ([]dtos.CourseShortResponse, error) {
	entities, err := c.repo.GetAll()
	if err != nil {
		return nil, err
	}

	res := make([]dtos.CourseShortResponse, len(entities))
	for i, entity := range entities {
		err := copier.Copy(&res[i], &entity)
		if err != nil {
			return nil, err
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

	sort.Slice(res, func(i, j int) bool {
		return res[i].Score > res[j].Score
	})

	return res, nil
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

func (c *courseUseCase) Search(
    keyword string,
    page, size int,
) ([]dtos.CourseShortResponse, int, int, int64, int, error) {
    
    page, size, offset := c.validatePagination(page, size)
    
    entities, err := c.repo.Search(keyword, size, offset)
    if err != nil {
        return nil, 0, 0, 0, 0, err
    }
    
    total, err := c.repo.CountSearch(keyword)
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