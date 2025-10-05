package usecase

import (
	"github.com/B1gdawg0/real_course_review_backend/internal/dtos"
	m "github.com/B1gdawg0/real_course_review_backend/internal/model"
	repo "github.com/B1gdawg0/real_course_review_backend/internal/repository"
	"github.com/B1gdawg0/real_course_review_backend/internal/utils"
	"github.com/jinzhu/copier"
)

type professorUseCase struct {
	repo repo.ProfessorRepository
}

func NewProfessorUseCase(repo repo.ProfessorRepository) ProfessorUseCase {
	return &professorUseCase{repo: repo}
}

func (p *professorUseCase) GetAll() ([]dtos.ProfessorShortResponse, error) {
	entities, err := p.repo.GetAll()
	if err != nil {
		return nil, err
	}

	res := make([]dtos.ProfessorShortResponse, len(entities))
	for i, entity := range entities {
		err = copier.Copy(&res[i], &entity)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

func (p *professorUseCase) GetProfessorById(id string) (*dtos.ProfessorFullResponse, error) {
    entity, err := p.repo.GetProfessorById(id)
    if err != nil {
        return nil, err
    }
    
    var res dtos.ProfessorFullResponse
    res.ID = entity.ID
    res.Name = entity.Name
    res.Email = entity.Email
    res.Title = entity.Title
    res.ImageURL = entity.ImageURL
    res.Phone = entity.Phone
    res.Description = entity.Description
    res.UniRoomAddress = entity.UniRoomAddress
    // res.ReviewCount = 0
    
    res.Classes = make([]dtos.CourseShortResponse, 0, len(entity.Classes))
    for _, course := range entity.Classes {
        rate, avg := utils.ParseRate(course.Rate)
        res.Classes = append(res.Classes, dtos.CourseShortResponse{
            ID:          course.ID,
            Name:        course.Name,
            Description: course.Description,
            Semester:    course.Semester,
            Code:        course.Code,
            Credit:      course.Credit,
            ReviewCount: utils.ParseReviewCount(course.ReviewCount),
            Rate:        rate,
            AvgRate:     avg,
            Score:       course.Score,
            Tags:        mapTagsToDTO(course.Tags),
        })
    }
    
    return &res, nil
}

func mapTagsToDTO(tags []m.Tag) []dtos.TagResponse {
	res := make([]dtos.TagResponse, len(tags))
	for i, t := range tags {
		res[i] = dtos.TagResponse{
			ID:   t.ID,
			Name: t.Name,
		}
	}
	return res
}