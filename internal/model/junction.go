package model

type CourseTag struct {
    CourseID string `gorm:"type:uuid;not null;primaryKey"`
    TagID    string `gorm:"type:uuid;not null;primaryKey"`
    Course   Course `gorm:"foreignKey:CourseID;references:ID"`
    Tag      Tag    `gorm:"foreignKey:TagID;references:ID"`
}

type ReviewTag struct {
    ReviewID  string `gorm:"type:uuid;not null;primaryKey"`
    TagID     string `gorm:"type:uuid;not null;primaryKey"`
    Review    Review `gorm:"foreignKey:ReviewID;references:ID"`
    Tag       Tag    `gorm:"foreignKey:TagID;references:ID"`
}