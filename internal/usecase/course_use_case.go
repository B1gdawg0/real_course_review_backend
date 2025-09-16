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

	dtos := make([]dtos.CourseShortResponse, len(entities))
	for i, entity := range entities {
		err := copier.Copy(&dtos[i], &entity)
		if err != nil {
			return nil, err
		}
		dtos[i].Rate = utils.ParseRate(entity.Rate)
		dtos[i].AvgRate = 4.0
	}

	sort.Slice(dtos, func(i, j int) bool {
		return dtos[i].Score > dtos[j].Score
	})

	return dtos, nil
}

func (c *courseUseCase) GetCourseById(id string) (*dtos.CourseFullResponse, error) {
	entity, err := c.repo.GetCourseById(id)
	if err != nil {
		return nil, err
	}

	var dtos dtos.CourseFullResponse
	err = copier.Copy(&dtos, entity)
	if err != nil {
		return nil, err
	}

	dtos.Rate = utils.ParseRate(entity.Rate)
	dtos.AvgRate = 4.0

	return &dtos, nil
}