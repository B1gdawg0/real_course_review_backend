package model

import "time"

type Report struct {
	ID         string     `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID     string     `gorm:"type:uuid;not null;index"`
	ReviewID   string     `gorm:"type:uuid;not null;index"`
	ReportType ReportType `gorm:"type:int;not null;check:report_type IN (1,2,3,4)"`

	User   	   User   	  `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Review     Review 	  `gorm:"foreignKey:ReviewID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	CreatedAt time.Time
	UpdatedAt time.Time
}