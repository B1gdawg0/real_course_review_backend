package handler

import (
	"strconv"

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

func (ch *CourseHandler) Search(c *fiber.Ctx) error {
    keyword := c.Query("q", "")
    if keyword == "" {
        return fiber.NewError(fiber.StatusBadRequest, "missing keyword")
    }
    
    page, _ := strconv.Atoi(c.Query("page", "1"))
    size, _ := strconv.Atoi(c.Query("size", "10"))
    
    data, page, size, total, totalPages, err := ch.cc.Search(keyword, page, size)
    
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