package config

import (
	"fmt"
	"log"
	"os"

	m "github.com/B1gdawg0/real_course_review_backend/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DropAllTables drops all tables in the database (fresh start)
func DropAllTables(db *gorm.DB) error {
	log.Println("🗑️  Dropping all tables...")

	// Drop tables in order (respecting foreign key constraints)
	tables := []string{
		"review_tags",
		"course_tags",
		"course_professors",
		"reports",
		"votes",
		"reviews",
		"tags",
		"courses",
		"professors",
		"users",
		"system_configs",
	}

	for _, table := range tables {
		if err := db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE", table)).Error; err != nil {
			log.Printf("⚠️  Warning: Failed to drop table %s: %v", table, err)
		} else {
			log.Printf("✅ Dropped table: %s", table)
		}
	}

	log.Println("✅ All tables dropped successfully")
	return nil
}

func InitDB(cfg *m.Config, schema ...interface{}) *gorm.DB {
	// Debug log to verify connection parameters
	log.Printf("📡 Connecting to PostgreSQL...")
	log.Printf("   Host: %s", cfg.DB_HOST)
	log.Printf("   Port: %s", cfg.DB_PORT)
	log.Printf("   User: %s", cfg.DB_USER)

	// Build DSN - omit dbname if not specified
	var dsn string
	if cfg.DB_NAME != "" {
		log.Printf("   Database: %s", cfg.DB_NAME)
		dsn = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			cfg.DB_HOST, cfg.DB_PORT, cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_NAME,
		)
	} else {
		log.Printf("   Database: (default)")
		dsn = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s sslmode=disable",
			cfg.DB_HOST, cfg.DB_PORT, cfg.DB_USER, cfg.DB_PASSWORD,
		)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}
	log.Println("✅ Database connection established")

	// Check if RESET_DB flag is set
	if os.Getenv("RESET_DB") == "true" {
		log.Println("🔄 RESET_DB=true detected, dropping all tables...")
		if err := DropAllTables(db); err != nil {
			log.Fatal("Failed to drop tables: " + err.Error())
		}
	}

	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "pgcrypto"`).Error; err != nil {
		log.Fatal("Failed to set up database: " + err.Error())
	}

	log.Println("🔧 Running AutoMigrate for all schemas...")
	err = db.AutoMigrate(schema...)
	if err != nil {
		log.Fatal("Failed to auto migrate: " + err.Error())
	}
	log.Println("✅ AutoMigrate completed successfully")

	if err := MigrateCourseFTS(db); err != nil {
		log.Fatal("Failed to setup FTS:", err)
	}

	return db
}
