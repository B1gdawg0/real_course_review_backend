package config

import (
	"fmt"
	"log"

	m "github.com/B1gdawg0/real_course_review_backend/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg *m.Config, schema ...interface{}) *gorm.DB {
    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        cfg.DB_HOST, cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_NAME, cfg.DB_PORT,
    )
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: " + err.Error())
    }

	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "pgcrypto"`).Error; err != nil {
        log.Fatal("Failed to set up database: " + err.Error())
    }

	err = db.AutoMigrate(schema...)
	if err != nil {
		log.Fatal("Failed to auto migrate: " + err.Error())
	}

	if err := MigrateCourseFTS(db); err != nil {
        log.Fatal("Failed to setup FTS:", err)
    }

    return db
}