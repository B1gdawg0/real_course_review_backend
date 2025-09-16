package model

import "time"

type User struct{
	ID      string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name    string `gorm:"column:full_name"`
	Email   string `gorm:"column:email_address;uniqueIndex"`
	Role    string `gorm:"column:user_role;type:varchar(20);not null;default:'USER'"`
	UniYear string `gorm:"column:university_year"`
	
	Reviews []Review `gorm:"foreignKey:UserID"`
	Votes   []Vote   `gorm:"foreignKey:UserID"`
	
	CreatedAt time.Time
	UpdatedAt time.Time
}