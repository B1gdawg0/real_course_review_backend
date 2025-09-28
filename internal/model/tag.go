package model

import "time"

type Tag struct {
    ID        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
    Name      string    `gorm:"column:name;uniqueIndex;not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
}