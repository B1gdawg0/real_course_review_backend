package model

import "time"

type Course struct {
	ID          string  `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name        string  `gorm:"column:name;not null"`
	Description string  `gorm:"column:description;type:text"`
	Semester    string  `gorm:"column:semester;not null"`
	Code        string  `gorm:"column:code;uniqueIndex;not null"`
	Credit      int     `gorm:"column:credit;not null"`
	ReviewCount int     `gorm:"column:review_count;default:0"`
	Rate        string `gorm:"column:rate;default:0"`
	Score 		float64 `gorm:"column:score;default:0"`

	Professors []Professor `gorm:"many2many:course_professors;"`
	
	CreatedAt time.Time
	UpdatedAt time.Time
}