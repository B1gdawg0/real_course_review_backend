package handler

import (
	"errors"
	"strconv"

	"github.com/B1gdawg0/real_course_review_backend/internal/dtos"
	"github.com/B1gdawg0/real_course_review_backend/internal/model"
	uc "github.com/B1gdawg0/real_course_review_backend/internal/usecase"
	"github.com/B1gdawg0/real_course_review_backend/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type CourseHandler struct {
	cc uc.CourseUseCase
}

func NewClassHandler(cc uc.CourseUseCase) *CourseHandler {
	return &CourseHandler{cc: cc}
}

func (ch *CourseHandler) GetCourses(c *fiber.Ctx) error {
	id := c.Query("id")
	keyword := c.Query("q")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	size, _ := strconv.Atoi(c.Query("size", "10"))

	filter := model.CourseFilter{
		Semester:   c.Query("semester"),
        CourseType: c.Query("course_type"),
        TagIDs:     utils.ParseCommaSeparated(c.Query("tag_ids")),
	}

	if id != "" && keyword != "" {
		return fiber.NewError(fiber.StatusBadRequest, "cannot use both 'id' and 'q' together")
	}

	if id != "" {
		data, err := ch.cc.GetCourseById(id)
		return dtos.Respond(c, data, err)
	}

	if keyword != "" {
		page, _ := strconv.Atoi(c.Query("page", "1"))
		size, _ := strconv.Atoi(c.Query("size", "10"))

		data, page, size, total, totalPages, err := ch.cc.Search(keyword, page, size, filter)
		return dtos.RespondWithMeta(c, data, err, dtos.PaginatedResponse{
			Page:       page,
			PageSize:   size,
			Total:      total,
			TotalPages: totalPages,
		})
	}

	data, page, size, total, totalPages, err := ch.cc.GetAll(page, size, filter)
	return dtos.RespondWithMeta(c, data, err, dtos.PaginatedResponse{
		Page:       page,
		PageSize:   size,
		Total:      total,
		TotalPages: totalPages,
	})
}

func (ch *CourseHandler) CompareCourseByIds(c *fiber.Ctx) error {
	first := c.Params("first")
	second := c.Params("second")
	userId, _ := c.Locals("userID").(string)

	if first == "" || second == "" {
		return fiber.NewError(fiber.StatusBadRequest, "both course IDs must be provided")
	}

	data, err := ch.cc.CompareCoursesById(userId, first, second)
	return dtos.Respond(c, data, err)
}

func (ch *CourseHandler) GetAISummary(c *fiber.Ctx) error {
	first := c.Params("first")
	second := c.Params("second")

	if first == "" || second == "" {
		return fiber.NewError(fiber.StatusBadRequest, "both course IDs must be provided")
	}

	data, err := ch.cc.GetAISummaryWithN8N(first, second)
    
    if err != nil {
        return dtos.Respond(c, dtos.CourseAISummaryResponse{}, err)
    }

    if data == nil {
        return dtos.Respond(c, dtos.CourseAISummaryResponse{}, errors.New("no data returned"))
    }

    return dtos.Respond(c, *data, nil)
}
