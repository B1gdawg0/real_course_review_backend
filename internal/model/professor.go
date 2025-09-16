package model

import "time"

type Professor struct {
	ID              string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name            string `gorm:"column:full_name;not null"`
	Email           string `gorm:"column:email;not null;uniqueIndex"`
	Title           string `gorm:"column:title"`
	ImageURL        string `gorm:"column:image_url"`
	Phone           string `gorm:"column:phone"`
	Description     string `gorm:"column:description;type:text"`
	UniRoomAddress  string `gorm:"column:uni_room_address"`

	Classes []Course `gorm:"many2many:course_professors;foreignKey:ID;joinForeignKey:ProfessorID;References:ID;joinReferences:CourseID"`
	
	CreatedAt time.Time
	UpdatedAt time.Time
}