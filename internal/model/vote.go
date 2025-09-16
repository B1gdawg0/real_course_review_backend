package model

import "time"

type Vote struct {
	ID       string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID   string `gorm:"column:user_id;type:uuid;not null;index"`
	ReviewID string `gorm:"column:review_id;type:uuid;not null;index"`
	Vote     int    `gorm:"column:vote;check:vote IN (-1, 1)"`

	User   User   `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Review Review `gorm:"foreignKey:ReviewID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	
	CreatedAt time.Time
	UpdatedAt time.Time
}