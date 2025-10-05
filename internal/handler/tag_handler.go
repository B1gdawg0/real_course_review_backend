package handler

import (
	"github.com/B1gdawg0/real_course_review_backend/internal/dtos"
	uc "github.com/B1gdawg0/real_course_review_backend/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type TagHandler struct {
	tc uc.TagUseCase
}

func NewTagHandler(tc uc.TagUseCase) *TagHandler{
	return &TagHandler{tc: tc}
}

func (th *TagHandler) GetAllOrOne(c *fiber.Ctx) error {
	id := c.Query("id")

	if id != "" {
		data, err := th.tc.GetTagById(id)
		return dtos.Respond(c, data, err)
	}

	data, err := th.tc.GetAll()
	return dtos.Respond(c, data, err)
}