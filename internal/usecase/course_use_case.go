package usecase

import (
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

	return &res, nil
}