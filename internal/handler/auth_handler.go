package handler

import (
	"github.com/B1gdawg0/real_course_review_backend/internal/dtos"
	"github.com/B1gdawg0/real_course_review_backend/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	usecase usecase.AuthUsecase
}

func NewAuthHandler(usecase usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{usecase: usecase}
}

func (h *AuthHandler) HandleAuth(c *fiber.Ctx) error {
	var req dtos.AuthRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	resp, err := h.usecase.LoginOrRegister(&req)
	
	return dtos.Respond(c, resp, err)
}
