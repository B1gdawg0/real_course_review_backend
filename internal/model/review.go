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
	DecayedScore float64 `gorm:"column:decayed_score;->;-:migration"`

	UpCount   int `gorm:"->;-:migration"`
	DownCount int `gorm:"->;-:migration"`
	UserVote  *int  `gorm:"column:user_vote"` // plan to ->;-:migration his ass but not sure so gotta leave it like this for now
	ReportCount int `gorm:"default:0"`
	IsAnonymous bool `gorm:"default:false"`

	// sync with course (reference key) later
	Grade       string			`gorm:"column:reviewer_grade"`
	Year		string 			`gorm:"column:course_year"`
	Sec			string 			`gorm:"column:section_id"`

	RecStatus	bool			`gorm:"column:rec_status;default:true"`

	User       User       `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Course      Course      `gorm:"foreignKey:CourseID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Votes      []Vote     `gorm:"foreignKey:ReviewID"`
	Rate	    string   		  `gorm:"column:rate"`
	Tags        []Tag     `gorm:"many2many:review_tags;"`

	CreatedAt time.Time
	UpdatedAt time.Time
}