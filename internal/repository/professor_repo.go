package repository

import (
	"fmt"

	m "github.com/B1gdawg0/real_course_review_backend/internal/model"
	"gorm.io/gorm"
)

type profRepo struct {
	db *gorm.DB
}

func NewProfessorRepository(db *gorm.DB) ProfessorRepository {
	return &profRepo{
		db: db,
	}
}

func (p *profRepo) GetAll() ([]m.Professor, error) {
	var profs []m.Professor
	err := p.db.Find(&profs).Error
	return profs, err
}

func (p *profRepo) GetProfessorById(id string) (*m.Professor, error) {
    var prof m.Professor
    err := p.db.
        Preload("Classes.Tags").
        First(&prof, "id = ?", id).Error
    if err != nil {
        return nil, err
    }
    
    // Debug: Check if duplicates exist here
    fmt.Printf("Number of classes loaded: %d\n", len(prof.Classes))
    for i, class := range prof.Classes {
        fmt.Printf("Class %d: %s (ID: %s)\n", i, class.Name, class.ID)
    }
    
    return &prof, nil
}