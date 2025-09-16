package model

import "time"

type SystemConfig struct {
    ID        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
    Key       string    `gorm:"unique;type:text"`
    Value     string    `gorm:"type:text"`

    CreatedAt time.Time
    UpdatedAt time.Time
}