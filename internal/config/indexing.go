package config

import (
	"gorm.io/gorm"
)

func MigrateCourseFTS(db *gorm.DB) error {
    if err := db.Exec(`
        CREATE OR REPLACE FUNCTION courses_fts_trigger() RETURNS trigger AS $$
        BEGIN
            NEW.fts := to_tsvector('english', 
                coalesce(NEW.name, '') || ' ' || 
                coalesce(NEW.code, '') || ' ' || 
                coalesce(NEW.description, '')
            );
            RETURN NEW;
        END;
        $$ LANGUAGE plpgsql;
    `).Error; err != nil {
        return err
    }

    if err := db.Exec(`
        DROP TRIGGER IF EXISTS courses_fts_update ON courses;
        CREATE TRIGGER courses_fts_update 
        BEFORE INSERT OR UPDATE ON courses 
        FOR EACH ROW EXECUTE FUNCTION courses_fts_trigger();
    `).Error; err != nil {
        return err
    }

    if err := db.Exec(`
        UPDATE courses SET fts = to_tsvector('english', 
            coalesce(name, '') || ' ' || 
            coalesce(code, '') || ' ' || 
            coalesce(description, '')
        )
        WHERE fts IS NULL
    `).Error; err != nil {
        return err
    }
    
    return nil
}