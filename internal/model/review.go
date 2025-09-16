package model

import "time"

type Review struct {
	ID          string  `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID      string  `gorm:"type:uuid;not null;index"`
	CourseID     string  `gorm:"type:uuid;not null;index"`

	Status      string  `gorm:"type:varchar(20);check:status IN ('ban','publish','delete','hidden');default:'publish'"`
	Description string  `gorm:"type:text"`

	AvgRate     float64 `gorm:"column:avg_rate"`
	Score       float64 `gorm:"column:score;default:0"`

	UpCount     int `gorm:"default:0"`
	DownCount   int `gorm:"default:0"`
	ReportCount int `gorm:"default:0"`
	IsAnonymous bool `gorm:"default:false"`

	User       User       `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Course      Course      `gorm:"foreignKey:CourseID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Votes      []Vote     `gorm:"foreignKey:ReviewID"`
	Rate string   		  `gorm:"column:rate"`

	CreatedAt time.Time
	UpdatedAt time.Time
}