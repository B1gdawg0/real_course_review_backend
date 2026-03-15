package handler

import (
	"strconv"

	"github.com/B1gdawg0/real_course_review_backend/internal/dtos"
	uc "github.com/B1gdawg0/real_course_review_backend/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type ReviewHandler struct {
	rc uc.ReviewUseCase
}

func NewReviewHandler(rc uc.ReviewUseCase) *ReviewHandler {
	return &ReviewHandler{rc: rc}
}

func (h *ReviewHandler) GetReviews(c *fiber.Ctx) error {
	filterType := c.Params("filterType")
	id := c.Params("id")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	size, _ := strconv.Atoi(c.Query("size", "10"))
	userID, _ := c.Locals("userID").(string)

	var data []dtos.ReviewFullResponse
	var total int64
	var totalPages int
	var err error

	switch filterType {
	case "c":
		data, page, size, total, totalPages, err = h.rc.GetReviewsByCourseID(userID, id, page, size)
	// case "u":
	// 	data, page, size, total, totalPages, err = h.rc.GetReviewsByUserID(userID, page, size)
	default:
		return fiber.NewError(fiber.StatusBadRequest, "invalid filter type")
	}

	if err != nil {
		return dtos.Respond(c, data, err)
	}


	return dtos.RespondWithMeta(c, data, err, dtos.PaginatedResponse{
		Page:       page,
		PageSize:   size,
		Total:      total,
		TotalPages: totalPages,
	})
}

func (h *ReviewHandler) CreateReview(c *fiber.Ctx) error {
	var req dtos.CreateReviewRequest
	var err error
	userID, _ := c.Locals("userID").(string)

	if err = c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	if req.CourseID == "" || req.ProfessorID == "" {
		return fiber.NewError(fiber.StatusBadRequest, "missing required fields")
	}

    reviewDTO, err := h.rc.CreateReview(req, userID)

    return dtos.Respond(c,reviewDTO, err)
}

func (h *ReviewHandler) UpdateReviewRecStatus(c *fiber.Ctx) error {
	role, _ := c.Locals("role").(string)
	if role != "ADMIN" {
		return fiber.NewError(fiber.StatusForbidden, "only admins can update review recommendation status")
	}

	var req dtos.UpdateReviewRecStatusRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	if req.ID == "" {
		return fiber.NewError(fiber.StatusBadRequest, "missing required fields")
	}

	err := h.rc.UpdateReviewRecStatus(req.ID, req.RecStatus)

	return dtos.RespondNoContent(c, err)
}