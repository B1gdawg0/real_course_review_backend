package main

import (
	"log"

	"github.com/B1gdawg0/real_course_review_backend/internal/background"
	"github.com/B1gdawg0/real_course_review_backend/internal/config"
	"github.com/B1gdawg0/real_course_review_backend/internal/dtos"
	"github.com/B1gdawg0/real_course_review_backend/internal/handler"
	m "github.com/B1gdawg0/real_course_review_backend/internal/model"
	f "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	eventQueue := make(chan dtos.Event, 5000)
	background.StartEventWorker(eventQueue)

	cfg := config.Load()

	db := config.InitDB(cfg, m.ALL_SCHEMA...)

	config.Seed(db)

	app := f.New()
	app.Use(cors.New())

	app.Get("/health", func(c *f.Ctx) error {
		return c.SendString("OK")
	})

	app.Post("/events", handler.CollectEvent(eventQueue))

	config.RegisterRoutesV1(app, db, cfg)

	log.Printf("Starting server on port %s...", cfg.PORT)
	if err := app.Listen(":" + cfg.PORT); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
