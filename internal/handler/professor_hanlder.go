package handler

import (
	"github.com/B1gdawg0/real_course_review_backend/internal/dtos"
	uc "github.com/B1gdawg0/real_course_review_backend/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type ProfessorHandler struct {
	pc uc.ProfessorUseCase
}

func NewProfessorHandler(pc uc.ProfessorUseCase) *ProfessorHandler{
	return &ProfessorHandler{pc: pc}
}

func (ph *ProfessorHandler) GetAllOrOne(c *fiber.Ctx) error {
	id := c.Query("id")

	if id != "" {
		data, err := ph.pc.GetProfessorById(id)
		return dtos.Respond(c, data, err)
	}

	data, err := ph.pc.GetAll()
	return dtos.Respond(c, data, err)
}