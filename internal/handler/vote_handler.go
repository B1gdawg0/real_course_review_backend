package handler

import (
	"github.com/B1gdawg0/real_course_review_backend/internal/dtos"
	uc "github.com/B1gdawg0/real_course_review_backend/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type VoteHandler struct {
	vc uc.VoteUseCase
}

func NewVoteHandler(vc uc.VoteUseCase) *VoteHandler{
	return &VoteHandler{vc: vc}
}

func (h *VoteHandler) Vote(c *fiber.Ctx) error{
	var req dtos.VoteReviewRequest
	var err error
	userID, _ := c.Locals("userID").(string)

	if err = c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	if req.ReviewID == "" || userID == "" {
		return fiber.NewError(fiber.StatusBadRequest, "missing required fields")
	}

    err = h.vc.AddVoteToReview(userID, req.ReviewID, req.Vote)

    return dtos.Respond(c, "",err)
}