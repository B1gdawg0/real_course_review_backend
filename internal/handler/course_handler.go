package handler

import (
	"github.com/B1gdawg0/real_course_review_backend/internal/dtos"
	uc "github.com/B1gdawg0/real_course_review_backend/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type CourseHandler struct {
	cc uc.CourseUseCase
}

func NewClassHandler(cc uc.CourseUseCase) *CourseHandler{
	return &CourseHandler{cc: cc}
}

func (ch *CourseHandler) GetAllOrOne(c *fiber.Ctx) error {
	id := c.Query("id")

	if id != "" {
		data, err := ch.cc.GetCourseById(id)
		return dtos.Respond(c, data, err)
	}

	data, err := ch.cc.GetAll()
	return dtos.Respond(c, data, err)
}