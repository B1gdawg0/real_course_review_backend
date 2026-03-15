package handler

import (
	"github.com/B1gdawg0/real_course_review_backend/internal/dtos"
	uc "github.com/B1gdawg0/real_course_review_backend/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type ReportHandler struct {
	rc uc.ReportUseCase
}

func NewReportHandler(rc uc.ReportUseCase) *ReportHandler{
	return &ReportHandler{rc: rc}
}


func (rh *ReportHandler) CreateReport(c *fiber.Ctx) error {
	var req dtos.ReportRequest
	var err error
	userID, _ := c.Locals("userID").(string)

	if err = c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	if req.ReviewID == ""{
		return fiber.NewError(fiber.StatusBadRequest, "missing required fields")
	}

	req.UserID = userID

    err = rh.rc.AddReportToReview(req)

    return dtos.Respond(c,"", err)
}

func (rh *ReportHandler) GetAll(c *fiber.Ctx) error {
	role, _ := c.Locals("role").(string)
	if role != "ADMIN" {
		return fiber.NewError(fiber.StatusForbidden, "only admins can delete courses")
	}

	data, err := rh.rc.GetAllReports()
	return dtos.Respond(c, data, err)
}

func (rh *ReportHandler) SolveReport(c *fiber.Ctx) error {
	role, _ := c.Locals("role").(string)
	if role != "ADMIN" {
		return fiber.NewError(fiber.StatusForbidden, "only admins can delete courses")
	}

	var req dtos.SolveReportRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	if req.ReviewID == "" || req.Reason == "" {
		return fiber.NewError(fiber.StatusBadRequest, "missing required fields")
	}

	err := rh.rc.SolveReport(req.ID, req.ReviewID, req.Action, req.Reason)
	return dtos.RespondNoContent(c, err)
}