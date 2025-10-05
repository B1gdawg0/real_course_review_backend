package usecase

import (
	"github.com/B1gdawg0/real_course_review_backend/internal/dtos"
	repo "github.com/B1gdawg0/real_course_review_backend/internal/repository"
	"github.com/jinzhu/copier"
)

type tagUseCase struct {
	repo repo.TagRepository
}

func NewTagUseCase(repo repo.TagRepository) TagUseCase {
	return &tagUseCase{repo: repo}
}

func (t *tagUseCase) GetAll() ([]dtos.TagResponse, error) {
	entities, err := t.repo.GetAll()
	if err != nil {
		return nil, err
	}

	res := make([]dtos.TagResponse, len(entities))
	for i, entity := range entities {
		err = copier.Copy(&res[i], &entity)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

func (t *tagUseCase) GetTagById(id string) (*dtos.TagResponse, error) {
	entity, err := t.repo.GetTagById(id)
	if err != nil {
		return nil, err
	}

	var res dtos.TagResponse
	err = copier.Copy(&res, &entity)
	if err != nil {
		return nil, err
	}
	
	return  &res, nil
}