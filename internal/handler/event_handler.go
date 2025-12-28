package handler

import (
	"time"

	"github.com/B1gdawg0/real_course_review_backend/internal/dtos"
	"github.com/gofiber/fiber/v2"
)

func CollectEvent(queue chan<- dtos.Event) fiber.Handler {
    return func(c *fiber.Ctx) error {
        var e dtos.Event
        if err := c.BodyParser(&e); err != nil {
            return c.SendStatus(400)
        }

        e.Ts = time.Now().Unix()

        select {
        case queue <- e:
        default:
        }

        return c.SendStatus(204)
    }
}
