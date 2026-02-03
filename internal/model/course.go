package model

import "time"

type Course struct {
	ID          string  `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name        string  `gorm:"column:name;not null"`
	Description string  `gorm:"column:description;type:text"`
	Semester    string  `gorm:"column:semester;not null"`
	Code        string  `gorm:"column:code;uniqueIndex;not null"`
	Credit      int     `gorm:"column:credit;not null"`
	CourseType  string  `gorm:"column:course_type;not null;default:'ELECTIVE_SPECIALIZED'"`
	ReviewCount string  `gorm:"column:review_count;default:0"`
	Rate        string  `gorm:"column:rate;default:0"`
	Score       float64 `gorm:"column:score;default:0"`

	FTS string `gorm:"column:fts;type:tsvector;index:,type:gin"`

	Professors []Professor `gorm:"many2many:course_professors;"`
	Tags       []Tag       `gorm:"many2many:course_tags;"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
